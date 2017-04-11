package mirth

import (
	"encoding/json"
	"encoding/xml"

	"io/ioutil"
	"os"
)

func ParseChannelFile(p string) Channel {

	return readChannelFile(p)
}

func readChannelFile(p string) Channel {
	f, err := os.Open(p)
	defer f.Close()

	check(err)

	b, err := ioutil.ReadAll(f)
	check(err)

	var q Channel
	q.FilePath = p

	err = xml.Unmarshal(b, &q)
	check(err)

	return q
}

func ToJson(c Channel) string {
	b, err := json.Marshal(c)
	check(err)
	return string(b)
}
