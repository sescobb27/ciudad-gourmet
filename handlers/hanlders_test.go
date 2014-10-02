package handlers

import (
        "encoding/json"
        "github.com/sescobb27/ciudad-gourmet/models"
        "io/ioutil"
        // . "github.com/smartystreets/goconvey/convey"
        "github.com/stretchr/testify/assert"
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestGetLocations(t *testing.T) {
        t.Parallel()
        locations := []*models.Location{}
        server := httptest.NewServer(Locations_Handler(locations))
        res, err := http.Get(server.URL)

        assert.NoError(t, err)
        assert.Equal(t, 200, res.StatusCode)
        assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
        // assert.NotEmpty(t, locations)
        body, err := ioutil.ReadAll(res.Body)
        assert.NoError(t, err)
        res_locations := []*models.Location{}
        err = json.Unmarshal(body, &res_locations)
        assert.NoError(t, err)
        assert.NotEmpty(t, res_locations)
}

func TestGetCategories(t *testing.T) {
        t.Parallel()
        categories := []*models.Category{}
        server := httptest.NewServer(Categories_Handler(categories))
        res, err := http.Get(server.URL)
        assert.NoError(t, err)
        assert.Equal(t, 200, res.StatusCode)
        assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
        res_categories := []*models.Category{}
        body, err := ioutil.ReadAll(res.Body)
        assert.NoError(t, err)
        err = json.Unmarshal(body, &res_categories)
        assert.NoError(t, err)
        assert.NotEmpty(t, res_categories)
}
