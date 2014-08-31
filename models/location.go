package models

import (
        "fmt"
        . "github.com/sescobb27/ciudad-gourmet/db"
)

type Location struct {
        Id      int8
        Name    string
}

func (l *Location) Create() {
        db, err := StablishConnection()
        if err != nil {
                panic(err)
        }
        defer db.Close()

        query := `INSERT INTO locations(name) VALUES ($1)`
        _, err = db.Exec(query, l.Name)

        if err != nil {
                panic(err)
        }
}

func GetLocations() []*Location {
        db, err := StablishConnection()
        if err != nil {
                panic(err)
        }
        defer db.Close()

        query := `SELECT name FROM locations`

        location_rows, err := db.Query(query)
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

func (l *Location) MarshalJSON() ([]byte, error) {
        str := fmt.Sprintf(`{"name": "%s"}`, l.Name)
        return []byte(str), nil
}
