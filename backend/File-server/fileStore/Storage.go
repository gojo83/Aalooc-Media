package fileStore

import "io"

type Storage interface {
	Save(path string, content io.Reader) error
}
