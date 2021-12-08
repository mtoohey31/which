package which

import (
	"errors"
	"os"
	"os/user"
	p "path"
	"strings"
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

		if isExecutableBy(currentUid, currentGids, info) {
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
