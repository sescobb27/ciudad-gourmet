package models

import (
        . "github.com/sescobb27/ciudad-gourmet/db"
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

func seedUsers() {
        for i := 0; i < 4; i++ {
                now := time.Now()
                data := []string{passwords[i], now.String()}
                p := helpers.EncryptPassword(data)
                u := &User{
                        CreatedAt:    now,
                        Username:     usernames[i],
                        Email:        emails[i],
                        LastName:     last_names[i],
                        Name:         names[i],
                        PasswordHash: p,
                        Rate:         0.0}
                u.Create()
        }
}

func rollbackUsers(t *testing.T) {
        db, err := StablishConnection()
        if err != nil {
                t.Fatal(err)
        }
        defer db.Close()

        _, err = db.Exec("truncate table users restart identity CASCADE")

        if err != nil {
                t.Fatal("Error truncating table users")
        }
}

var testFindByEmail = func(t *testing.T) {
        for _, email := range emails {
                u, err := FindByEmail(&email)
                if err != nil {
                        t.Fatal(err)
                }
                if u == nil {
                        t.Fatalf("User with email %s should exist", email)
                }
        }
}

var testFindByUsername = func(t *testing.T) {
        for _, uname := range usernames {
                u, err := FindByUsername(&uname)
                if err != nil {
                        t.Fatal(err)
                }
                if u == nil {
                        t.Fatalf("User with username %s should exist", uname)
                }
        }
}

var testFindAllUsers = func(t *testing.T) {
        users, err := FindAllUsers()

        if err != nil {
                t.Fatal(err)
        }

        if len(users) != 4 {
                t.Fatalf("Should be 4 users and were found %d", len(users))
        }
}

func TestUserQueries(t *testing.T) {
        seedUsers()

        testFindByEmail(t)
        testFindByUsername(t)
        testFindAllUsers(t)

        rollbackUsers(t)
}
