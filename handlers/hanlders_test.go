package handlers

import (
    "encoding/json"
    "github.com/julienschmidt/httprouter"
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
    ps := httprouter.Params{}

    Locations_Handler(recorder, req, ps)
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
    ps := httprouter.Params{}

    Categories_Handler(recorder, req, ps)
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
        strings.NewReader("username=test&email=test27@eafit.edu.co&lastname=Escobar&name=Simon&password=12345"),
    )
    assert.NoError(t, err)
    ps := httprouter.Params{}
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    SignUp_Handler(recorder, req, ps)
    assert.Equal(t, 200, recorder.Code, recorder.Body.String())
    assert.True(t, models.UserExist("test", "test27@eafit.edu.co"))
}

func TestUserSignIn(t *testing.T) {
    t.Parallel()
    recorder := httptest.NewRecorder()
    req, err := http.NewRequest(
        "POST",
        "/signin",
        strings.NewReader("username=sescob&password=qwerty"),
    )
    assert.NoError(t, err)
    ps := httprouter.Params{}
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    SignIn_Handler(recorder, req, ps)
    assert.Equal(t, 200, recorder.Code)
}
