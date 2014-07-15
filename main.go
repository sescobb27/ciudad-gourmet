package main

import (
        "fmt"
        "net/http"
)

func log_handler(handler http.HandlerFunc) http.HandlerFunc {
        return func(res http.ResponseWriter, req *http.Request) {
                fmt.Printf("%v", req)
                handler.ServeHTTP(res, req)
        }
}

func main() {
        server := http.NewServeMux()
        server.Handle("/", log_handler(Index_Handler))
        server.Handle("/signin", log_handler(SignIn_Handler))
        server.Handle("/logout", log_handler(LogOut_Handler))
        server.Handle("/register", log_handler(Register_Handler))
        server.Handle("/categories", log_handler(Categories_Handler))
        server.Handle("/locations", log_handler(Locations_Handler))
        server.Handle("/products", log_handler(Products_Handler))
        server.Handle("/products/findby", log_handler(FindbyProduct_Handler))
        server.Handle("/purchase", log_handler(Purchase_Handler))
        server.Handle("/chefs", log_handler(Chefs_Handler))
        server.Handle("/chefs/addproduct", log_handler(AddProduct_Handler))
        server.Handle("/chefs/listproducts", log_handler(ListProducts_Handler))
        http.ListenAndServe(":3000", server)
}

func Index_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func SignIn_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func LogOut_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Register_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Categories_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Locations_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Products_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func FindbyProduct_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Purchase_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Chefs_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func AddProduct_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func ListProducts_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}
