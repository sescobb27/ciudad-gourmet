package models

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var (
    p_names = []string{
        "plato1",
        "plato2",
        "plato3",
        "plato4",
    }
)

func TestFindProductsByName(t *testing.T) {
    t.Parallel()
    for _, name := range p_names {
        products, err := FindProductsByName(name)
        assert.NoError(t, err)
        assert.NotEmpty(t, products)
    }
}

func TestFindProductsByUserName(t *testing.T) {
    t.Parallel()
    for _, username := range u_usernames {
        products, err := FindProductsByUserName(username)
        assert.NoError(t, err)
        assert.NotEmpty(t, products)
    }
}
