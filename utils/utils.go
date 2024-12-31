package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func GetPublicPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return findPublicPath(dir)
}

func findPublicPath(_path string) (string, error) {
	public := filepath.Join(_path, "public")
	exists, err := directoryExists(public)
	if err != nil {
		return "", err
	}
	if exists == true {
		return public, nil
	}
	if _path == "/" {
		return "", errors.New("reached root without finding public directory")
	}
	parent := filepath.Dir(_path)
	return findPublicPath(parent)
}

func directoryExists(_path string) (bool, error) {
	info, err := os.Stat(_path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}
