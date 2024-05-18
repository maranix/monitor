package fsutil

import (
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
