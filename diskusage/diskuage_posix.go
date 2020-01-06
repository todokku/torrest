// +build !windows

package diskusage

import (
	"syscall"
)

// disk usage of path/disk
func DiskUsage(path string) (*DiskStatus, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return nil, err
	}
	status := &DiskStatus{
		All:  int64(fs.Blocks) * fs.Bsize,
		Free: int64(fs.Bfree) * fs.Bsize,
	}
	status.Used = status.All - status.Free
	return status, nil
}
