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
func (s *Server) CreatedMentorController(c *gin.Context) {
	mentor := models.Mentor{}

	fileImage, err := helper.UploadImage(c)
	fmt.Println(fileImage)

	if fileImage == "" || err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	specialistInt, _ := strconv.Atoi(c.PostForm("specialist"))

	mentor.Nama = c.PostForm("nama")
	mentor.Profile = fileImage
	mentor.Specialist = specialistInt
	mentor.Portofolio = c.PostForm("portofolio")
	mentor.Salary = c.PostForm("salary")
	mentor.NamaRekening = c.PostForm("namaRekening")
	mentor.NomorRekening = c.PostForm("nomorRekening")

	mentorCreated, err := mentor.CreatedMentor(s.DB)

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, mentorCreated.Id))
	response.JSON(c, http.StatusCreated, "Succes")

}

// get all mentor
func (s *Server) GetAllMentorController(c *gin.Context) {
	page := c.Query("page")
	offset := c.Query("offset")

	mentor := models.Mentor{}
	getAllMentor, count, _ := mentor.GetAllMentor(s.DB, page, offset)

	response.GetJsonResponse(
		c, count, http.StatusOK, "Succes", getAllMentor)
}

// update
func (s *Server) UpdateMentorController(c *gin.Context) {
	id := c.Param("id")

	idMentor, _ := strconv.ParseInt(id, 10, 64)

	mentor := models.Mentor{}

	if err := c.BindJSON(&mentor); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	updateMentor, err := mentor.UpdateMentor(s.DB, uint32(idMentor))

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}
	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, updateMentor.Id))
	response.JSON(c, http.StatusCreated, "Succes Update")
}

// delete
func (s *Server) DeleteMentorController(c *gin.Context) {
	id := c.Param("id")

	idMentor, _ := strconv.ParseInt(id, 10, 64)

	mentor := models.Mentor{}

	deleteMentor, err := mentor.DeleteMentor(s.DB, uint32(idMentor))

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}
	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, deleteMentor.Id))
	response.JSON(c, http.StatusCreated, "Succes Delete")
}
