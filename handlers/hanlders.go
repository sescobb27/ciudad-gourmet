package handlers

import (
    "encoding/json"
    "fmt"
    "github.com/julienschmidt/httprouter"
    "github.com/sescobb27/ciudad-gourmet/models"
    "github.com/sescobb27/ciudad-gourmet/services/log"
    "io/ioutil"
    "net/http"
    "strconv"
)

func Index_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "text/html")
    file, err := ioutil.ReadFile("resources/index.html")
    if err != nil {
        log.Log.Error(err.Error())
        http.Error(res, err.Error(), http.StatusNotFound)
        return
    }
    res.Write(file)
}

func formatReq(req *http.Request) string {
    return fmt.Sprintf("%v", (*req))
}

func Categories_Handler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    log.Log.Info(formatReq(req))

    res.Header().Set("Content-Type", "application/json")
    categories, err := models.GetCategories()
    if err != nil {
        log.Log.Error(err.Error())
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    json_categories, err := json.Marshal(categories)

    if err != nil {
        log.Log.Error(err.Error())
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    res.Write(json_categories)
}

func Locations_Handler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    log.Log.Info(formatReq(req))

    res.Header().Set("Content-Type", "application/json")
    locations, err := models.GetLocations()
    if err != nil {
        log.Log.Error(err.Error())
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    json_locations, err := json.Marshal(locations)

    if err != nil {
        log.Log.Error(err.Error())
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    res.Write(json_locations)
}

func Products_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
}

func FindProducts_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
    err := req.ParseForm()
    if err != nil {
        http.Error(res, err.Error(), http.StatusBadRequest)
        return
    }

    // the request should be /products/find?find_by={VAR}&{VAR}={ITEM}
    findBy := req.Form.Get("find_by")
    if findBy == "" {
        http.Error(res, "There is no key to find", http.StatusBadRequest)
        return
    }

    filter := req.Form.Get(findBy)
    if filter == "" {
        http.Error(res, "There is no key to filter", http.StatusBadRequest)
        return
    }

    var products []*models.Product
    switch findBy {
    case "id":
        // strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
        id, err := strconv.ParseInt(filter, 10, 0)
        if err != nil {
            break
        }
        products, err = models.FindProductsById(id)
    case "location":
        products, err = models.FindProductsByLocation(filter)
    case "category":
        products, err = models.FindProductsByCategory(filter)
    case "product_name":
        products, err = models.FindProductsByName(filter)
    case "username":
        products, err = models.FindProductsByUserName(filter)
    }

    if err != nil {
        log.Log.Error(err.Error())
        http.Error(res, err.Error(), http.StatusNotFound)
        return
    }

    products_json, err := json.Marshal(products)
    if err != nil {
        log.Log.Error(err.Error())
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    res.Write(products_json)
}

func Purchase_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
}

func Chefs_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
}

func ChefAddProduct_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
}
