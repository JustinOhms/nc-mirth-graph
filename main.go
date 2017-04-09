// mirth-chart project main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justinohms/mirthchart/searchers"
	"github.com/justinohms/mirthchart/structs"
)

var srcDirP = flag.String("src", "/Users/justinohms/Dropbox/nc/src/salmon-mirth/src/mirth/channel-groups", "The directory containing the mirth source xml files.")

//var srcDirP = flag.String("src", "", "The directory containing the mirth source xml files.")

var channels = make(map[string]structs.MirthChannel)
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

	channelPaths, err := searchers.FindAllChannelFiles(srcDir)
	check(err)

	fmt.Println(channelPaths)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
