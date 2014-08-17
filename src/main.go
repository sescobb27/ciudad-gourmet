package main

import (
        "encoding/json"
        "flag"
        "fmt"
        "io/ioutil"
        "models"
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

func Index_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "text/html")
        file, err := ioutil.ReadFile("resources/index.html")
        if err != nil {
                panic(err)
        }
        res.Write(file)
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
        categories := models.GetCategories()
        json_categories, err := json.Marshal(categories)

        if err != nil {
                panic(err)
        }

        fmt.Println(string(json_categories))

        res.Write(json_categories)
}

func Locations_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
        locations := models.GetLocations()
        json_locations, err := json.Marshal(locations)

        if err != nil {
                panic(err)
        }

        fmt.Println(string(json_locations))

        res.Write(json_locations)
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
