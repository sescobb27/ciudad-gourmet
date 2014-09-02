package models

import (
        "testing"
        "time"
)

var (
        product_names = []string{"plato1", "plato2", "plato3", "plato4"}
        descriptions  = []string{"Descripcion1", "Descripcion2",
                "Descripcion3", "Descripcion4"}
        prices          = []float64{18500.0, 12300.0, 5000.0, 7300.0}
        image   string  = "images/default.png"
        rates           = []float32{1.9, 2.5, 3.2, 4.8}
)

func seedProducts() {
        seedUsers()
        for i := 0; i < 4; i++ {
                u := &User{Id: int64(i)}
                p := &Product{Id: int64(i),
                        CreatedAt:   time.Now(),
                        Image:       image,
                        Description: descriptions[i],
                        Name:        product_names[i],
                        Price:       prices[i],
                        Rate:        rates[i],
                        Chef:        u}
                p.Create()
        }
}

func rollbackProducts(t *testing.T) {
        rollbackUsers(t)
}

func TestFindByName(t *testing.T) {
        seedProducts()
        for _, name := range product_names {
                products, err := FindByName(name)
                if err != nil {
                        t.Fatal(err)
                }
                if products == nil || len(products) == 0 {
                        t.Fatalf("Error: Product %s shoud exist", name)
                }
        }
        rollbackProducts(t)
}