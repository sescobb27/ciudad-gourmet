package models

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGetLocations(t *testing.T) {
    locations, err := GetLocations()
    assert.NoError(t, err)
    assert.NotEmpty(t, locations)
}
