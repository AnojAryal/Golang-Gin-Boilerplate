package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/anojaryal/JWT-Authentication/initializers"
	"github.com/anojaryal/JWT-Authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {

	fmt.Println("In Middleware")

	//get the cookie of req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode/validate it 
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{

		//check the expression
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0{
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to req
		c.Set("user", user)

		//continue
		c.Next()
	}else{
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}