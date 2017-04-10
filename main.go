// mirth-chart project main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justinohms/mirthchart/fileutils"
	"github.com/justinohms/mirthchart/mirth"
)

//read from the command line
var srcDirP = flag.String("src", "", "The directory containing the mirth source xml files.")

var channels = make(map[string]mirth.Channel)
var channelPaths []string

func main() {
	flag.Parse()

	srcDir := ""
	fmt.Println("Mirth Chart")
	if *srcDirP == "" {
		//read from a settings file
		fileSetting := fileutils.FileSetting(".settings")
		if fileSetting == "" {
			srcDir, _ = os.Getwd()
		} else {
			srcDir = fileSetting
		}
	} else {
		srcDir = *srcDirP
	}
	fmt.Println("Source Directory:", srcDir)

	//load the channelPaths
	s := mirth.Scanner{}
	channelPaths, err := s.FindAllChannelFiles(srcDir)
	check(err)

	//parse the channels
	for i := 0; i < len(channelPaths); i++ {
		ch := mirth.ParseChannelFile(channelPaths[i])
		channels[ch.FilePath] = ch

		fmt.Println(ch)

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
