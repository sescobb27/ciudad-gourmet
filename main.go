package main

import (
        "flag"
        . "github.com/sescobb27/ciudad-gourmet/handlers"
        "net/http"
        "os"
)

func main() {
        seed := flag.Bool("seed", false, "Seed the database")
        restore := flag.Bool("restore", false, "Restore the database")
        flag.Parse()

        if *seed {
                Seed()
                os.Exit(0)
        } else if *restore {
                Restore()
                os.Exit(0)
        }

        server := http.NewServeMux()

        server.Handle("/", Get(Index_Handler))
        server.Handle("/categories", Get(Categories_Handler))
        server.Handle("/locations", Get(Locations_Handler))
        server.Handle("/products", Get(Products_Handler))
        server.Handle("/products/find", Get(FindProducts_Handler))

        server.Handle("/signin", Post(SignIn_Handler))
        server.Handle("/signout", Post(SignOut_Handler))
        server.Handle("/signup", Post(SignUp_Handler))
        server.Handle("/purchase", Post(Purchase_Handler))

        server.Handle("/chefs", Get(Chefs_Handler))
        server.Handle("/chefs/product/add", Post(ChefAddProduct_Handler))

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
