package models

import (
    "time"
)

type Discount struct {
    Id                 int64
    CreatedAt          time.Time
    Title, Description string
    DiscountPercent    float32
    FinishAt           time.Time
    Finished           bool
    Chef               *User
}
