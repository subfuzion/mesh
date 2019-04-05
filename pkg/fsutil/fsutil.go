package fsutil

import (
	"io/ioutil"
	"net/http"
)

// Assets wraps an http.FileSystem to provide helper methods
type Assets struct {
	fs http.FileSystem
}

// NewAssets is a factory method for creating Assets
func NewAssets(fs http.FileSystem) *Assets {
	return &Assets{
		fs: fs,
	}
}

// Open returns a File specified by the path
func (a *Assets) Open(path string) (http.File, error) {
	return a.fs.Open(path)
}

// ReadAll returns the contents of a file
func (a *Assets) ReadAll(path string) ([]byte, error) {
	f, err := a.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
