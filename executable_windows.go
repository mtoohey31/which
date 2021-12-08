//go:build windows

package which

import "io/fs"

func isExecutableBy(uid string, gids []string, fileInfo fs.FileInfo) bool {
	return true
}
