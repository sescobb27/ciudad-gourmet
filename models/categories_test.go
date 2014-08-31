package models

import (
        "testing"
)

func TestGetCategories(t *testing.T) {
        categories := GetCategories()
        if len(categories) == 0 {
                t.Fatal("Categories should not be empty")
        }
}
