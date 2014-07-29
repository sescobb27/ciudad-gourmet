package models

import (
        "db"
        "log"
        "time"
)

type Product struct {
        Id                       int64
        CreatedAt                time.Time
        Image, Description, Name string
        Price                    float64
        Rate                     float32
        Errors                   []string
        Chef                     *User
        Categories               []*Category
        Discounts                []*Discount
}

func (p *Product) Create() {
        db, err := db.StablishConnection()
        if err != nil {
                log.Fatal(err)
                panic(err)
        }
        defer db.Close()

        query := `INSERT INTO products(
            id, created_at, description, image, name, price, rate, chef_id)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

        _, err = db.Exec(query,
                p.Id,
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
func (p *Product) FindByName() {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //     FROM products AS p
        //     WHERE LOWER(p.name) LIKE %$1% ORDER BY p.rate DESC`
}

//  ======
func (p *Product) FindByCategory() {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM products AS p
        //       INNER JOIN products_categories as pc ON ( p.id = pc.product_id )
        //           INNER JOIN categories AS c ON ( pc.category_id = c.id )
        //   WHERE LOWER(c.name) = LOWER('$1') ORDER BY p.rate DESC`

}

//  ====== TODO =======
func (p *Product) FindByLocation() {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM Product p
        //   INNER JOIN locations AS loc ON (p.) INNER JOIN pl.location l
        //   WHERE LOWER(l.name) = LOWER(:location) ORDER BY p.rate DESC`

}

//  ======
func (p *Product) FindByUserName() {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM products as p
        //   INNER JOIN users as u on ( p.chef_id = u.id )
        //   WHERE LOWER(u.username) = LOWER($1)`

}

//  ======
func (p *Product) FindByPrice() {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM products as p
        //   WHERE p.price >= $1 AND p.price <= $2
        //   ORDER BY p.price ASC`

}

//  ======
func (p *Product) FindById() {
        // query := `SELECT p.id, p.name, p.description, p.price, p.image, p.rate
        //   FROM products as p
        //   WHERE p.id = $1`

}
