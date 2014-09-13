package models

import (
        "errors"
        "github.com/sescobb27/ciudad-gourmet/helpers"
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
        for _, email := range emails {
                u, err := Stub_FindUserByEmail(&email)
                if err != nil {
                        t.Fatal(err)
                }
                if u == nil {
                        t.Fatalf("User with email %s should exist", email)
                }
        }
}

func TestFindUserByUsername(t *testing.T) {
        for _, uname := range usernames {
                u, err := Stub_FindUserByUsername(&uname)
                if err != nil {
                        t.Fatal(err)
                }
                if u == nil {
                        t.Fatalf("User with username %s should exist", uname)
                }
        }
}

func TestFindAllUsers(t *testing.T) {
        users, err := Stub_FindAllUsers()

        if err != nil {
                t.Fatal(err)
        }

        if len(users) != 4 {
                t.Fatalf("Should be 4 users and were found %d", len(users))
        }
}
