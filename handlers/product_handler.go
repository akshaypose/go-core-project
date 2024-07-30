package handlers

import (
	"encoding/json"
	"go-core-project/utils"
	"log"
	"net/http"
)

type Products struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type APIResponse struct {
	ProductList []Products `json:"products"`
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	body, err := utils.MakeCurlCall("https://dummyjson.com/products")
	if err != nil {
		http.Error(w, "Failed to make API call", http.StatusInternalServerError)
		return
	}

	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
		return
	}

	if data.ProductList == nil {
		jsonResponse(w, http.StatusOK, map[string]string{"Data": "No Products Found"})
		return
	}

	jsonResponse(w, http.StatusOK, data.ProductList)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	body, err := utils.MakeCurlCall("https://dummyjson.com/products/" + id)
	if err != nil {
		http.Error(w, "Failed to make API call", http.StatusInternalServerError)
		return
	}

	var data Products
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, http.StatusOK, data)
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}
