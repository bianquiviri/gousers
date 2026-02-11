package handlers

import (
	"gousers/internal/database"
	"gousers/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	database.DB = db
}

func TestGetUsers(t *testing.T) {
	SetupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users", GetUsers)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
