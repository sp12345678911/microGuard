package test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRulChecker(t *testing.T) {
	r := gin.Default()
	r.GET("/rul-checker", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "RUL Checker is running"})
	})

	w := performRequest(r, "GET", "/rul-checker")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expected := `{"message":"RUL Checker is running"}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, w.Body.String())
	}
}
