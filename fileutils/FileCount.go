package fileutils

import (
	"io/ioutil"
	"path/filepath"
)

//FileCount returns the total number of files under this path
func FileCount(path string) (int, error) {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, file := range files {
		if !file.IsDir() {
			i++
		}
	}
	return i, nil
}

//FileCount returns the total number of files under this path
func FileCountRecursive(path string) (int, error) {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, file := range files {
		if file.IsDir() {
			j, _ := FileCountRecursive(filepath.Join(path, file.Name()))
			i = i + j
		} else {
			i++
		}
	}
	return i, nil
}
