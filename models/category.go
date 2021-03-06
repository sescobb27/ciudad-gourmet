package models

import (
    sql "github.com/sescobb27/ciudad-gourmet/db"
)

type Category struct {
    Id          int8   `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

func (c *Category) Create() (bool, error) {
    query := `INSERT INTO categories(name, description)
      VALUES ($1, $2) RETURNING id`
    var category_id int8
    err := sql.DB.QueryRow(query, c.Name, c.Description).Scan(&category_id)

    if err != nil {
        return false, err
    }
    c.Id = category_id

    return true, nil
}

func GetCategories() ([]*Category, error) {
    query := `SELECT id, name FROM categories`
    categories_rows, err := sql.DB.Query(query)
    categories := []*Category{}
    if err != nil {
        return categories, err
    }
    defer categories_rows.Close()

    if categories_rows == nil {
        return categories, nil
    }

    for categories_rows.Next() {
        category := Category{}
        err = categories_rows.Scan(
            &category.Id,
            &category.Name,
        )
        categories = append(categories, &category)
    }

    return categories, nil
}
