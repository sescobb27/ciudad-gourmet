package handlers

import (
        // . "github.com/smartystreets/goconvey/convey"
        "github.com/stretchr/testify/assert"
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestGetLocations(t *testing.T) {
        t.Parallel()
        recorder := httptest.NewRecorder()
        req, err := http.NewRequest("GET", "/locations", nil)
        Locations_Handler(recorder, req)
        assert.NoError(t, err)
        assert.Equal(t, 200, recorder.Code)
        assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
}

func TestGetCategories(t *testing.T) {
        t.Parallel()
        recorder := httptest.NewRecorder()
        req, err := http.NewRequest("GET", "/categories", nil)
        Categories_Handler(recorder, req)
        assert.NoError(t, err)
        assert.Equal(t, 200, recorder.Code)
        assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
}
