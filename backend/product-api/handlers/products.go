package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Shahriar-shudip/golang-microservies-tuitorial/product-api/data"
	"github.com/gorilla/mux"
)

//Products exported
type Products struct {
	l *log.Logger
}

//NewProducts exported
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {

}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	prod.AddProduct()
}

func (p *Products) ProductDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	prod := data.GetProduct(string(params["id"]))
	err := prod.ToJson(w)
	if err != nil {
		fmt.Println(err)
	}

}

type KeyProduct struct{}

func (p *Products) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJson(r.Body)

		if err != nil {
			http.Error(w, "unable to parse", http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		p.l.Println(ctx.Value(KeyProduct{}))
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})

}
