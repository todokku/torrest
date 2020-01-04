package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/i96751414/torrest/bittorrent"
)

type NewTorrentResponse struct {
	InfoHash string `json:"info_hash" example:"000102030405060708090a0b0c0d0e0f10111213"`
}

type FileInfoResponse []*bittorrent.FileInfo

type TorrentInfoResponse struct {
	InfoHash string                    `json:"info_hash"`
	Status   *bittorrent.TorrentStatus `json:"status,omitempty"`
}

// @Summary Add Magnet
// @Description add magnet to service
// @ID add-magnet
// @Produce json
// @Param uri query string true "magnet URI"
// @Success 200 {object} NewTorrentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /add/magnet [get]
func addMagnet(service *bittorrent.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		magnet := ctx.Query("uri")
		if !strings.HasPrefix(magnet, "magnet:") {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid magnet provided"})
			return
		}
		if infoHash, err := service.AddMagnet(magnet); err != nil {
			ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		} else {
			ctx.JSON(http.StatusOK, NewTorrentResponse{InfoHash: infoHash})
		}
	}
}

// @Summary List Torrents
// @Description list all torrents from service
// @ID list-torrents
// @Produce json
// @Param status query boolean false "get torrents status"
// @Success 200 {array} TorrentInfoResponse
// @Router /torrents [get]
func listTorrents(service *bittorrent.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		torrents := service.Torrents()
		response := make([]TorrentInfoResponse, len(torrents))
		for i, torrent := range torrents {
			response[i].InfoHash = torrent.InfoHash()
		}
		if ctx.DefaultQuery("status", "false") == "true" {
			for i, torrent := range torrents {
				response[i].Status = torrent.GetStatus()
			}
		}
		ctx.JSON(http.StatusOK, response)
	}
}

// @Summary Remove Torrent
// @Description remove torrent from service
// @ID remove-torrent
// @Produce json
// @Param infoHash path string true "torrent info hash"
// @Param delete query boolean false "delete files"
// @Success 200 {object} MessageResponse
// @Failure 404 {object} ErrorResponse
// @Router /torrents/{infoHash}/remove [get]
func removeTorrent(service *bittorrent.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		infoHash := ctx.Param("infoHash")
		removeFiles := ctx.DefaultQuery("delete", "true") == "true"

		if err := service.RemoveTorrent(infoHash, removeFiles); err == nil {
			ctx.JSON(http.StatusOK, MessageResponse{Message: fmt.Sprintf("Torrent '%s' deleted", infoHash)})
		} else {
			ctx.JSON(http.StatusNotFound, NewErrorResponse(err))
		}
	}
}

// @Summary Get Torrent Status
// @Description get torrent status
// @ID torrent-status
// @Produce json
// @Param infoHash path string true "torrent info hash"
// @Success 200 {object} bittorrent.TorrentStatus
// @Failure 404 {object} ErrorResponse
// @Router /torrents/{infoHash}/status [get]
func torrentStatus(service *bittorrent.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		infoHash := ctx.Param("infoHash")
		if torrent, err := service.GetTorrent(infoHash); err == nil {
			ctx.JSON(http.StatusOK, torrent.GetStatus())
		} else {
			ctx.JSON(http.StatusNotFound, NewErrorResponse(err))
		}
	}
}

// @Summary Get Torrent Files
// @Description get a list of the torrent files and its details
// @ID torrent-files
// @Produce json
// @Param infoHash path string true "torrent info hash"
// @Success 200 {object} FileInfoResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /torrents/{infoHash}/files [get]
func torrentFiles(service *bittorrent.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		infoHash := ctx.Param("infoHash")
		if torrent, err := service.GetTorrent(infoHash); err == nil {
			if torrent.HasMetadata() {
				files := torrent.Files()
				response := make(FileInfoResponse, len(files))
				for i, file := range files {
					response[i] = file.Info()
				}
				ctx.JSON(http.StatusOK, response)
			} else {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no metadata"})
			}
		} else {
			ctx.JSON(http.StatusNotFound, NewErrorResponse(err))
		}
	}
}
