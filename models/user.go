package models

import (
        . "db"
        "errors"
        "helpers"
        "log"
        "strings"
        "time"
)

type User struct {
        Id        int64
        CreatedAt time.Time
        Username, Email,
        LastName, Name,
        PasswordHash string
        Rate      float32
        Errors    []string
        Discounts []*Discount
        Products  []*Product
        Purchases []*Purchase
        Locations []*Location
}

func (u *User) Create() {
        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `INSERT INTO users(
            id, created_at, email, lastname, name, password_hash, rate, username)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

        _, err = db.Exec(query,
                u.Id,
                u.CreatedAt,
                u.Email,
                u.LastName,
                u.Name,
                u.PasswordHash,
                u.Rate,
                u.Username)

        if err != nil {
                log.Fatal(err)
                panic(err)
        }
}

func FindByEmail(email *string) (*User, error) {
        err := errors.New("Correo Invalido")
        if email == nil || len(*email) == 0 {
                return nil, err
        }

        (*email) = strings.ToLower(*email)
        if !helpers.EmailValidator(*email) {
                return nil, err
        }

        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `SELECT id, email, username, name, lastname, password_hash
            FROM users  AS u
            WHERE LOWER(u.email) = LOWER($1) LIMIT 1`

        user_row := db.QueryRow(query, email)
        if user_row == nil {
                return nil, errors.New("No User With That Email")
        }

        user := new(User)
        user_row.Scan(&user.Id,
                &user.Email,
                &user.Username,
                &user.Name,
                &user.LastName,
                &user.PasswordHash)
        return user, nil
}

func FindByUsername(username *string) (*User, error) {
        if username == nil || len(*username) == 0 ||
                !helpers.UniqueNamesValidator(*username) {
                return nil, errors.New("Correo Invalido")
        }
        (*username) = strings.ToLower(*username)

        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `SELECT id, email, username, name, lastname, password_hash
            FROM users AS u
            WHERE LOWER(u.username) = LOWER($1) LIMIT 1`

        user_row := db.QueryRow(query, username)
        if user_row == nil {
                return nil, errors.New("No User With That Username")
        }

        user := new(User)
        user_row.Scan(&user.Id,
                &user.Email,
                &user.Username,
                &user.Name,
                &user.LastName,
                &user.PasswordHash)
        return user, nil
}

func (u *User) FindByProductId() {
        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()
        // query := `SELECT u.id, u.name, u.username, u.email
        //     FROM products AS p
        //     INNER JOIN users AS u on (p.chef_id = u.id)
        //     WHERE p.id = $1`
}
