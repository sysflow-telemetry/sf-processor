package ioutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ListFilePaths lists file paths with extension fileExt in path if
// path is a valid directory, otherwise, it returns path if path is
// a valid path and has extension fileExt.
func ListFilePaths(path string, fileExt string) ([]string, error) {
	var paths []string
	if fi, err := os.Stat(path); os.IsNotExist(err) {
		return paths, err
	} else if fi.IsDir() {
		var files []os.FileInfo
		var err error
		if files, err = ioutil.ReadDir(path); err != nil {
			return paths, err
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == fileExt {
				f := path + "/" + file.Name()
				paths = append(paths, f)
			}
		}
		return paths, nil
	} else {
		if filepath.Ext(path) == fileExt {
			return append(paths, path), nil
		}
		return paths, nil
	}
}
