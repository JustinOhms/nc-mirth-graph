package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"

	"path"
)

func (_escStaticFS) IoCopy(name string, w io.Writer) {
	f, present := _escData[path.Clean(name)]
	if !present {
		//return nil, os.ErrNotExist
		return
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}

		io.Copy(w, gr)
	})

}

func (_escLocalFS) IoCopy(name string, w io.Writer) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return
	}
	file, _ := os.Open(f.local)
	io.Copy(w, file)
}

func FSIoCopy(useLocal bool, name string, w io.Writer) {
	if useLocal {
		_escLocal.IoCopy(name, w)
	} else {
		_escStatic.IoCopy(name, w)
	}
}
