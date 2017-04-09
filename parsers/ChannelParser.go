package parsers

import (
	"encoding/xml"

	"io/ioutil"
	"os"

	"github.com/justinohms/mirthchart/structs"
)

func ParseChannelFile(p string) structs.MirthChannel {

	return readChannelFile(p)
}

func readChannelFile(p string) structs.MirthChannel {
	f, err := os.Open(p)
	defer f.Close()

	check(err)

	b, err := ioutil.ReadAll(f)
	check(err)

	var q structs.MirthChannel
	q.FilePath = p

	err = xml.Unmarshal(b, &q)
	check(err)

	return q
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
