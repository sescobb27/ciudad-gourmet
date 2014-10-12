package main

import (
        "github.com/sescobb27/ciudad-gourmet/db"
        "github.com/sescobb27/ciudad-gourmet/helpers"
        "github.com/sescobb27/ciudad-gourmet/models"
        "time"
)

var (
        categoriesSeed = []string{
                "Pasta",
                "Carne",
                "Pollo",
                "Ensalada",
                "Desayuno",
                "Almuerzo",
                "Postre",
                "Sopa",
                "Vegetariana",
                "Menu Infantil",
                "Comida Rapida",
                "Almuerzo para 2",
                "Desayuno para 2",
                "Comida para 2",
                "Ensalada de Frutas",
                "Gourmet",
        }
        descriptionsSeed = []string{
                "Ricas Pasta",
                "Ricas Carnes",
                "Rico Pollo",
                "Ricas Ensaladas",
                "Ricos Desayunos",
                "Ricos Almuerzos",
                "Ricos Postres",
                "Ricas Sopas",
                "Rica Comida Vegetariana",
                "Ricos Menu Infantil",
                "Rica Comida Rapida",
                "Ricos Almuerzo para 2",
                "Ricos Desayuno para 2",
                "Ricas Comidas para 2",
                "Ricas Ensaladas de Frutas",
                "Ricas Comida Gourmet",
        }
        locationsSeed = []string{
                "Poblado",
                "Laureles",
                "Envigado",
                "Caldas",
                "Sabaneta",
                "Colores",
                "Estadio",
                "Calazans",
                "Bello",
                "Boston",
                "Prado Centro",
                "Itagui",
                "Belen",
                "Guayabal",
        }
)

func insert_Categories() []*models.Category {
        categories := models.GetCategories()
        if len(categories) == 0 {
                for i, category := range categoriesSeed {
                        c := &models.Category{
                                Name:        category,
                                Description: descriptionsSeed[i],
                        }
                        c.Create()
                        categories = append(categories, c)
                }
        }
        return categories
}

func insert_Locations() []*models.Location {
        locations := models.GetLocations()
        if len(locations) == 0 {
                for _, location := range locationsSeed {
                        l := &models.Location{Name: location}
                        l.Create()
                        locations = append(locations, l)
                }
        }
        return locations
}

var (
        u_names = []string{
                "Simon",
                "Edgardo",
                "Juan",
                "Camilo",
        }
        u_last_names = []string{
                "Escobar",
                "Sierra",
                "Norenia",
                "Mejia",
        }
        u_usernames = []string{
                "sescob",
                "easierra",
                "jknore",
                "jcmejia",
        }
        u_emails = []string{
                "sescob@gmail.com",
                "easierra@gmail.com",
                "jknore@gmail.com",
                "jcmejia@gmail.com",
        }
        u_passwords = []string{
                "qwerty",
                "123456",
                "AeIoU!@",
                "S3CUR3P455W0RD!\"#$%&/()=",
        }
)

func insert_Users() []*models.User {
        users, err := models.FindAllUsers()
        if err != nil {
                panic(err)
        }
        if len(users) == 0 {
                for i := 0; i < 4; i++ {
                        now := time.Now().Local()
                        data := []string{now.Format(time.RFC850), u_passwords[i]}
                        passwordHash := helpers.EncryptPassword(data)
                        u := &models.User{
                                CreatedAt:    now,
                                Username:     u_usernames[i],
                                Email:        u_emails[i],
                                LastName:     u_last_names[i],
                                Name:         u_names[i],
                                PasswordHash: passwordHash,
                                Rate:         0.0,
                        }
                        u.Create()
                        users = append(users, u)
                }
        }
        return users
}

var (
        p_names = []string{
                "plato1",
                "plato2",
                "plato3",
                "plato4",
        }
        p_descriptions = []string{
                "Descripcion1",
                "Descripcion2",
                "Descripcion3",
                "Descripcion4",
        }
        p_prices = []float64{
                18500.0,
                12300.0,
                5000.0,
                7300.0,
        }
        p_rates = []float32{
                1.9,
                2.5,
                3.2,
                4.8,
        }
        p_image string  = "images/default.png"
)

func insert_Products(users []*models.User, locations []*models.Location, categories []*models.Category) {
        for i, user := range users {
                p := &models.Product{
                        CreatedAt:   time.Now().Local(),
                        Image:       p_image,
                        Description: p_descriptions[i],
                        Name:        p_names[i],
                        Price:       p_prices[i],
                        Rate:        p_rates[i],
                        Chef:        user,
                }
                p.Create()
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

func Restore() {
        database, err := db.StablishConnection()
        if err != nil {
                panic(err)
        }

        tables := "users, categories, locations, products"
        _, err = database.Exec("TRUNCATE TABLE " + tables + " RESTART IDENTITY CASCADE")

        if err != nil {
                panic(err)
        }

        sequences := []string{
                "categories_id_sequence",
                "discounts_id_sequence",
                "locations_id_sequence",
                "payment_types_id_sequence",
                "products_id_sequence",
                "products_categories_id_sequence",
                "products_discounts_id_sequence",
                "products_payment_types_id_sequence",
                "purchases_id_sequence",
                "purchases_discounts_id_sequence",
                "purchases_products_id_sequence",
                "user_types_id_sequence",
                "users_id_sequence",
                "users_locations_id_sequence",
                "users_user_types_id_sequence",
        }

        for _, sequence := range sequences {
                _, err := database.Exec("ALTER SEQUENCE " + sequence + " RESTART WITH 1")

                if err != nil {
                        panic(err)
                }
        }
}

func Seed() {
        categories := insert_Categories()
        locations := insert_Locations()
        users := insert_Users()
        insert_Products(users, locations, categories)
        // insertProductsCategories()
        // insertProductsLocations()
}
