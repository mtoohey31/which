package which

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	p "path"
	"strings"
	"syscall"
)

func Which(executable string) (string, error) {
	path := os.Getenv("PATH")

	errs := []error{}

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	currentUid := currentUser.Uid
	currentGids, err := currentUser.GroupIds()
	if err != nil {
		return "", err
	}

	for _, element := range strings.Split(path, ":") {
		path := p.Join(element, executable)
		info, err := os.Stat(path)

		if err != nil {
			continue
		}

		mode := info.Mode()
		stat := info.Sys().(*syscall.Stat_t)

		if isExecOwner(mode) && fmt.Sprint(stat.Uid) == currentUid {
			return path, nil
		} else if isExecGroup(mode) && contains(currentGids, fmt.Sprint(stat.Gid)) {
			return path, nil
		} else if isExecOther(mode) {
			return path, nil
		}
	}

	if len(errs) == 1 {
		return "", errs[0]
	} else if len(errs) > 1 {
		errStrs := make([]string, len(errs))
		for i, err := range errs {
			errStrs[i] = err.Error()
		}
		return "", errors.New(strings.Join(errStrs, ", "))
	}

	return "", errors.New(executable + " not found")
}

func contains(arr []string, item string) bool {
	for _, other := range arr {
		if item == other {
			return true
		}
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
