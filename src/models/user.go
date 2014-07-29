package models

import (
        "db"
        "log"
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
        db, err := db.StablishConnection()
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

func (u *User) FindByEmail() {
        db, err := db.StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()
        // query := `SELECT * FROM users  AS u
        //     WHERE LOWER(u.email) = LOWER($1) LIMIT 1`
}

func (u *User) FindByUsername() {
        db, err := db.StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()
        // query := `SELECT * FROM users AS u
        //     WHERE LOWER(u.username) = LOWER($1) LIMIT 1`
}

func (u *User) FindByProductId() {
        db, err := db.StablishConnection()
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
