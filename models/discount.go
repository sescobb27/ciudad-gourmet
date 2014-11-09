package models

import (
    "time"
)

type Discount struct {
    Id              int64     `json:"id"`
    CreatedAt       time.Time `json:"created_at,omitempty"`
    Title           string    `json:"title"`
    Description     string    `json:"description"`
    DiscountPercent float32   `json:"discount_percent"`
    FinishAt        time.Time `json:"finish_at"`
    Finished        bool      `json:"finished"`
    Owner           *User     `json:"user"`
}
