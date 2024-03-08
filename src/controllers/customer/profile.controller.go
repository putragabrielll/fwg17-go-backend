package customerController

import (
	"fmt"
	"net/http"
	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/middlewares"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT PROFILE USERS
func ProfileUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))
	// fmt.Println(id)

	profileUsers, err := models.FindUsersId(id)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail users!",
		Results: profileUsers,
	})
}

// UPDATE PROFILE
func UpdateUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))
	// fmt.Println(id, "masuk update!")

	profileData := services.Person{}
	c.ShouldBind(&profileData) // untuk memasukkan data dari form ke struck Person{}
	
	if profileData.Password != "" {
		paswdhash := []byte(profileData.Password) // proses hashing password
		hasedPasswd, _ := argonize.Hash(paswdhash)
		profileData.Password = hasedPasswd.String()
	}

	// -------------
	file, _ := c.FormFile("picture")
	if file != nil {
		cekFile, err := middlewares.UploadFile(c, "picture", "profile") // fungsi upload file
		if err != nil {
			msg := fmt.Sprintf("%v", err)
			helpers.Utils(err, msg, c) // Error Handler
			return
		}
		profileData.Picture = cekFile
	}
	// -------------
	profileData.Id = id // mengarahkan isi dari profileData dengan value "id" di ambil dari id di atas.

	updatedUsers, err := models.UpdateUsers(profileData)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Users updated successfully!",
		Results: updatedUsers,
	})
}
