package models

import (
        "errors"
        . "github.com/sescobb27/ciudad-gourmet/db"
        "github.com/sescobb27/ciudad-gourmet/helpers"
        "log"
        "strings"
        "time"
)

type User struct {
        Id           int64       `json:"id,omitempty"`
        CreatedAt    time.Time   `json:"createdAt,omitempty"`
        Username     string      `json:"username"`
        Email        string      `json:"email"`
        LastName     string      `json:"lastname"`
        Name         string      `json:"name"`
        PasswordHash string      `json:"-"`
        Rate         float32     `json:"rate"`
        Errors       []string    `json:"errors"`
        Discounts    []*Discount `json:"discounts"`
        Products     []*Product  `json:"products"`
        Purchases    []*Purchase `json:"purchases"`
        Locations    []*Location `json:"locations"`
}

func (u *User) Create() error {
        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `INSERT INTO users(
            created_at, email, lastname, name, password_hash, rate, username)
            VALUES ($1, $2, $3, $4, $5, $6, $7)`

        _, err = db.Exec(
                query,
                u.CreatedAt.Format(time.RFC850),
                u.Email,
                u.LastName,
                u.Name,
                u.PasswordHash,
                u.Rate,
                u.Username,
        )

        return err
}

func FindUserByEmail(email *string) (*User, error) {
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

        user := &User{}
        user_row.Scan(
                &user.Id,
                &user.Email,
                &user.Username,
                &user.Name,
                &user.LastName,
                &user.PasswordHash,
        )
        return user, nil
}

func FindUserByUsername(username *string) (*User, error) {
        if username == nil || len(*username) == 0 ||
                !helpers.UniqueNamesValidator(*username) {
                return nil, errors.New("Nombre de Usuario Invalido")
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

        user := &User{}
        user_row.Scan(
                &user.Id,
                &user.Email,
                &user.Username,
                &user.Name,
                &user.LastName,
                &user.PasswordHash,
        )
        return user, nil
}

func FindAllUsers() ([]*User, error) {
        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `SELECT id, email, username, name, lastname, rate FROM users`

        user_rows, err := db.Query(query)

        if err != nil {
                return nil, err
        }

        users := []*User{}
        if user_rows == nil {
                return users, nil
        }

        for user_rows.Next() {
                user := &User{}
                user_rows.Scan(
                        &user.Id,
                        &user.Email,
                        &user.Username,
                        &user.Name,
                        &user.LastName,
                        &user.Rate,
                )
                users = append(users, user)
        }
        return users, nil
}

func (u *User) FindUserByProductId() {
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
