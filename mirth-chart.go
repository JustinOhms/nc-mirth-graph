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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findAllXmlFiles(path string) {
	err := filepath.Walk(path, visit)
	check(err)
}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)

	if !f.IsDir() {
		fl, err := os.Open(path)
		defer fl.Close()
		check(err)
		fmt.Printf("name: %s\n", f.Name())
		b1 := make([]byte, 20)
		fl.ReadAt(b1, 0)
		fmt.Printf("%s\n", string(b1))
	}

	return nil
}
