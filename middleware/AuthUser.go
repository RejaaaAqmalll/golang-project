package middleware

import (
	"fmt"
	"net/http"
	"os"
	"set-up-Golang/config"
	"set-up-Golang/helper"
	"set-up-Golang/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Autharization")

	if err != nil {
		helper.Unauthorized(c, http.StatusUnauthorized, err)
		return
	}

	// Decode token yang diperoleh
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			helper.ExpiredToken(c, http.StatusUnauthorized, err)
			return
		}

		// temukan user yang login
		var user model.User

		id := claims["id_user"].(float64)
		config.DB.Find(&user, id)

		if user.Id == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"msg":    "User can't found",
				"reason": "User can't found",
				"data":   nil,
			})
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		panic(err)
	}
}
