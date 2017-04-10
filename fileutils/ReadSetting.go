package fileutils

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func FileSetting(p string) (content string) {
	absp, err := filepath.Abs(p)
	check(err)
	//	fmt.Println(absp)
	b, err := ioutil.ReadFile(absp)

	//fmt.Println(b)
	//fmt.Println(err)

	if err != nil && strings.Contains(string(err.Error()), ".settings: no such file or directory") {
		return ""
	} else {
		check(err)
	}

	return strings.TrimSpace(string(b))
}
