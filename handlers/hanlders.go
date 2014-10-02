package handlers

import (
        "encoding/json"
        "github.com/sescobb27/ciudad-gourmet/helpers"
        "github.com/sescobb27/ciudad-gourmet/models"
        "io/ioutil"
        "net/http"
        "strconv"
        "time"
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

func SignUp_Handler(res http.ResponseWriter, req *http.Request) {
        err := req.ParseForm()
        if err != nil {
                http.Error(res, err.Error(), http.StatusBadRequest)
                return
        }

        username := req.Form.Get("username")
        email := req.Form.Get("email")
        lastname := req.Form.Get("lastname")
        name := req.Form.Get("name")
        password := req.Form.Get("password")
        timeNow := time.Now().Local()
        dataToEncrypt := []string{timeNow.Format(time.RFC850), password}

        passwordHash := helpers.EncryptPassword(dataToEncrypt)

        user := &models.User{
                CreatedAt:    timeNow,
                Username:     username,
                Email:        email,
                LastName:     lastname,
                Name:         name,
                PasswordHash: passwordHash,
                Rate:         0.0,
        }

        err = user.Create()

        if err != nil {
                http.Error(res, err.Error(), http.StatusBadRequest)
        }
}

func SignOut_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func NewChef_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func Categories_Handler(categorieService models.CategoryService) http.HandlerFunc {
        return func(res http.ResponseWriter, req *http.Request) {
                res.Header().Set("Content-Type", "application/json")
                categories := []*models.Category{}
                categories = append(categories, categorieService.GetCategories()...)
                json_categories, err := json.Marshal(categories)

                if err != nil {
                        panic(err)
                }

                res.Write(json_categories)
        }
}

func Locations_Handler(locationService models.LocationService) http.HandlerFunc {
        return func(res http.ResponseWriter, req *http.Request) {
                res.Header().Set("Content-Type", "application/json")
                locations := []*models.Location{}
                locations = append(locations, locationService.GetLocations()...)
                json_locations, err := json.Marshal(locations)

                if err != nil {
                        panic(err)
                }

                res.Write(json_locations)
        }
}

func Products_Handler(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "application/json")
}

func FindProducts_Handler(res http.ResponseWriter, req *http.Request) {
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
                http.Error(res, err.Error(), http.StatusNotFound)
                return
        }

        products_json, err := json.Marshal(products)
        if err != nil {
                http.Error(res, err.Error(), http.StatusInternalServerError)
                return
        }

        res.Write(products_json)
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
