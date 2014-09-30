package models

import (
        "errors"
        . "github.com/sescobb27/ciudad-gourmet/db"
        "github.com/sescobb27/ciudad-gourmet/helpers"
        "log"
        "time"
)

type Product struct {
        Id          int64       `json:"id"`
        CreatedAt   time.Time   `json:"createdat"`
        Image       string      `json:"image"`
        Description string      `json:"description"`
        Name        string      `json:"name"`
        Price       float64     `json:"price"`
        Rate        float32     `json:"rate"`
        Errors      []string    `json:"errors"`
        Chef        *User       `json:"chef"`
        Categories  []*Category `json:"categories"`
        Discounts   []*Discount `json:"discounts"`
}

func (p *Product) Create() {
        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `INSERT INTO products(
            created_at, description, image, name, price, rate, chef_id)
            VALUES ($1, $2, $3, $4, $5, $6, $7)`

        _, err = db.Exec(query,
                p.CreatedAt,
                p.Description,
                p.Image,
                p.Name,
                p.Price,
                p.Rate,
                p.Chef.Id)

        if err != nil {
                log.Fatal(err)
                panic(err)
        }
}

//  ======
func FindProductsByName(name string) ([]*Product, error) {
        if len(name) == 0 || !helpers.ProductNameValidator(name) {
                return nil, errors.New("Nombre del Producto Invalido")
        }

        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
            FROM products AS p
            WHERE LOWER(p.name) LIKE '%' || $1 || '%' ORDER BY p.rate DESC`

        product_rows, err := db.Query(query, name)

        if err != nil {
                return nil, err
        }

        if product_rows == nil {
                return nil, errors.New("No Products Named " + name)
        }

        products := []*Product{}
        for product_rows.Next() {
                product := new(Product)
                err = product_rows.Scan(&product.Id,
                        &product.Name,
                        &product.Description,
                        &product.Price,
                        &product.Image,
                        &product.Rate)
                if err != nil {
                        panic(err)
                }
                products = append(products, product)
        }
        return products, nil
}

//  ======
func FindProductsByCategory(category string) ([]*Product, error) {
        db, err := StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
          FROM products AS p
              INNER JOIN products_categories as pc ON ( p.id = pc.product_id )
                  INNER JOIN categories AS c ON ( pc.category_id = c.id )
          WHERE LOWER(c.name) = LOWER('$1') ORDER BY p.rate DESC`

        product_rows, err := db.Query(query, category)

        if err != nil {
                return nil, err
        }

        if product_rows == nil {
                return nil, errors.New("No Products For Category: " + category)
        }

        products := []*Product{}

        for product_rows.Next() {
                product := new(Product)
                err = product_rows.Scan(&product.Id,
                        &product.Name,
                        &product.Description,
                        &product.Price,
                        &product.Image,
                        &product.Rate)
                if err != nil {
                        panic(err)
                }
                products = append(products, product)
        }
        return products, nil
}

//  ====== TODO =======
func FindProductsByLocation(location string) ([]*Product, error) {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM Product p
        //   INNER JOIN locations AS loc ON (p.) INNER JOIN pl.location l
        //   WHERE LOWER(l.name) = LOWER(:location) ORDER BY p.rate DESC`
        return nil, nil
}

//  ======
func FindProductsByUserName(username string) ([]*Product, error) {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM products as p
        //   INNER JOIN users as u on ( p.chef_id = u.id )
        //   WHERE LOWER(u.username) = LOWER($1)`
        return nil, nil
}

//  ======
func FindProductsByPrice(price float64) ([]*Product, error) {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM products as p
        //   WHERE p.price >= $1 AND p.price <= $2
        //   ORDER BY p.price ASC`
        return nil, nil
}

//  ======
func FindProductsById(id int64) ([]*Product, error) {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM products as p
        //   WHERE p.id = $1`
        return nil, nil
}
