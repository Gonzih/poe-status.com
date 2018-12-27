package server

import (
	"net/http"
	"strings"
)

type safeFileSystem struct {
	fs http.FileSystem
}

func (sfs safeFileSystem) Open(path string) (http.File, error) {
	f, err := sfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := sfs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
