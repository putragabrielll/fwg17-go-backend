package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// var identityKey = "id"

// 2) PAYLOAD(membuat JWT) memperoleh data dari AUTHENTICATOR
func payload(data interface{}) jwt.MapClaims { // PayloadFunc
	user := data.(*services.PersonNet)
	return jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
	}
}

// 3) IdentityHandler membaca dari PAYLOAD
func identity(c *gin.Context) interface{} { // IdentityHandler
	claims := jwt.ExtractClaims(c)
	return &services.PersonNet{
		Id:   int(claims["id"].(float64)),
		Role: claims["role"].(string),
	}
}

// 1) AUTHENTICATOR mengembalikan data dari "models.FindUsersByEmail"
func authenticator(c *gin.Context) (interface{}, error) { // Authenticator
	loginauth := services.RLUsers{}
	err := c.ShouldBind(&loginauth)
	if err != nil {
		return nil, errors.New("Format Email not Support!")
	}

	finduser, err := models.FindUsersByEmail(loginauth.Email)
	if err != nil {
		return nil, errors.New("Email not register!")
	}

	checkpas, _ := argonize.DecodeHashStr(finduser.Password)
	paswdcheck := []byte(loginauth.Password)
	if !checkpas.IsValidPassword(paswdcheck) {
		return nil, errors.New("Password wrong!")
	
	}
	return &services.PersonNet{
		Id:   finduser.Id,
		Role: finduser.Role,
	}, nil
}

// 4) Authorizator membaca dari IdentityHandler
func authorizator(data interface{}, c *gin.Context) bool { // Authorizator
	userRoleCheck := data.(*services.PersonNet)
	if strings.HasPrefix(c.Request.URL.Path, "/admin") {
		if userRoleCheck.Role == "admin" {
			return true
		}
	} else if strings.HasPrefix(c.Request.URL.Path, "/customer") {
		if userRoleCheck.Role == "customer" {
			return true
		}
	}
	return false // Untuk role
}

func unauth(c *gin.Context, code int, message string) { // Unauthorized
	c.JSON(http.StatusUnauthorized, &services.ResponseBack{
		Success: false,
		Message: message,
	})
}

func loginresp(c *gin.Context, code int, token string, time time.Time) { // LoginResponse
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Login success!",
		Results: services.TokenTemp{
			Token: token,
		},
	})
}





// flow dari JWT dan verifikasi.
// AUTHENTICATOR => PAYLOADFUNC => UNAUTHORIZED => IDENTITYHANDLER => AUTHORIZATOR => LOGINRESPONSE



func Auth() (*jwt.GinJWTMiddleware, error) {
	godotenv.Load()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "go-backend",
		Key:             []byte(os.Getenv("APP_SECRET")),
		IdentityKey:     "id",
		PayloadFunc:     payload,
		IdentityHandler: identity,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauth,
		LoginResponse:   loginresp,
	})
	if err != nil {
		return nil, err
	}
	return authMiddleware, err
}
