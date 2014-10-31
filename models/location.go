package models

import (
    sql "github.com/sescobb27/ciudad-gourmet/db"
)

type Location struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func (l *Location) Create() (bool, error) {
    query := `INSERT INTO locations(name) VALUES ($1)`
    _, err := sql.DB.Exec(query, l.Name)

    if err != nil {
        return false, err
    }
    return true, nil
}

func GetLocations() []*Location {
    query := `SELECT name FROM locations`

    location_rows, err := sql.DB.Query(query)
    if err != nil {
        panic(err)
    }

    if location_rows == nil {
        panic(location_rows)
    }

    locations := []*Location{}
    for location_rows.Next() {
        location := Location{}
        err = location_rows.Scan(&location.Name)
        locations = append(locations, &location)
    }

    return locations
}
