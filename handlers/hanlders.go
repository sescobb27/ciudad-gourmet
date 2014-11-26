package handlers

import (
    "encoding/json"
    "fmt"
    "github.com/julienschmidt/httprouter"
    "github.com/sescobb27/ciudad-gourmet/models"
    "github.com/sescobb27/ciudad-gourmet/services/log"
    "github.com/sescobb27/ciudad-gourmet/services/session"
    "golang.org/x/crypto/bcrypt"
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
)

var (
    logFactory *log.LogFactory
)

func init() {
    var err error
    logFactory, err = log.NewLogFactory("./ciudad-gourmet.log")
    if err != nil {
        panic(err)
    }
}

func Index_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "text/html")
    file, err := ioutil.ReadFile("resources/index.html")
    if err != nil {
        logFactory.Error(err.Error())
        http.Error(res, err.Error(), http.StatusNotFound)
        return
    }
    res.Write(file)
}

func formatReq(req *http.Request) string {
    return fmt.Sprintf("%v", (*req))
}

func SignIn_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    username := req.PostFormValue("username")
    password := req.PostFormValue("password")
    logFactory.Info(formatReq(req))
    sessionStore, err := session.Manager.SessionStart(res, req)
    if err != nil {
        logFactory.Error(err.Error())
    } else {
        if userSession := sessionStore.Get("user"); userSession != nil {
            var user *models.User
            var u_password string
            user = userSession.(*models.User)
            u_password = user.PasswordHash
            err = bcrypt.CompareHashAndPassword(
                []byte(u_password),
                []byte(password),
            )
            if err == nil && username == user.Username {
                return
            }
        }
    }
    user, err := models.FindUserByUsername(&username)
    if err != nil {
        http.Error(res, err.Error(), http.StatusNotFound)
        logFactory.Error(err.Error())
        return
    }
    err = bcrypt.CompareHashAndPassword(
        []byte(user.PasswordHash),
        []byte(password),
    )
    if err != nil {
        http.Error(res, err.Error(), http.StatusNotFound)
        logFactory.Error(err.Error())
        return
    }
}

func SignUp_Handler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

    username := req.PostFormValue("username")
    email := req.PostFormValue("email")
    lastname := req.PostFormValue("lastname")
    name := req.PostFormValue("name")
    password := req.PostFormValue("password")
    timeNow := time.Now().Local()
    logFactory.Info(formatReq(req))

    passwordHash, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )

    if err != nil {
        http.Error(res, err.Error(), http.StatusBadRequest)
        logFactory.Error(err.Error())
        return
    }

    user := &models.User{
        CreatedAt:    timeNow,
        Username:     username,
        Email:        email,
        LastName:     lastname,
        Name:         name,
        PasswordHash: string(passwordHash),
        Rate:         0.0,
    }

    if user.IsValid() {
        err = user.Create()
        if err != nil {
            http.Error(res, err.Error(), http.StatusBadRequest)
            logFactory.Error(err.Error())
            return
        }
        var sessionStore session.SessionStore
        sessionStore, err = session.Manager.SessionStart(res, req)
        if err != nil {
            logFactory.Error(err.Error())
        } else {
            sessionStore.Set("user", user)
        }
    } else {
        json_err, err := json.Marshal(user.Errors)
        if err != nil {
            http.Error(res, err.Error(), http.StatusInternalServerError)
            logFactory.Error(err.Error())
            return
        }
        res.Header().Set("Content-Type", "application/json")
        res.Write(json_err)
        res.WriteHeader(http.StatusBadRequest)
    }
}

func SignOut_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
}

func Categories_Handler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    logFactory.Info(formatReq(req))

    res.Header().Set("Content-Type", "application/json")
    categories, err := models.GetCategories()
    if err != nil {
        logFactory.Error(err.Error())
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    json_categories, err := json.Marshal(categories)

    if err != nil {
        logFactory.Error(err.Error())
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    res.Write(json_categories)
}

func Locations_Handler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    logFactory.Info(formatReq(req))

    res.Header().Set("Content-Type", "application/json")
    locations, err := models.GetLocations()
    if err != nil {
        logFactory.Error(err.Error())
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    json_locations, err := json.Marshal(locations)

    if err != nil {
        logFactory.Error(err.Error())
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
        logFactory.Error(err.Error())
        http.Error(res, err.Error(), http.StatusNotFound)
        return
    }

    products_json, err := json.Marshal(products)
    if err != nil {
        logFactory.Error(err.Error())
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
