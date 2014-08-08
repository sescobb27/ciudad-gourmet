package main

import (
        "models"
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
        insert_Categories()
        insert_Locations()
        // insertProductsCategories()
        // insertProductsLocations()
}
