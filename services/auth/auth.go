package auth

import (
    "errors"
    "github.com/dgrijalva/jwt-go"
    "github.com/sescobb27/ciudad-gourmet/models"
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
}

// MakeToken creates a new RS256 signed token. it sets the claims map
func MakeToken(user *models.User, expirationTime time.Time) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("RS256"))
    token.Claims["id"] = user.Id
    token.Claims["username"] = user.Username
    token.Claims["email"] = user.Email
    // see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
    token.Claims["exp"] = expirationTime // time.Now().AddDate(1, 0, 0)
    tokenString, err := token.SignedString(privateKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// GetUserFromToken checks for the exsistence of a Token in the Request Header,
// then if there is no error and the token is valid we get the tokenString
// which is our in memory sessionId and try to retrieve us a user from there
func GetUserFromToken(req *http.Request) (*models.User, error) {
    token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
        return publicKey, nil
    })

    if err != nil {
        return nil, err
    }

    if token.Valid {
        user := &models.User{}
        user.Id = int64(token.Claims["id"].(float64))
        user.Username = token.Claims["username"].(string)
        user.Email = token.Claims["email"].(string)
        return user, nil
    } else {
        return nil, InvalidToken
    }
}
