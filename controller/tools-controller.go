package controller

import (
	"develov_be/helper"
	"develov_be/models"
	"develov_be/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// created
func (s *Server) CreatedToolsController(c *gin.Context) {
	tools := models.Tools{}

	nama := c.Request.PostFormValue("nama")
	fmt.Println("value  : ", nama)

	fileImage, err := helper.UploadImage(c)
	fmt.Println(fileImage)

	if fileImage == "" || err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	tools.Nama = nama
	tools.Logo = fileImage

	createdTools, err := tools.CreatedTools(s.DB)

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, createdTools.Id))
	response.JSON(c, http.StatusCreated, "Succes")

}

// read
func (s *Server) ReadToolsController(c *gin.Context) {
	tools := models.Tools{}
	getAllTools, count, _ := tools.GetTools(s.DB)

	response.GetJsonResponse(
		c, count, http.StatusOK, "Succes", getAllTools,
	)
}

// update
func (s *Server) UpdateToolsController(c *gin.Context) {
	tools := models.Tools{}

	fileImage, _ := helper.UploadImage(c)
	fmt.Println(fileImage)

	nama := c.Request.PostFormValue("nama")
	fmt.Println("value  : ", nama)

	tools.Nama = nama
	tools.Logo = fileImage

	// token := auth.ExtractToken(c)

	updateTools, err := tools.UpdateTools(s.DB, 2)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, updateTools.Id))
	response.JSON(c, http.StatusCreated, "Succes")
}

// mentor tools
func (s *Server) CeatedMentorToolsController(c *gin.Context) {
	var tools []models.MentorTools

	paramIdMentor := c.Param("idMentor")
	idMentor, _ := strconv.Atoi(paramIdMentor)

	if err := c.BindJSON(&tools); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	for _, tool := range tools {
		tool.IdMentor = uint32(idMentor)
		_, err := tool.CreatedMentorTools(s.DB)
		if err != nil {
			response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		}

	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, tools))
	response.JSON(c, http.StatusCreated, "Succes")

}
