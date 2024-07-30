package routes

import (
	"go-core-project/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", handlers.GetAllProducts)
	mux.HandleFunc("/product/{id}", handlers.GetProductByID)
	return mux
}
