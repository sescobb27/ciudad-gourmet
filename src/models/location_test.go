package models

import (
        "testing"
)

func TestGetLocations(t *testing.T) {
        locations := GetLocations()
        if len(locations) == 0 {
                t.Fatal("Locations should not be empty")
        }
}
