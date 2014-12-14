package auth

import (
    "github.com/sescobb27/ciudad-gourmet/models"
    "github.com/stretchr/testify/assert"
    "net/http"
    "sync"
    "testing"
    "time"
)

var (
    u_usernames = []string{
        "sescob",
        "easierra",
        "jknore",
        "jcmejia",
    }
    u_emails = []string{
        "sescob@gmail.com",
        "easierra@gmail.com",
        "jknore@gmail.com",
        "jcmejia@gmail.com",
    }
)

func makeUsers() []*models.User {
    users := []*models.User{}
    for i := 0; i < 10; i++ {
        u := &models.User{
            Id:       int64(i),
            Username: u_usernames[i%4],
            Email:    u_emails[i%4],
        }
        users = append(users, u)
    }
    return users
}

func makeRequest() []*http.Request {
    reqs := make([]*http.Request, 0, 10)
    methods := []string{
        "GET",
        "POST",
    }
    urls := []string{
        "/",
        "/signup",
    }
    var err error
    var req *http.Request
    for i := 0; i < 10; i++ {
        req, err = http.NewRequest(
            methods[i%2],
            urls[i%2],
            nil,
        )
        if err != nil {
            panic(err)
        }
        reqs = append(reqs, req)
    }
    return reqs
}

func TestMakeToken(t *testing.T) {
    t.Parallel()
    reqs := makeRequest()
    var wg sync.WaitGroup
    for i, user := range makeUsers() {
        wg.Add(1)
        go func(u *models.User, req *http.Request) {
            token, err := MakeToken(u, time.Now().AddDate(1, 0, 0))
            assert.NoError(t, err)
            req.Header.Set("Authorization", "BEARER "+token)
            userSession, err := GetUserFromToken(req)
            assert.NoError(t, err)
            assert.Equal(t, u, userSession)
            wg.Done()
        }(user, reqs[i])
    }
    wg.Wait()
}
