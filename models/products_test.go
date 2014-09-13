package models

import (
        "github.com/stretchr/testify/assert"
        "testing"
        "time"
        "unsafe"
)

var (
        product_names = []string{"plato1", "plato2", "plato3", "plato4"}
        descriptions  = []string{"Descripcion1", "Descripcion2",
                "Descripcion3", "Descripcion4"}
        prices          = []float64{18500.0, 12300.0, 5000.0, 7300.0}
        image   string  = "images/default.png"
        rates           = []float32{1.9, 2.5, 3.2, 4.8}
)

type MockProduct Product

func Stub_FindProductsByName(p_name string) ([]*MockProduct, error) {
        mock_products := seedProducts()
        result := make([]*MockProduct, 0, 4)
        for _, product := range mock_products {
                if product.Name == p_name {
                        result = append(result, product)
                }
        }
        return result, nil
}

func seedProducts() []*MockProduct {
        seedUsers()
        users, _ := Stub_FindAllUsers()
        mock_products := make([]*MockProduct, 0, 4)
        for i, user := range users {
                p := &MockProduct{CreatedAt: time.Now(),
                        Image:       image,
                        Description: descriptions[i],
                        Name:        product_names[i],
                        Price:       prices[i],
                        Rate:        rates[i],
                        Chef:        (*User)(unsafe.Pointer(&user))}
                mock_products = append(mock_products, p)
        }
        return mock_products
}

func TestFindProductsByName(t *testing.T) {
        for _, name := range product_names {
                products, err := Stub_FindProductsByName(name)
                assert.NoError(t, err)
                assert.NotEmpty(t, products)
        }
}
