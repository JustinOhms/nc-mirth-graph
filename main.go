// mirth-chart project main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justinohms/mirthchart/mirth"
)

//The following line is the real line, uncomment it and delete the next line after
//var srcDirP = flag.String("src", "", "The directory containing the mirth source xml files.")
var srcDirP = flag.String("src", "/Users/justinohms/Dropbox/nc/src/salmon-mirth/src/mirth/channel-groups", "The directory containing the mirth source xml files.")

var channels = make(map[string]mirth.Channel)
var channelPaths []string

func main() {
	flag.Parse()

	srcDir := ""
	fmt.Println("Mirth Chart")
	if *srcDirP == "" {
		srcDir, _ = os.Getwd()
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
