package middlewares

import (
	"errors"
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context, dest string) (string, error) {
	file, _ := c.FormFile("picture") // "picture" => nama field / nama form
	extensionFile := file.Header["Content-Type"][0]
	fmt.Println(file.Header)

	ext := map[string]string{
		"image/png":  ".png",
		"image/jpeg": ".jpeg",
		"image/jpg":  ".jpg",
	}
	fileName := fmt.Sprintf("uploads/%v/%v%v", dest, uuid.NewString(), ext[extensionFile])

	// START VALIDATION FILE EXTENSION
	validTypeFile := []string{
		"image/png",
		"image/jpeg",
		"image/jpg",
	}
	if !(strings.Contains(strings.Join(validTypeFile, ","), extensionFile)) {
		return "", errors.New("Extension not support!")
	}
	if file.Size > (2 * 1024 * 1024) {
		return "", errors.New(("File to over size, Max file upload 2MB!"))
	}
	// END VALIDATION

	c.SaveUploadedFile(file, fileName)
	return fileName, c.Err()
}
