//go:build !windows

package which

import (
	"fmt"
	"io/fs"
	"os"
	"syscall"
)

func isExecutableBy(uid string, gids []string, fileInfo fs.FileInfo) bool {
	mode := fileInfo.Mode()
	stat := fileInfo.Sys().(*syscall.Stat_t)

	if isExecOwner(mode) && fmt.Sprint(stat.Uid) == uid {
		return true
	} else if isExecGroup(mode) && contains(gids, fmt.Sprint(stat.Gid)) {
		return true
	} else if isExecOther(mode) {
		return true
	}
	return false
}

// source: https://stackoverflow.com/a/60128480
func isExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}

// source: https://stackoverflow.com/a/60128480
func isExecGroup(mode os.FileMode) bool {
	return mode&0010 != 0
}

// source: https://stackoverflow.com/a/60128480
func isExecOther(mode os.FileMode) bool {
	return mode&0001 != 0
}
