package main

import (
        "crypto/sha256"
        "fmt"
        "io"
        "models"
        "time"
)

var (
        categories = []string{"Pasta", "Carne", "Pollo", "Ensalada", "Desayuno",
                "Almuerzo", "Postre", "Sopa", "Vegetariana", "Menu Infantil",
                "Comida Rapida", "Almuerzo para 2", "Desayuno para 2", "Comida para 2",
                "Ensalada de Frutas", "Gourmet"}
        descriptions = []string{"Ricas Pasta", "Ricas Carnes", "Rico Pollo",
                "Ricas Ensaladas", "Ricos Desayunos", "Ricos Almuerzos", "Ricos Postres",
                "Ricas Sopas", "Rica Comida Vegetariana", "Ricos Menu Infantil",
                "Rica Comida Rapida", "Ricos Almuerzo para 2", "Ricos Desayuno para 2",
                "Ricas Comidas para 2", "Ricas Ensaladas de Frutas", "Ricas Comida Gourmet"}
        locations = []string{"Poblado", "Laureles", "Envigado", "Caldas",
                "Sabaneta", "Colores", "Estadio", "Calazans", "Bello", "Boston",
                "Prado Centro", "Itagui", "Belen", "Guayabal"}
)

func encryptPassword(data []string) string {
        hash := sha256.New()
        for _, v := range data {
                io.WriteString(hash, v)
        }
        return fmt.Sprintf("%x", hash.Sum(nil))
}

func insert_Users() {
        names := []string{"Simon", "Edgardo", "Juan", "Camilo"}
        last_names := []string{"Escobar", "Sierra", "Norenia", "Mejia"}
        usernames := []string{"sescob", "easierra", "jknore", "jcmejia"}
        emails := []string{"sescob@gmail.com", "easierra@gmail.com",
                "jknore@gmail.com", "jcmejia@gmail.com"}
        passwords := []string{"qwerty", "123456", "AeIoU!@",
                "S3CUR3P455W0RD!\"#$%&/()="}

        for i := 0; i < 4; i++ {
                now := time.Now()
                data := []string{passwords[i], now.String()}
                p := encryptPassword(data)
                u := &models.User{Id: int64(i),
                        CreatedAt:    now,
                        Username:     usernames[i],
                        Email:        emails[i],
                        LastName:     last_names[i],
                        Name:         names[i],
                        PasswordHash: p,
                        Rate:         0.0}
                u.Create()
        }
}

// func insertProducts(p *Product) {
//         query := `insert into products
//       (id, created_at, description, image, name, price, rate, chef_id)
//   values
//       ($1,$2,$3,$4,$5,$6,$7,$8)`
//         db, err := stablishConnection()
//         assertNoError(err)
//         defer db.Close()

//         _, err = db.Exec(query, p.Id, p.Created_at, p.Description, p.Image, p.Name,
//                 p.Price, p.Rate, p.Chef_id)
//         assertNoError(err)
// }

// func makeProducts() {
//         names := []string{"plato1", "plato2", "plato3", "plato4"}
//         descriptions := []string{"Descripcion1", "Descripcion2",
//                 "Descripcion3", "Descripcion4"}
//         prices := []float64{18500.0, 12300.0, 5000.0, 7300.0}
//         images := []string{"images/default.png",
//                 "images/default.png", "images/default.png", "images/default.png"}
//         rates := []float32{1.9, 2.5, 3.2, 4.8}

//         for i := 0; i < 4; i++ {
//                 p := &Product{i, time.Now(), descriptions[i], names[i], prices[i], rates[i],
//                         images[i], i}
//                 insertProducts(p)
//         }
// }

func insert_Categories() {
        for i, category := range categories {
                c := &models.Category{int8(i), category, descriptions[i]}
                c.Create()
        }
}

func insert_Locations() {
        for i, location := range locations {
                l := &models.Location{int8(i), location}
                l.Create()
        }
}

// func insertProductsLocations() {
//         db, err := stablishConnection()
//         assertNoError(err)
//         defer db.Close()

//         query := `insert into products_locations (id, location_id, product_id)
//   values ($1,$2, $3)`
//         for i, _ := range locations {
//                 _, err = db.Exec(query, i, i, i%4)
//                 assertNoError(err)
//         }
// }

// func insertProductsCategories() {
//         db, err := stablishConnection()
//         assertNoError(err)
//         defer db.Close()

//         query := `insert into products_categories (id, category_id, product_id)
//   values ($1,$2, $3)`
//         for i, _ := range categories {
//                 _, err = db.Exec(query, i, i, i%4)
//                 assertNoError(err)
//         }
// }

// func restore() {
//         db, err := stablishConnection()
//         assertNoError(err)
//         defer db.Close()
//         tables := []string{"users", "categories", "locations", "products"}
//         for _, t := range tables {
//                 _, err = db.Exec("truncate table " + t + " restart identity CASCADE")
//                 assertNoError(err)
//         }
// }

func Run() {
        // restore()
        insert_Users()
        insert_Categories()
        insert_Locations()
        // makeProducts()
        // insertProductsCategories()
        // insertProductsLocations()
}
