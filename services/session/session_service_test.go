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

func TestSessionStart(t *testing.T) {
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
        ws            sync.WaitGroup
        sessionStore  SessionStore
        sessionStore2 SessionStore
    )
    ws.Add(1)
    go func(cookie *CookieManager) {
        var err1 error
        sessionStore, err1 = cookie.SessionStart(recorder, req)
        assert.NoError(t, err1)
        assert.NotNil(t, sessionStore)
        assert.NotNil(t, recorder.HeaderMap.Get("Set-Cookie"))
        ws.Done()
    }(cookie)
    ws.Add(1)
    go func(cookie *CookieManager) {
        var err2 error
        sessionStore2, err2 = cookie.SessionStart(recorder, req)
        assert.NoError(t, err2)
        assert.NotNil(t, sessionStore2)
        assert.NotNil(t, recorder.HeaderMap.Get("Set-Cookie"))
        ws.Done()
    }(cookie)
    ws.Wait()
    assert.NotNil(t, sessionStore)
    assert.NotNil(t, sessionStore2)
    assert.Equal(t, sessionStore, sessionStore2)
}
