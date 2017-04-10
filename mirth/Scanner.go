package mirth

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/justinohms/mirthchart/fileutils"
)

var channelPaths []string

func FindAllChannelFiles(path string) ([]string, error) {
	dirsize, _ := fileutils.FileCountRecursive(path)

	channelPaths = make([]string, 0, dirsize)

	err := filepath.Walk(path, visit)
	return channelPaths, err
}

func visit(p string, f os.FileInfo, err error) error {

	//if it's not a dir and is an xml file we are interested
	if !f.IsDir() && strings.ToLower(path.Ext(p)) == ".xml" {
		fl, err := os.Open(p)
		defer fl.Close()
		check(err)
		//fmt.Printf("name: %s\n", f.Name())
		b1 := make([]byte, 25)
		fl.ReadAt(b1, 0)
		fl.Close()

		if strings.Contains(string(b1), "<channel version=") {
			channelPaths = append(channelPaths, p)
		}
	}
	return nil
}
