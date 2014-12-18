package handlers

import (
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "github.com/sescobb27/ciudad-gourmet/models"
    "github.com/sescobb27/ciudad-gourmet/services/log"
    "github.com/sescobb27/ciudad-gourmet/services/session"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "time"
)

func SignIn_Handler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    username := req.PostFormValue("username")
    password := req.PostFormValue("password")
    log.Log.Info(formatReq(req))
    sessionStore, err := session.Manager.SessionStart(res, req)
    if err != nil {
        log.Log.Error(err.Error())
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
        log.Log.Error(err.Error())
        return
    }
    err = bcrypt.CompareHashAndPassword(
        []byte(user.PasswordHash),
        []byte(password),
    )
    if err != nil {
        http.Error(res, err.Error(), http.StatusNotFound)
        log.Log.Error(err.Error())
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
    log.Log.Info(formatReq(req))

    passwordHash, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )

    if err != nil {
        http.Error(res, err.Error(), http.StatusBadRequest)
        log.Log.Error(err.Error())
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
            log.Log.Error(err.Error())
            return
        }
        var sessionStore session.SessionStore
        sessionStore, err = session.Manager.SessionStart(res, req)
        if err != nil {
            log.Log.Error(err.Error())
        } else {
            sessionStore.Set("user", user)
        }
    } else {
        json_err, err := json.Marshal(user.Errors)
        if err != nil {
            http.Error(res, err.Error(), http.StatusInternalServerError)
            log.Log.Error(err.Error())
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
