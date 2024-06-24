package fsutil

import (
	"errors"
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
	stat, err := os.Lstat(p)
	if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return true, nil
	}

	return false, nil
}

func Validate(path string) error {
	if path == "" {
		return errors.New("**Invalid Target:**\nThe provided target path must be a valid file or directory.")
	}

	fileInfo, err := IsPathValidAndAccessible(path)
	if err != nil {
		return err
	}

	// Throw if the target path is neither a file or a directory
	if !fileInfo.IsDir() && !fileInfo.Mode().IsRegular() {
		return errors.New("**Invalid Target:**\nExpected the target to be either a directory or file.")
	}

	return nil
}

func IsPathValidAndAccessible(path string) (os.FileInfo, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("**Invalid Target:**\nThe provided path does not exist.")
		}

		if os.IsPermission(err) {
			return nil, errors.New("**Permission Denied:**\nCould not access the provided path.")
		}

		return nil, errors.New("**Validation Error:**\nAn unknown error occurred while validating the target.")
	}

	return fileInfo, nil
}
