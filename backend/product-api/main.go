package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Shahriar-shudip/golang-microservies-tuitorial/product-api/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	l := log.New(os.Stdout, "api", log.LstdFlags)

	//handlers
	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()
	getRoutes := sm.Methods("GET").Subrouter()
	getRoutes.HandleFunc("/products", ph.GetProducts)
	getRoutes.HandleFunc("/product/{id:[a-zA-Z0-9]+}", ph.ProductDetails)

	postRoutes := sm.Methods(http.MethodPost).Subrouter()
	postRoutes.HandleFunc("/addproduct", ph.AddProduct)
	postRoutes.Use(ph.Middleware)
	hand := cors.Default().Handler(sm)

	server := http.Server{
		Addr:         ":9090",
		Handler:      hand,
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
