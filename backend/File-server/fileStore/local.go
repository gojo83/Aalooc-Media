package fileStore

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Local struct {
	basePath string
	size     int
}

func NewLocal(b string, size int) (*Local, error) {
	p, err := filepath.Abs(b)
	if err != nil {
		return nil, err
	}
	return &Local{p, size}, nil
}

func (l *Local) Save(path string, content io.Reader) error {
	//start := time.Now()
	//adds with the basepath+pathoffile
	fp := l.fullpath(path)

	d := filepath.Dir(fp) //return the whole path excluding the last element or file name
	fmt.Println(fp)
	fmt.Println(d)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return err
	}
	file, _ := os.Create(fp)
	defer file.Close()
	_, err = io.Copy(file, content)
	if err != nil {
		return err
	}
	//elapsed := time.Since(start)
	//fmt.Println(elapsed)
	return nil
}

func (l *Local) fullpath(path string) string {

	return filepath.Join(l.basePath, path)
}
