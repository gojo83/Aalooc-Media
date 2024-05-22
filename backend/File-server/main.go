package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Shahriar-shudip/golang-microservies-tuitorial/File-server/fileStore"
	"github.com/Shahriar-shudip/golang-microservies-tuitorial/File-server/handlers"
	goCros "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "api", log.LstdFlags)

	localStorage, _ := fileStore.NewLocal("imageStore", 1024*1000*5)
	fileHandler := handlers.NewFile(localStorage, l)

	sm := mux.NewRouter()
	ch := goCros.CORS(goCros.AllowedOrigins([]string{"*"}))

	uploadFile := sm.Methods(http.MethodPost).Subrouter()
	uploadFile.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fileHandler.UploadRest)
	uploadFile.HandleFunc("/", fileHandler.UploadMultiPart)

	getFile := sm.Methods(http.MethodGet).Subrouter()
	getFile.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z0-9]+\\.[a-z]{3}}", http.StripPrefix("/images/", http.FileServer(http.Dir("./imageStore"))))

	server := http.Server{
		Addr:         ":9001",
		Handler:      ch(sm), //to allow corsorigin
		IdleTimeout:  123 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("gracful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc) // shuts the server when users has done with the request
}
