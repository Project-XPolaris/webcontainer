package filesystem

import (
	"net/http"
	"os"
)

type FileSystem struct {
	Root http.FileSystem
}

func (fs *FileSystem) Open(name string) (http.File, error) {
	f, err := fs.Root.Open(name)
	if os.IsNotExist(err) {
		return fs.Root.Open("index.html")
	}
	return f, err
}
