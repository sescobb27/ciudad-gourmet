package session

import (
    "crypto/rand"
    "encoding/hex"
    "net/http"
    "net/url"
    "time"
)

// Based on from github.com/astaxie/beego/session
type CookieManager struct {
    cookieName      string
    expiresOn       time.Time
    sessionIdLength int
    domain          string
    provider        Provider
    semaphore       chan signal
}

func NewCookieManager(cookieName string, expiresOn time.Time, sessionIdLength int, provider Provider) *CookieManager {
    cookie := &CookieManager{
        cookieName:      cookieName,
        expiresOn:       expiresOn,
        sessionIdLength: sessionIdLength,
        provider:        provider,
        semaphore:       make(chan signal, 1),
    }
    cookie.semaphore <- signal{}
    return cookie
}

// Start session. generate or read the session id from http request.
// if session id exists, return SessionStore with this id.
func (cManager *CookieManager) SessionStart(res http.ResponseWriter, req *http.Request) (SessionStore, error) {
    var sessionStore SessionStore
    <-cManager.semaphore
    cookie, err := req.Cookie(cManager.cookieName)
    if err != nil || cookie.Value == "" {
        session_id, err := cManager.genSessionId()
        if err != nil {
            return nil, err
        }
        sessionStore = cManager.provider.SessionStore(session_id)
        cookie = cManager.genCookie(session_id)
        http.SetCookie(res, cookie)
        req.AddCookie(cookie)
    } else {
        session_id, err := url.QueryUnescape(cookie.Value)
        if err != nil {
            return nil, err
        }
        sessionStore = cManager.provider.SessionRead(session_id)
        if sessionStore == nil {
            sessionStore = cManager.provider.SessionStore(session_id)
            cookie = cManager.genCookie(session_id)
            http.SetCookie(res, cookie)
            req.AddCookie(cookie)
        }
    }
    cManager.semaphore <- signal{}
    return sessionStore, nil
}

func (cManager *CookieManager) genSessionId() (string, error) {
    b := make([]byte, cManager.sessionIdLength)
    n, err := rand.Read(b)
    if n != len(b) || err != nil {
        return "", err
    }
    return hex.EncodeToString(b), nil
}

// generate http cookie for the specified session id
func (cManager *CookieManager) genCookie(session_id string) *http.Cookie {
    return &http.Cookie{
        Name:     cManager.cookieName,
        Value:    url.QueryEscape(session_id),
        Path:     "/",
        HttpOnly: true,
        // Secure:   true,
        Domain:  cManager.domain,
        Expires: cManager.expiresOn,
    }
}

// Destroy session by its id in http request cookie.
func (cManager *CookieManager) SessionDestroy(res http.ResponseWriter, req *http.Request) {
    <-cManager.semaphore
    cookie, err := req.Cookie(cManager.cookieName)
    if err != nil || cookie.Value == "" {
        return
    } else {
        cManager.provider.SessionDestroy(cookie.Value)
        cookie.Expires = time.Now()
        http.SetCookie(res, cookie)
    }
    cManager.semaphore <- signal{}
}
