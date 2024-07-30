package routes

import (
	"go-core-project/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", handlers.GetAllProducts)
	mux.HandleFunc("GET /product/{id}", handlers.GetProductByID)
	return mux
}
