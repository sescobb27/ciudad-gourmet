package session

import (
    "github.com/sescobb27/ciudad-gourmet/models"
    "github.com/stretchr/testify/assert"
    "sync"
    "testing"
)

var (
    u_names = []string{
        "Simon",
        "Edgardo",
        "Juan",
        "Camilo",
    }
    u_last_names = []string{
        "Escobar",
        "Sierra",
        "Norenia",
        "Mejia",
    }
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
    u_passwords = []string{
        "qwerty",
        "123456",
        "AeIoU!@",
        "S3CUR3P455W0RD!\"#$%&/()=",
    }
)

func makeUsers() []*models.User {
    users := []*models.User{}
    for i := 0; i < 50; i++ {
        u := &models.User{
            Username:     u_usernames[i%4] + string(i),
            Email:        u_emails[i%4] + string(i),
            LastName:     u_last_names[i%4] + string(i),
            Name:         u_names[i%4] + string(i),
            PasswordHash: u_passwords[i%4] + string(i),
            Rate:         0.0,
        }
        users = append(users, u)
    }
    return users
}

func TestMemorySession_Set_and_Get(t *testing.T) {
    t.Parallel()
    users := makeUsers()
    mStore := NewMemorySessionStore("memSession")
    var wg sync.WaitGroup
    for _, u := range users {
        wg.Add(1)
        go func(u *models.User) {
            mStore.Set(u.Username, u)
            userSession := mStore.Get(u.Username)
            assert.Equal(t, u, userSession)
            wg.Done()
        }(u)
    }
    wg.Wait()
}

func TestMemorySession_Set_and_Delete(t *testing.T) {
    t.Parallel()
    users := makeUsers()
    mStore := NewMemorySessionStore("memSession")
    var wg sync.WaitGroup
    for _, u := range users {
        wg.Add(1)
        go func(u *models.User) {
            mStore.Set(u.Username, u)
            mStore.Delete(u.Username)
            assert.Nil(t, mStore.Get(u.Username))
            wg.Done()
        }(u)
    }
    wg.Wait()
}
