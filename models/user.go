package models

import (
    stdsql "database/sql"
    "errors"
    sql "github.com/sescobb27/ciudad-gourmet/db"
    "github.com/sescobb27/ciudad-gourmet/helpers"
    "strings"
    "time"
)

type User struct {
    Id           int64       `json:"id,omitempty"`
    CreatedAt    time.Time   `json:"created_at,omitempty"`
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
    existChan := make(chan bool)
    go func() {
        existChan <- UserExist(u.Username, u.Email)
    }()
    var err error
    if exist := <-existChan; !exist {
        query := `INSERT INTO users(
            created_at, email, lastname, name, password_hash, rate, username)
            VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
        err = sql.DB.QueryRow(
            query,
            u.CreatedAt.Format(time.RFC850),
            u.Email,
            u.LastName,
            u.Name,
            u.PasswordHash,
            u.Rate,
            u.Username,
        ).Scan(&u.Id)
    } else {
        err = errors.New("User already exist")
    }
    return err
}

func (u *User) IsValid() bool {
    if !helpers.EmailValidator(u.Email) {
        u.Errors = append(u.Errors, "Invalid email")
    }
    if !helpers.UniqueNamesValidator(u.Username) {
        u.Errors = append(u.Errors, "Invalid username")
    }
    if !helpers.UserNamesValidator(u.Name) {
        u.Errors = append(u.Errors, "Invalid name")
    }
    if !helpers.UserNamesValidator(u.LastName) {
        u.Errors = append(u.Errors, "Invalid last name")
    }
    return len(u.Errors) == 0
}

func UserExist(username, email string) bool {
    query := `SELECT id FROM users  AS u
            WHERE LOWER(u.email) = LOWER($1) OR
            LOWER(u.username) = LOWER($2) LIMIT 1`
    var id stdsql.NullInt64
    sql.DB.QueryRow(query, email, username).Scan(&id)
    return id.Int64 != 0
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

    query := `SELECT id, email, username, name, lastname, password_hash
            FROM users  AS u
            WHERE LOWER(u.email) = LOWER($1) LIMIT 1`

    user_row := sql.DB.QueryRow(query, email)
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

    query := `SELECT id, email, username, name, lastname, password_hash
            FROM users AS u
            WHERE LOWER(u.username) = LOWER($1) LIMIT 1`

    user_row := sql.DB.QueryRow(query, username)
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
    query := `SELECT id, email, username, name, lastname, rate FROM users`

    user_rows, err := sql.DB.Query(query)

    users := []*User{}
    if err != nil {
        return users, err
    }
    defer user_rows.Close()

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
    //     FROM products AS p
    //     INNER JOIN users AS u on (p.chef_id = u.id)
    //     WHERE p.id = $1`
}
