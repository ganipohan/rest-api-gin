package handlers

import (
	"gin-rest-api/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("faijsdhfAJSDFAjasdjKNKJJkjhkjfKAKSDJFkjka2342039uakdnflkan09809")

// Register mengelola pendaftaran pengguna
func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := models.CreateUser(user.Username, user.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// Login mengelola autentikasi pengguna dan menghasilkan token JWT
func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    dbUser, err := models.GetUserByUsername(user.Username)
    if err != nil || bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":    dbUser.ID,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })
    
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
