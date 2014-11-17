package session

import (
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "strings"
    "sync"
    "testing"
    "time"
)

// 2 requests at the same time should be treated as different so they should
// return 2 different SessionStore values
func TestSessionStart_2DifferentRequest(t *testing.T) {
    t.Parallel()
    // First simulated request and response
    recorder1 := httptest.NewRecorder()
    req1, err := http.NewRequest(
        "POST",
        "/signin",
        strings.NewReader("username=sescob&password=qwerty"),
    )

    // Second simulated request and response
    recorder2 := httptest.NewRecorder()
    req2, err := http.NewRequest(
        "POST",
        "/signin",
        strings.NewReader("username=sescob&password=qwerty"),
    )
    assert.NoError(t, err)
    expiresOn := time.Now().AddDate(1, 0, 0) // expires 1 year after this one
    provider := NewSessionProvider()
    cookie := NewCookieManager("cookieTest", expiresOn, 16, provider)

    var (
        ws            sync.WaitGroup
        sessionStore  SessionStore
        sessionStore2 SessionStore
    )
    ws.Add(1)
    go func(tmpCookie *CookieManager, res http.ResponseWriter, r *http.Request) {
        var err1 error
        sessionStore, err1 = tmpCookie.SessionStart(res, r)
        assert.NoError(t, err1)
        assert.NotNil(t, sessionStore)
        cookieStr := res.Header().Get("Set-Cookie")
        assert.True(t, strings.Contains(cookieStr, sessionStore.SessionID()))
        assert.NotNil(t, res.Header().Get("Set-Cookie"))

        ws.Done()
    }(cookie, recorder1, req1)
    ws.Add(1)
    go func(tmpCookie *CookieManager, res http.ResponseWriter, r *http.Request) {
        var err2 error
        sessionStore2, err2 = tmpCookie.SessionStart(res, r)
        assert.NoError(t, err2)
        assert.NotNil(t, sessionStore2)
        cookieStr := res.Header().Get("Set-Cookie")
        assert.True(t, strings.Contains(cookieStr, sessionStore2.SessionID()))
        assert.NotNil(t, res.Header().Get("Set-Cookie"))

        ws.Done()
    }(cookie, recorder2, req2)
    ws.Wait()
    assert.NotNil(t, sessionStore)
    assert.NotNil(t, sessionStore2)
    assert.NotEqual(t, sessionStore, sessionStore2)
}

// 2 requests at the same time should be treated as different so the second one
// can't read the session deleted by the previous one
func TestSessionDestroy_2DifferentRequest(t *testing.T) {
    t.Parallel()
    recorder := httptest.NewRecorder()
    req, err := http.NewRequest(
        "POST",
        "/signin",
        strings.NewReader("username=sescob&password=qwerty"),
    )
    assert.NoError(t, err)
    expiresOn := time.Now().AddDate(1, 0, 0) // expires 1 year after this one
    provider := NewSessionProvider()
    cookie := NewCookieManager("cookieTest", expiresOn, 16, provider)

    var (
        sessionStore SessionStore
        ws           sync.WaitGroup
    )
    sessionStore, err = cookie.SessionStart(recorder, req)
    assert.NoError(t, err)

    ws.Add(1)
    done := make(chan signal)
    go func(rw http.ResponseWriter, r *http.Request) {
        cookie.SessionDestroy(rw, r)
        done <- signal{}
        assert.Nil(t, cookie.provider.SessionRead(sessionStore.SessionID()))
        ws.Done()
    }(recorder, req)

    ws.Add(1)
    go func(rw http.ResponseWriter, r *http.Request) {
        <-done
        assert.Nil(t, cookie.provider.SessionRead(sessionStore.SessionID()))
        ws.Done()
    }(recorder, req)
    ws.Wait()
}
