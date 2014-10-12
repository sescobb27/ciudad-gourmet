package handlers

import (
        "encoding/json"
        "github.com/sescobb27/ciudad-gourmet/models"
        "io/ioutil"
        "strings"
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

        locations := []*models.Location{}
        body, err := ioutil.ReadAll(recorder.Body)
        assert.NoError(t, err)

        err = json.Unmarshal(body, &locations)
        assert.NoError(t, err)
        assert.NotEmpty(t, locations)

}

func TestGetCategories(t *testing.T) {
        t.Parallel()

        recorder := httptest.NewRecorder()
        req, err := http.NewRequest("GET", "/categories", nil)

        Categories_Handler(recorder, req)
        assert.NoError(t, err)
        assert.Equal(t, 200, recorder.Code)
        assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

        categories := []*models.Category{}
        body, err := ioutil.ReadAll(recorder.Body)
        assert.NoError(t, err)

        err = json.Unmarshal(body, &categories)
        assert.NoError(t, err)
        assert.NotEmpty(t, categories)
}

func TestUserSignUp(t *testing.T) {
        t.Parallel()
        recorder := httptest.NewRecorder()
        req, err := http.NewRequest(
                "POST",
                "/signup",
                strings.NewReader("username=sescob27&email=sescob27@eafit.edu.co&lastname=Escobar&name=Simon&password=12345"),
        )
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
        assert.NoError(t, err)
        SignUp_Handler(recorder, req)
        assert.Equal(t, 200, recorder.Code, recorder.Body.String())
        username := "sescob27"
        user, err := models.FindUserByUsername(&username)
        assert.NoError(t, err)
        assert.Equal(t, "sescob27", user.Username)
        assert.Equal(t, "sescob27@eafit.edu.co", user.Email)
}
