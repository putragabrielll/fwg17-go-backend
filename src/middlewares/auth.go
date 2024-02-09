package middlewares

import (
	"net/http"
	"os"
	"time"
	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// var identityKey = "id"

func payload(data interface{}) jwt.MapClaims { // PayloadFunc
	user := data.(*services.PersonNet)
	return jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
	}
}

func identity(c *gin.Context) interface{} { // IdentityHandler
	claims := jwt.ExtractClaims(c)
	return &services.PersonNet{
		Id:   int(claims["id"].(float64)),
		Role: claims["role"].(string),
	}
}

func authenticator(c *gin.Context) (interface{}, error) { // Authenticator
	loginauth := services.RLUsers{}
	err := c.ShouldBind(&loginauth)
	if err != nil {
		return nil, err
	}

	finduser, err := models.FindUsersByEmail(loginauth.Email)
	if err != nil {
		return nil, err
	}

	checkpas, _ := argonize.DecodeHashStr(finduser.Password)
	paswdcheck := []byte(loginauth.Password)
	if !checkpas.IsValidPassword(paswdcheck) {
		return nil, err
	}
	return &services.PersonNet{
		Id: finduser.Id,
		Role: finduser.Role,
	}, nil
}

func authorizator(data interface{}, c *gin.Context) bool{ // Authorizator
	return true // Untuk role
}

func unauth(c *gin.Context, code int, message string) { // Unauthorized
	c.JSON(http.StatusUnauthorized, &services.ResponseBack{
		Success: false,
		Message: "Unauthorized!",
	})
}

func loginresp(c *gin.Context, code int, token string, time time.Time){ // LoginResponse
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Login success!",
		Results: struct{
			Token string `json:"token"`
		}{
			Token: token,
		},
	})
}








func Auth() (*jwt.GinJWTMiddleware, error){
	godotenv.Load()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           	"go-backend",
		Key:             	[]byte(os.Getenv("APP_SECRET")),
		IdentityKey:     	"id",
		PayloadFunc:     	payload,
		IdentityHandler: 	identity,
		Authenticator:   	authenticator,
		Authorizator: 		authorizator,
		Unauthorized:    	unauth,
		LoginResponse: 		loginresp,
	})
	if err != nil {
		return nil, err
	}
	return authMiddleware, err
}
