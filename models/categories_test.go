package models

import (
        "github.com/stretchr/testify/assert"
        "testing"
)

func TestGetCategories(t *testing.T) {
        categoryService := CategoryMock{}
        categories := categoryService.GetCategories()
        assert.NotEmpty(t, categories)
}
