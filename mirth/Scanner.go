package mirth

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/justinohms/mirthchart/fileutils"
)

const ChannelSig = "<channel version=" //this is the first few characters of the file

type Scanner struct {
	channelPaths []string
}

func (s *Scanner) FindAllChannelFiles(path string) ([]string, error) {
	dirsize, _ := fileutils.FileCountRecursive(path)

	s.channelPaths = make([]string, 0, dirsize)

	err := filepath.Walk(path, s.visit)
	return s.channelPaths, err
}

func (s *Scanner) visit(p string, f os.FileInfo, err error) error {

	//if it's not a dir and is an xml file we are interested
	if !f.IsDir() && strings.ToLower(path.Ext(p)) == ".xml" {
		fl, err := os.Open(p)
		check(err)
		defer fl.Close()

		//read in just the first part of the file so we can check it
		b1 := make([]byte, 25)
		fl.ReadAt(b1, 0)
		fl.Close()

		if strings.Contains(string(b1), ChannelSig) {
			s.channelPaths = append(s.channelPaths, p)
		}
	}
	return nil
}
