package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"

	"path"
)

func (_escStaticFS) IoCopy(w io.Writer, name string) {
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

func FSIoCopy(w io.Writer, name string) {
	_escStatic.IoCopy(w, name)

}
