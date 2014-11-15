package session

import (
    "github.com/sescobb27/ciudad-gourmet/models"
    "github.com/stretchr/testify/assert"
    "sync"
    "testing"
)

func TestMemoryProvider_Store_and_Read(t *testing.T) {
    t.Parallel()
    var wg sync.WaitGroup
    users := makeUsers()
    provider := NewSessionProvider()
    for _, u := range users {
        wg.Add(1)
        go func(u *models.User) {
            provider.SessionStore(u.Username)
            sessionStore := provider.SessionRead(u.Username)
            assert.NotNil(t, sessionStore)
            wg.Done()
        }(u)
    }
    wg.Wait()
}

func TestMemoryProvider_Store_and_Delete(t *testing.T) {
    t.Parallel()
    var wg sync.WaitGroup
    users := makeUsers()
    provider := NewSessionProvider()
    for _, u := range users {
        wg.Add(1)
        go func(u *models.User) {
            provider.SessionStore(u.Username)
            provider.SessionDestroy(u.Username)
            sessionStore := provider.SessionRead(u.Username)
            assert.Nil(t, sessionStore)
            wg.Done()
        }(u)
    }
    wg.Wait()
}
