package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/xdouglas90/go_studies/gin-rest-api/controllers"
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
	r := RoutesTestSetup()
	r.GET("/students", controllers.GetStudents)
	w := performRequest(r, "GET", "/students")

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
