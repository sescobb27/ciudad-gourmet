package models

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGetCategories(t *testing.T) {
    categories, err := GetCategories()
    assert.NoError(t, err)
    assert.NotEmpty(t, categories)
}
