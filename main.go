package main

import (
        "flag"
        "fmt"
        "github.com/sescobb27/ciudad-gourmet/handlers"
        "net/http"
        "os"
)

func log_handler(handler http.HandlerFunc) http.HandlerFunc {
        return func(res http.ResponseWriter, req *http.Request) {
                fmt.Printf("%v\n\n", req)
                handler.ServeHTTP(res, req)
        }
}

func main() {
        seed := flag.Bool("seed", false, "Seed the database")
        flag.Parse()

        if *seed {
                Seed()
                os.Exit(0)
        }

        server := http.NewServeMux()

        server.Handle("/",
                log_handler(handlers.Index_Handler))

        server.Handle("/signin",
                log_handler(handlers.SignIn_Handler))

        server.Handle("/logout",
                log_handler(handlers.LogOut_Handler))

        server.Handle("/register",
                log_handler(handlers.Register_Handler))

        server.Handle("/categories",
                log_handler(handlers.Categories_Handler))

        server.Handle("/locations",
                log_handler(handlers.Locations_Handler))

        server.Handle("/products",
                log_handler(handlers.Products_Handler))

        server.Handle("/products/findby",
                log_handler(handlers.FindbyProduct_Handler))

        server.Handle("/purchase",
                log_handler(handlers.Purchase_Handler))

        server.Handle("/chefs",
                log_handler(handlers.Chefs_Handler))

        server.Handle("/chefs/addproduct",
                log_handler(handlers.AddProduct_Handler))

        server.Handle("/chefs/listproducts",
                log_handler(handlers.ListProducts_Handler))

        server.Handle("/images/",
                http.StripPrefix("/images/",
                        http.FileServer(http.Dir("resources/images"))))

        server.Handle("/css/",
                http.StripPrefix("/css/",
                        http.FileServer(http.Dir("resources/css"))))

        server.Handle("/js/",
                http.StripPrefix("/js/",
                        http.FileServer(http.Dir("resources/js"))))

        server.Handle("/catalog/",
                http.StripPrefix("/catalog/",
                        http.FileServer(http.Dir("resources/catalog"))))

        http.ListenAndServe(":3000", server)
}
