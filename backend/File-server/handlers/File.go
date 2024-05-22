package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/Shahriar-shudip/golang-microservies-tuitorial/File-server/fileStore"
	"github.com/gorilla/mux"
)

var wg sync.WaitGroup

//Export File
type File struct {
	log   *log.Logger
	store fileStore.Storage
}

func NewFile(s fileStore.Storage, l *log.Logger) *File {
	return &File{log: l, store: s}
}

func (f *File) UploadRest(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	id := vars["id"]
	filename := vars["filename"]
	wg.Add(1)
	go f.saveFile(id, filename, rw, r)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
	//	wg.Wait()
}

var start time.Time

func (f *File) UploadMultiPart(rw http.ResponseWriter, r *http.Request) {
	start = time.Now()
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		http.Error(rw, "Expected multipart ", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	f.log.Println(id)
	if err != nil {
		http.Error(rw, "Expected multipart ", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(rw, "Expected multipart ", http.StatusBadRequest)
		return
	}
	f.saveMultipartFile(r.FormValue("id"), fileHeader.Filename, rw, file)
	end := time.Since(start)
	f.log.Printf("time took to handle multipart %v", end)
}

func (f *File) UploadMultipleFiles(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("UploadM")
	reader, err := r.MultipartReader()

	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			fmt.Println(err)
			break
		}

		//if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			continue
		}
		dst, err := os.Create("/home/shudip/i" + part.FileName())
		defer dst.Close()

		if err != nil {
			fmt.Println(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (f *File) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Println("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}

func (f *File) saveMultipartFile(id, path string, rw http.ResponseWriter, r io.Reader) {
	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Println("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}

}
