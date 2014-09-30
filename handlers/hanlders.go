package handlers

import (
        "encoding/json"
        "github.com/sescobb27/ciudad-gourmet/models"
        "io/ioutil"
        "net/http"
)

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

func SignOut_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func NewChef_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Categories_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
        categories := models.GetCategories()
        json_categories, err := json.Marshal(categories)

        if err != nil {
                panic(err)
        }

        res.Write(json_categories)
}

func Locations_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
        locations := models.GetLocations()
        json_locations, err := json.Marshal(locations)

        if err != nil {
                panic(err)
        }

        res.Write(json_locations)
}

func Products_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func FindProduct_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Purchase_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Chefs_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func ChefAddProduct_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}
