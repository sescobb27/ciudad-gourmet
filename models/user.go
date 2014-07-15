package models

import (
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
        Discounts []Discount
        Products  []Product
        Purchases []Purchase
        Locations []Location
}
