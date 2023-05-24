package controller

import (
	"fmt"
	"lify_backend/models"
	"lify_backend/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) LoginUser(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	token, err := user.SignIn(s.DB, user.Id, user.Email, user.Password)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
	}

	fmt.Printf("password user : %s \n", user.Password)

	fmt.Printf("token user : %s \n", token)
	user.UpdateTokenUser(s.DB, user.Email, token)

	userDetail, err := user.GetUserByEmail(s.DB, user.Email)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
	}

	response.JSONLOGIN(c, http.StatusOK, "Berhasil Login", userDetail.Token)
}
