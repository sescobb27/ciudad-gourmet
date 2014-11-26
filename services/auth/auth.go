package auth

import (
    "errors"
    "github.com/dgrijalva/jwt-go"
    "github.com/sescobb27/ciudad-gourmet/models"
    "github.com/sescobb27/ciudad-gourmet/services/session"
    "io/ioutil"
    "net/http"
    "time"
)

const (
    privateKeyPath = "../../cg.rsa"
    publicKeyPath  = "../../cg.rsa.pub"
)

var (
    privateKey   []byte
    publicKey    []byte
    memProvider  *session.MemoryProvider
    InvalidToken = errors.New("Invalid Token")
)

func init() {
    var err error
    privateKey, err = ioutil.ReadFile(privateKeyPath)
    if err != nil {
        panic(err)
    }
    publicKey, err = ioutil.ReadFile(publicKeyPath)
    if err != nil {
        panic(err)
    }
    memProvider = session.NewSessionProvider()
}

// MakeToken creates a new RS256 signed token. it sets the claims map
func MakeToken(user *models.User) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("RS256"))
    token.Claims["id"] = user.Id
    token.Claims["username"] = user.Username
    token.Claims["email"] = user.Email
    token.Claims["name"] = user.Name
    token.Claims["lastname"] = user.LastName
    // see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
    token.Claims["exp"] = time.Now().AddDate(1, 0, 0)
    tokenString, err := token.SignedString(privateKey)
    if err != nil {
        return "", err
    }
    sessionStore := memProvider.SessionStore(tokenString)
    sessionStore.Set("user", user)
    return tokenString, nil
}

// VerifyHeader checks for the exsistence of a Token in the Request Header
func GetUserFromToken(req *http.Request) (*models.User, error) {
    token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
        return publicKey, nil
    })

    if err != nil {
        return nil, err
    }

    if token.Valid {
        tokenString, err := token.SignedString(privateKey)
        if err != nil {
            return nil, err
        }
        sessionStore := memProvider.SessionRead(tokenString)
        if sessionStore == nil {
            return nil, InvalidToken
        }
        return sessionStore.Get("user").(*models.User), nil
    } else {
        return nil, InvalidToken
    }
}
