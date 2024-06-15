package fsutil

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func AbsPath(p string) (string, error) {
	path, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}

	return path, nil
}

func IsDirectory(p string) (bool, error) {
	stat, err := os.Stat(p)
	if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return true, nil
	}

	return false, nil
}

func GetSubDirPaths(p string) ([]string, error) {
	subdirs := make([]string, 0)

	err := filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		_, err = IsDirectory(path)
		if err != nil {
			return err
		}
		subdirs = append(subdirs, path)

		return nil
	})
	if err != nil {
		return subdirs, err
	}

	return subdirs, nil
}

func GetParentDir(p string) string {
	return filepath.Dir(p)
}

func GetFileFromPath(p string) string {
	return filepath.Base(p)
}

func IsValidPath(path string) error {
	if path == "" {
		return errors.New("**Invalid Target:**\nThe provided target path must be a valid file or directory.")
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("**Invalid Target:**\nThe provided path does not exist.")
		}

		if os.IsPermission(err) {
			return errors.New("**Permission Denied:**\nCould not access the provided path.")
		}

		return errors.New("**Validation Error:**\nAn unknown error occurred while validating the target.")
	}

	if !fileInfo.IsDir() || !fileInfo.Mode().IsRegular() {
		return errors.New("**Invalid Target:**\nExpected the target to be either a directory or file.")
	}

	return nil
}
