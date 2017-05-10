// mirth-chart project main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justinohms/mirthgraph/fileutils"
	"github.com/justinohms/mirthgraph/launcher"
	"github.com/justinohms/mirthgraph/mirth"
	"github.com/justinohms/mirthgraph/server"
)

//read from the command line
var srcDirP = flag.String("src", "", "The directory containing the mirth source xml files.")

var channels = make(map[string]mirth.Channel)
var channelPaths []string

/*
NOTE about static vs. local content

TL/DR; Adding a new file? Create empty file, run 'go generate', set "useLocal" flag

When setting the UI to serve up live file system files using the "useLocal" flag
if you add new files you must still run the 'go generate' command in order for those
files to be served locally.  This is because even if the file is being served from
from the local file system, the presense of the file is checked using the static
resources table.

The following line is used by go generate to build in static resources DO NOT DELETE
*/
//go:generate esc -o server/ui.go -pkg server ui

func main() {
	flag.Parse()

	srcDir := ""
	fmt.Println("Mirth Chart\nscanning directory...")
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
	fmt.Println(srcDir)

	//load the channelPaths
	s := mirth.Scanner{}
	channelPaths, err := s.FindAllChannelFiles(srcDir)
	check(err)

	//parse the channels
	for i := 0; i < len(channelPaths); i++ {
		ch := mirth.ParseChannelFile(channelPaths[i])
		channels[ch.Id] = ch

	}

	g := mirth.ToGraphJson(channels)
	//fmt.Println(g)

	portchannel := make(chan int)

	datachannel := make(chan string)

	finishedchannel := make(chan bool)

	go server.ServeDynamicContent(portchannel, datachannel, finishedchannel)
	//fmt.Println("after")
	serverport := <-portchannel
	fmt.Println("server started on port:", serverport)

	// send the data in
	datachannel <- g

	//this is just for development
	//return

	url := fmt.Sprintf("http://127.0.0.1:%d/ui/graph.html", serverport)
	launcher.OpenURL(url)

	for {
		//loop forever
		//maybe we monitor the folder in this loop
		//wait for something from finished channel then exit
		<-finishedchannel
		fmt.Println("Complete, see your browser for directed graph diagram.\nexiting")

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
