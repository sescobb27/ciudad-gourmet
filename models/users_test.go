package models

import (
        "errors"
        "github.com/sescobb27/ciudad-gourmet/helpers"
        "github.com/stretchr/testify/assert"
        "testing"
        "time"
)

var (
        names      = []string{"Simon", "Edgardo", "Juan", "Camilo"}
        last_names = []string{"Escobar", "Sierra", "Norenia", "Mejia"}
        usernames  = []string{"sescob", "easierra", "jknore", "jcmejia"}
        emails     = []string{"sescob@gmail.com", "easierra@gmail.com",
                "jknore@gmail.com", "jcmejia@gmail.com"}
        passwords = []string{"qwerty", "123456", "AeIoU!@",
                "S3CUR3P455W0RD!\"#$%&/()="}
)

type MockUser User

func Stub_FindUserByEmail(email *string) (*MockUser, error) {
        mock_users := seedUsers()
        for _, user := range mock_users {
                if user.Email == (*email) {
                        return user, nil
                }
        }
        return nil, errors.New("No User With That Email")
}

func Stub_FindUserByUsername(username *string) (*MockUser, error) {
        mock_users := seedUsers()
        for _, user := range mock_users {
                if user.Username == (*username) {
                        return user, nil
                }
        }
        return nil, errors.New("No User With That Username")
}

func Stub_FindAllUsers() ([]*MockUser, error) {
        return seedUsers(), nil
}

func (u *MockUser) Stub_FindUserByProductId() {
}

func seedUsers() []*MockUser {
        mock_users := make([]*MockUser, 0, 4)
        for i := 0; i < 4; i++ {
                now := time.Now()
                data := []string{passwords[i], now.String()}
                p := helpers.EncryptPassword(data)
                u := &MockUser{
                        CreatedAt:    now,
                        Username:     usernames[i],
                        Email:        emails[i],
                        LastName:     last_names[i],
                        Name:         names[i],
                        PasswordHash: p,
                        Rate:         0.0}
                mock_users = append(mock_users, u)
        }
        return mock_users
}

func TestFindUserByEmail(t *testing.T) {
        t.Parallel()
        for _, email := range emails {
                u, err := Stub_FindUserByEmail(&email)
                assert.NoError(t, err)
                assert.NotNil(t, u, "User with email "+email+"should exist")
                assert.Equal(t, email, u.Email)
        }
}

func TestFindUserByUsername(t *testing.T) {
        t.Parallel()
        for _, uname := range usernames {
                u, err := Stub_FindUserByUsername(&uname)
                assert.NoError(t, err)
                assert.NotNil(t, u, "User with username %s should exist", uname)
                assert.Equal(t, uname, u.Username)
        }
}

func TestFindAllUsers(t *testing.T) {
        t.Parallel()
        users, err := Stub_FindAllUsers()
        assert.NoError(t, err)
        assert.NotEmpty(t, users)
        assert.Equal(t, 4, len(users))
}
