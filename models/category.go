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
      VALUES ($1, $2)`

    _, err := sql.DB.Exec(
        query,
        c.Name,
        c.Description,
    )

    if err != nil {
        return false, err
    }
    return true, nil
}

func GetCategories() []*Category {
    query := `SELECT name FROM categories`
    categories_rows, err := sql.DB.Query(query)
    if err != nil {
        panic(err)
    }

    if categories_rows == nil {
        panic(categories_rows)
    }

    categories := []*Category{}
    for categories_rows.Next() {
        category := Category{}
        err = categories_rows.Scan(&category.Name)
        categories = append(categories, &category)
    }

    return categories
}
