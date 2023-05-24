package controller

import (
	"fmt"
	"lify_backend/models"
	"lify_backend/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// created (register)
func (s *Server) CreatedUserController(c *gin.Context) {
	// _, cancel := context.WithCancel(context.Background())
	// defer cancel()

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
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, userCreated.Id))
	response.JSON(c, http.StatusCreated, "Succes", userCreated)

}

// get all mapping
func (s *Server) GetAllUserMapping(c *gin.Context) {
	// _, cancel := context.WithCancel(context.Background())
	// defer cancel()

	page := c.Query("page")
	offset := c.Query("offset")

	user := models.User{}
	getAllUser, count, _ := user.GetAllUser(s.DB, page, offset)

	getAllUserMapping := user.GetAllUserMap(s.DB, *getAllUser)

	response.GetJsonResponse(
		c, count, http.StatusOK, "Succes", getAllUserMapping,
	)

}

// // get all
// func (s *Server) GetAllUser(c *gin.Context) {
// 	page := c.Query("page")
// 	offset := c.Query("offset")

// 	user := models.User{}
// 	getAllUser, count, _ := user.GetAllUser(s.DB, page, offset)

// 	response.GetJsonResponse(
// 		c, count, http.StatusOK, "Succes", getAllUser,
// 	)

// }
