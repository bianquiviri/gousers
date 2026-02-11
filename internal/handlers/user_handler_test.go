package handlers

import (
	"bytes"
	"encoding/json"
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

	// Add a test user
	database.DB.Create(&models.User{Name: "Test User", Email: "test@example.com"})

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []models.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "Test User", users[0].Name)
}

func TestCreateUser(t *testing.T) {
	SetupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/users", CreateUser)

	userPayload := models.User{
		Name:  "New User",
		Email: "new@example.com",
	}
	jsonPayload, _ := json.Marshal(userPayload)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdUser models.User
	json.Unmarshal(w.Body.Bytes(), &createdUser)
	assert.Equal(t, userPayload.Name, createdUser.Name)
	assert.Equal(t, userPayload.Email, createdUser.Email)
}

func TestGetUser(t *testing.T) {
	SetupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users/:id", GetUser)

	// Add a test user
	user := models.User{Name: "Find Me", Email: "findme@example.com"}
	database.DB.Create(&user)

	// Test Success
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var foundUser models.User
	json.Unmarshal(w.Body.Bytes(), &foundUser)
	assert.Equal(t, "Find Me", foundUser.Name)

	// Test Not Found
	req404, _ := http.NewRequest("GET", "/users/999", nil)
	w404 := httptest.NewRecorder()
	r.ServeHTTP(w404, req404)
	assert.Equal(t, http.StatusNotFound, w404.Code)
}
