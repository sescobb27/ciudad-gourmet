package models

type Purchase struct {
    Id              int64
    TotalPrice      float64
    Chef, Purchaser *User
}
