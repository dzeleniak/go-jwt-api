package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dzeleniak/jwt-api/initializers"
	"github.com/dzeleniak/jwt-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {

	// Get the cookie off the req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		log.Println("Couldnt find Authorization cookie...")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/Validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Println("Authorization cookie is expired...")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Check the user with token sub
		var user models.User

		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			log.Println("Couldnt find requested user...")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()

	} else {
		log.Println("Token is invalid...")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
