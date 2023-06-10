package controller

import (
	"develov_be/auth"
	"develov_be/helper"
	"develov_be/models"
	"develov_be/response"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// login
func (s *Server) LoginUserController(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	token, err := user.SignIn(s.DB, user.Id, user.Email, user.Password)

	if err != nil {
		response.ErrorSigInResponse(c, http.StatusBadRequest, token)
		return
	}

	fmt.Printf("password user : %s \n", user.Password)

	fmt.Printf("token user : %s \n", token)
	user.UpdateTokenUser(s.DB, user.Email, token)

	userDetail, err := user.GetUserByEmail(s.DB, user.Email)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	response.JSONLOGIN(c, http.StatusOK, "Berhasil Login", userDetail.Token)
}

// created (register)
func (s *Server) CreatedUserController(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	newUUID := models.NewUUID()
	user.Uuid = newUUID
	fmt.Printf("uuid User : %s", newUUID)

	hashPassword := models.HashPasswordToSha256(user.Password)
	user.Password = hashPassword

	userCreated, err := user.CreateUser(s.DB)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, userCreated.Id))
	response.JSON(c, http.StatusCreated, "Succes")

}

// get all mapping
func (s *Server) GetAllUserMappingController(c *gin.Context) {
	page := c.Query("page")
	offset := c.Query("offset")

	user := models.User{}
	getAllUser, count, _ := user.GetAllUser(s.DB, page, offset)

	getAllUserMapping := user.GetAllUserMap(s.DB, *getAllUser)

	response.GetJsonResponse(
		c, count, http.StatusOK, "Succes", getAllUserMapping,
	)
}

// get profile
func (s *Server) GetUserByTokenController(c *gin.Context) {
	token := auth.ExtractToken(c)
	fmt.Printf("token User : %s", token)

	user := models.User{}

	getUserByToken, err := user.GetUserByToken(s.DB, token)

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	mappingUser := models.ResponseUserMapping{
		Nama:     getUserByToken.Nama,
		Profile:  getUserByToken.Profile,
		UserName: getUserByToken.UserName,
		Email:    getUserByToken.Email,
		NoHp:     getUserByToken.NoHp,
		Role:     getUserByToken.Role,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil",
		"data":    mappingUser,
	})

}

func (s *Server) UpdateImageController(c *gin.Context) {
	user := models.User{}

	token := auth.ExtractToken(c)
	fmt.Printf("token User : %s \n", token)

	fileImage, err := helper.UploadImage(c)
	fmt.Println(fileImage)

	if fileImage == "" || err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	updateImage, err := user.UpdateImage(s.DB, token, fileImage)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, updateImage.Uuid))
	response.JSON(c, http.StatusCreated, "Update Succes")
}

// read image helper
func (s *Server) ReadImagesController(c *gin.Context) {
	fileName := c.Param("file")
	img, err := os.Open("assets/images/" + fileName)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	defer img.Close()
	c.Writer.Header().Set("Content-Type", "image/jpeg")
	io.Copy(c.Writer, img)
}

// delete
func (s *Server) DeleteUserController(c *gin.Context) {
	id := c.Param("id")
	idMentor, _ := strconv.ParseInt(id, 10, 64)

	user := models.User{}

	deleteMentor, err := user.DeleteUser(s.DB, uint32(idMentor))

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}
	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, deleteMentor.Id))
	response.JSON(c, http.StatusCreated, "Succes Delete")
}

// delete image
func (s *Server) DeleteImageUserController(c *gin.Context) {
	user := models.User{}

	token := auth.ExtractToken(c)

	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	deleteImage, _ := user.UpdateImage(s.DB, token, user.Profile)

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, deleteImage.Id))
	response.JSON(c, http.StatusCreated, "Succes Delete Image")

}
