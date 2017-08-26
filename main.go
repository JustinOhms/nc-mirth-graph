// mirth-chart project main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/caimeo/iniflags"

	"github.com/justinohms/mirthgraph/launcher"
	"github.com/justinohms/mirthgraph/mirth"
	"github.com/justinohms/mirthgraph/server"

	"github.com/caimeo/console"
)

//read from the command line
var srcDirP = flag.String("src", "", "The directory containing the mirth source xml files.")
var customUIP = flag.Bool("custom_ui", false, "Serve custom UI instead of the embedded UI.")
var verboseMode = flag.Bool("verbose", false, "Verbose console output.")

var channels = make(map[string]mirth.Channel)
var channelPaths []string

var con console.Console

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
	iniflags.SetConfigFile(".settings")
	iniflags.SetAllowMissingConfigFile(true)
	iniflags.Parse()

	con := console.Init(*verboseMode, false)

	fmt.Println("Mirth Chart")

	useLocal := *customUIP
	if useLocal {
		con.Verbose("Using custom UI")
	}

	srcDir := ""
	con.Verbose("scanning directory...")
	if *srcDirP == "" {
		srcDir, _ = os.Getwd()
	} else {
		srcDir = *srcDirP
	}
	con.Always(srcDir)

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

	server.UseLocal = useLocal
	go server.ServeDynamicContent(portchannel, datachannel, finishedchannel)

	//fmt.Println("after")
	serverport := <-portchannel
	//fmt.Println("server started on port:", serverport)
	con.Verbose(fmt.Sprintf("server started on port: %d", serverport))

	// send the data in
	datachannel <- g

	//this is just for development
	//return

	url := fmt.Sprintf("http://127.0.0.1:%d/ui/graph.html", serverport)
	launcher.OpenURL(url)

	//for {
	//loop forever
	//maybe we monitor the folder in this loop
	//wait for something from finished channel then exit
	<-finishedchannel
	//fmt.Println("Complete, see your browser for directed graph diagram.\n")
	con.Always("Complete, see your browser for directed graph diagram.\n")
	//}

	//sleep for a bit in case the ui is still pulling resources
	time.Sleep(10 * time.Second)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
