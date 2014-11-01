package models

import (
    "errors"
    sql "github.com/sescobb27/ciudad-gourmet/db"
    "github.com/sescobb27/ciudad-gourmet/helpers"
    "log"
    "time"
)

type Product struct {
    Id          int64       `json:"id"`
    CreatedAt   time.Time   `json:"createdat,omitempty"`
    Image       string      `json:"image"`
    Description string      `json:"description"`
    Name        string      `json:"name"`
    Price       float64     `json:"price"`
    Rate        float32     `json:"rate"`
    Errors      []string    `json:"errors"`
    Chef        *User       `json:"chef"`
    Categories  []*Category `json:"categories"`
    Discounts   []*Discount `json:"discounts,omitempty"`
}

func (p *Product) Create() {
    query := `INSERT INTO products(
            created_at, description, image, name, price, rate, chef_id)
            VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
    query2 := `INSERT INTO products_categories(category_id, product_id)
            VALUES ($1, $2)`
    var tx_err, err error
    var product_id int64
    tx, tx_err := sql.DB.Begin()

    if tx_err != nil {
        err = tx.Rollback()
        if err != nil {
            log.Fatal(err)
        }
        log.Fatal(tx_err)
        panic(err)
    }

    tx_err = tx.QueryRow(
        query,
        p.CreatedAt,
        p.Description,
        p.Image,
        p.Name,
        p.Price,
        p.Rate,
        p.Chef.Id,
    ).Scan(&product_id)

    if tx_err != nil {
        err = tx.Rollback()
        if err != nil {
            log.Fatal(err)
        }
        log.Fatal(tx_err)
        panic(err)
    }

    for _, category := range p.Categories {
        _, tx_err = tx.Exec(query2, category.Id, product_id)
        if tx_err != nil {
            err = tx.Rollback()
            if err != nil {
                log.Fatal(err)
            }
            log.Fatal(tx_err)
            panic(err)
        }
    }

    tx_err = tx.Commit()
    if tx_err != nil {
        err = tx.Rollback()
        if err != nil {
            log.Fatal(err)
        }
        log.Fatal(tx_err)
        panic(err)
    }
}

//  ======
func FindProductsByName(name string) ([]*Product, error) {
    if len(name) == 0 || !helpers.ProductNameValidator(name) {
        return nil, errors.New("Nombre del Producto Invalido")
    }

    query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
            FROM products AS p
            WHERE LOWER(p.name) LIKE '%' || $1 || '%' ORDER BY p.rate DESC`

    product_rows, err := sql.DB.Query(query, name)

    if err != nil {
        return nil, err
    }

    if product_rows == nil {
        return nil, errors.New("No Products Named " + name)
    }

    defer product_rows.Close()

    products := []*Product{}
    for product_rows.Next() {
        product := new(Product)
        err = product_rows.Scan(
            &product.Id,
            &product.Name,
            &product.Description,
            &product.Price,
            &product.Image,
            &product.Rate,
        )
        if err != nil {
            panic(err)
        }
        products = append(products, product)
    }
    if err = product_rows.Err(); err != nil {

    }
    return products, nil
}

//  ======
func FindProductsByCategory(category string) ([]*Product, error) {
    query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
          FROM products AS p
              INNER JOIN products_categories as pc ON ( p.id = pc.product_id )
                  INNER JOIN categories AS c ON ( pc.category_id = c.id )
          WHERE LOWER(c.name) = LOWER('$1') ORDER BY p.rate DESC`

    product_rows, err := sql.DB.Query(query, category)

    if err != nil {
        return nil, err
    }

    products := []*Product{}
    if product_rows == nil {
        return products, errors.New("No Products For Category: " + category)
    }

    defer product_rows.Close()

    for product_rows.Next() {
        product := new(Product)
        err = product_rows.Scan(
            &product.Id,
            &product.Name,
            &product.Description,
            &product.Price,
            &product.Image,
            &product.Rate,
        )
        if err != nil {
            return products, err
        }
        products = append(products, product)
    }

    if err = product_rows.Err(); err != nil {
        return products, err
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
    query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
      FROM products as p
      INNER JOIN users as u on ( p.chef_id = u.id )
      WHERE LOWER(u.username) = LOWER($1)`

    product_rows, err := sql.DB.Query(query, username)

    if err != nil {
        return nil, err
    }

    products := []*Product{}
    if product_rows == nil {
        return products, errors.New("No Products For User: " + username)
    }

    defer product_rows.Close()

    for product_rows.Next() {
        product := new(Product)
        err = product_rows.Scan(
            &product.Id,
            &product.Name,
            &product.Description,
            &product.Price,
            &product.Image,
            &product.Rate,
        )
        if err != nil {
            return products, err
        }
        products = append(products, product)
    }

    if err = product_rows.Err(); err != nil {
        return products, err
    }
    return products, nil
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
