package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"time"

	"github.com/devShahriar/alocmedia/backend/auth/errorHandler"
	"github.com/devShahriar/alocmedia/backend/auth/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	l := log.New(os.Stdout, "Auth api", log.LstdFlags)

	userHandler := handlers.NewUserHandler(l)
	errorH := errorHandler.NewErrorHandler(l)

	sm := mux.NewRouter() //return a new router instance

	//ws handler

	wsHandler := sm.Methods(http.MethodGet).Subrouter()
	wsHandler.HandleFunc("/ws/validate/{userId:[a-z]+}", errorH.WsHandler)

	//insert user handler
	postUser := sm.Methods(http.MethodPost).Subrouter()
	postUser.HandleFunc("/insert/user", userHandler.InsertUser)

	pUser := sm.Methods(http.MethodPost).Subrouter()
	pUser.HandleFunc("/login", userHandler.Login)

	hand := cors.Default().Handler(sm)

	server := http.Server{
		Addr:         ":9000",
		Handler:      hand,
		IdleTimeout:  123 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go errorHandler.Hub.Run()
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
