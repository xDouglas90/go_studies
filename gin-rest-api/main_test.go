package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xdouglas90/go_studies/gin-rest-api/controllers"
	"github.com/xdouglas90/go_studies/gin-rest-api/database"
)

func RoutesTestSetup() *gin.Engine {
	r := gin.Default()
	return r
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestVerifySCGetStudents(t *testing.T) {
	database.Connect()
	r := RoutesTestSetup()
	r.GET("/students", controllers.GetStudents)
	w := performRequest(r, "GET", "/students")
	assert.Equal(t, http.StatusOK, w.Code, "OK response is expected")

	resMock := `[{"ID":1,"CreatedAt":"2022-11-07T17:44:00.439386-03:00","UpdatedAt":"2022-11-07T17:44:00.439386-03:00","DeletedAt":null,"name":"Júlia Maria","cpf":"01958459003","rg":"209897574"}]`
	resBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, resMock, string(resBody), "Response body mismatch")
}
