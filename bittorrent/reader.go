package bittorrent

import (
	"errors"
	"io"
	"time"
)

const (
	piecesRefreshDuration = 500 * time.Millisecond
)

type Reader interface {
	io.Reader
	io.Seeker
	io.Closer
}

type reader struct {
	storage        Reader
	torrent        *Torrent
	offset         int64
	length         int64
	pieceLength    int64
	priorityPieces int
	closing        chan interface{}
}

func (r *reader) waitForPiece(piece int) error {
	if r.torrent.handle.HavePiece(piece) {
		return nil
	}

	log.Infof("Waiting for piece %d", piece)

	pieceRefreshTicker := time.NewTicker(piecesRefreshDuration)
	defer pieceRefreshTicker.Stop()

	for !r.torrent.handle.HavePiece(piece) {
		select {
		case <-r.torrent.closing:
			log.Warningf("Unable to wait for piece %d as torrent was closed", piece)
			return errors.New("torrent was closed")
		case <-r.closing:
			log.Warningf("Unable to wait for piece %d as file was closed", piece)
			return errors.New("file was closed")
		case <-pieceRefreshTicker.C:
			continue
		}
	}
	return nil
}

func (r *reader) pieceFromOffset(offset int64) int {
	return int((r.offset + offset) / r.pieceLength)
}

func (r *reader) Read(b []byte) (n int, err error) {
	currentOffset, err := r.storage.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}
	// TODO: Check all the pieces
	piece := r.pieceFromOffset(currentOffset + int64(len(b)))
	if err := r.waitForPiece(piece); err != nil {
		return 0, err
	}

	return r.storage.Read(b)
}

func (r *reader) Close() error {
	close(r.closing)
	return r.storage.Close()
}

func (r *reader) Seek(off int64, whence int) (ret int64, err error) {
	seekingOffset := off

	switch whence {
	case io.SeekStart:
		// do nothing
	case io.SeekCurrent:
		if currentOffset, err := r.storage.Seek(0, io.SeekCurrent); err != nil {
			return currentOffset, err
		} else {
			seekingOffset += currentOffset
		}
	case io.SeekEnd:
		seekingOffset = r.length - seekingOffset
	default:
		return 0, errors.New("invalid whence")
	}

	piece := r.pieceFromOffset(seekingOffset)
	if !r.torrent.handle.HavePiece(piece) {
		log.Infof("We don't have piece %d, setting piece priorities", piece)
		p := 0
		for ; p < piece; p++ {
			r.torrent.handle.PiecePriority(p, 0)
		}
		for ; p <= piece+r.priorityPieces; p++ {
			r.torrent.handle.PiecePriority(p, 7)
		}
		for ; p <= r.pieceFromOffset(r.length); p++ {
			r.torrent.handle.PiecePriority(p, 6)
		}
	}

	return r.storage.Seek(off, whence)
}