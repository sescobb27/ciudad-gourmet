package models

import (
        "github.com/stretchr/testify/assert"
        "testing"
)

var (
        u_usernames = []string{"sescob", "easierra", "jknore", "jcmejia"}
        u_emails    = []string{"sescob@gmail.com", "easierra@gmail.com",
                "jknore@gmail.com", "jcmejia@gmail.com"}
)

func TestFindUserByEmail(t *testing.T) {
        t.Parallel()
        for _, email := range u_emails {
                u, err := FindUserByEmail(&email)
                assert.NoError(t, err)
                assert.NotNil(t, u, "User with email "+email+"should exist")
                assert.Equal(t, email, u.Email)
        }
}

func TestFindUserByUsername(t *testing.T) {
        t.Parallel()
        for _, uname := range u_usernames {
                u, err := FindUserByUsername(&uname)
                assert.NoError(t, err)
                assert.NotNil(t, u, "User with username %s should exist", uname)
                assert.Equal(t, uname, u.Username)
        }
}

func TestFindAllUsers(t *testing.T) {
        t.Parallel()
        users, err := FindAllUsers()
        assert.NoError(t, err)
        assert.NotEmpty(t, users)
        assert.Equal(t, 4, len(users))
}
