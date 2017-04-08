// mirth-chart project main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var srcDirP = flag.String("src", "/Users/justinohms/Dropbox/nc/src/salmon-mirth/src/mirth/channel-groups", "The directory containing the mirth source xml files.")

//var srcDirP = flag.String("src", "", "The directory containing the mirth source xml files.")

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

	findAllXmlFiles(srcDir)
}

func findAllXmlFiles(path string) {
	err := filepath.Walk(path, visit)
	fmt.Printf("filepath walk error %v \n", err)
}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}
