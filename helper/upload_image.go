package helper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ImageName() string {
	b := make([]byte, 12) //equals 24 characters
	rand.Read(b)
	s := hex.EncodeToString(b)

	return s
}

func UploadImage(c *gin.Context) (string, error) {
	c.Request.ParseMultipartForm(10 << 20)

	uploadedFile, handler, err := c.Request.FormFile("images")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return "", err
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd() // "/_(path)"
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	//name image random
	imageFile := ImageName()
	fileName := fmt.Sprintf("%s%s", imageFile, filepath.Ext(handler.Filename))

	fileLocation := filepath.Join(dir, "assets/images/", fileName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	filePath := `/assets/images/` + fileName
	fmt.Println(filePath)

	return filePath, nil
}
