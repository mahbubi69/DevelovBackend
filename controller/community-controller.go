package controller

import (
	"develov_be/helper"
	"develov_be/models"
	"develov_be/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// created
func (s *Server) CreatedCommunityController(c *gin.Context) {
	community := models.Community{}

	fileImage, _ := helper.UploadImage(c)
	fmt.Println(fileImage)

	community.Title = c.PostForm("title")
	community.Deskripsi = c.PostForm("deskripsi")
	community.Image = fileImage

	createdCommunity, err := community.CreatedCommunity(s.DB)

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, createdCommunity.Id))
	response.JSON(c, http.StatusCreated, "Succes")
}

// get all
func (s *Server) GetAllCommunityController(c *gin.Context) {
	page := c.Query("page")
	offset := c.Query("offset")

	community := models.Community{}
	getAllCommunity, count, _ := community.GetACommunity(s.DB, page, offset)

	response.GetJsonResponse(
		c, count, http.StatusOK, "Succes", getAllCommunity,
	)
}

// searc community by title
func (s *Server) SearchCommunityByTitleController(c *gin.Context) {
	title := c.Query("title")

	community := models.Community{}
	searchCommunity, count, _ := community.SearchCommunityByNama(s.DB, title)
	if count == 0 {
		response.ErrorResponseAuth(c, http.StatusOK, "maaf yang anda cari tidak ada")
		return
	}

	response.GetJsonResponse(
		c, count, http.StatusOK, "Succes", searchCommunity,
	)
}
