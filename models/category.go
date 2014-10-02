package models

import (
        "fmt"
        . "github.com/sescobb27/ciudad-gourmet/db"
)

type CategoryService interface {
        CreateCategory() (bool, error)
        GetCategories() []*Category
}

type Category struct {
        Id          int8
        Name        string
        Description string
}

func (c Category) CreateCategory() (bool, error) {
        db, err := StablishConnection()
        if err != nil {
                return false, err
        }
        defer db.Close()

        query := `INSERT INTO categories(name, description)
      VALUES ($1, $2)`

        _, err = db.Exec(query,
                c.Name,
                c.Description)

        if err != nil {
                return false, err
        }
        return true, nil
}

func (c Category) GetCategories() []*Category {
        db, err := StablishConnection()
        if err != nil {
                panic(err)
        }
        defer db.Close()

        query := `SELECT name FROM categories`
        categories_rows, err := db.Query(query)
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

func (c *Category) MarshalJSON() ([]byte, error) {
        str := fmt.Sprintf(`{"name": "%s", "description": "%s"}`, c.Name, c.Description)
        return []byte(str), nil
}

// ============ MOCKS and STUBS ============
type CategoryMock struct{}

func (l CategoryMock) CreateCategory() (bool, error) {
        return true, nil
}

func (l CategoryMock) GetCategories() []*Category {
        mock_category := &Category{
                Id:          1,
                Name:        "Category Mock",
                Description: "This is a Category Mock",
        }
        return []*Category{mock_category}
}
