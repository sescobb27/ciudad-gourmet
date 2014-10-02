package models

import (
        "github.com/stretchr/testify/assert"
        "testing"
)

func TestGetLocations(t *testing.T) {
        locationService := LocationMock{}
        locations := locationService.GetLocations()
        assert.NotEmpty(t, locations)
}
