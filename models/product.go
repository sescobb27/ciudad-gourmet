package models

import (
        "time"
)

type Product struct {
        Id                       int64
        CreatedAt                time.Time
        Image, Description, Name string
        Price                    float64
        Rate                     float32
        Errors                   []string
        Chef                     User
        Categories               []Category
        Discounts                []Discount
}
