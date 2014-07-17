package db

import (
        "../models"
)

type User models.User

func (u *User) Create() {
        // query := `INSERT INTO users(
        //     id, created_at, email, lastname, name, password_hash, rate, username)
        //     VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
}

func (u *User) FindByEmail() {
        // query := `SELECT * FROM users  AS u
        //     WHERE LOWER(u.email) = LOWER($1) LIMIT 1`
}

func (u *User) FindByUsername() {
        // query := `SELECT * FROM users AS u
        //     WHERE LOWER(u.username) = LOWER($1) LIMIT 1`
}

func (u *User) FindByProductId() {
        // query := `SELECT u.id, u.name, u.username, u.email
        //     FROM products AS p
        //     INNER JOIN users AS u on (p.chef_id = u.id)
        //     WHERE p.id = $1`
}
