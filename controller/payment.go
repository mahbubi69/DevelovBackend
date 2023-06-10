package controller

import (
	"develov_be/helper"
	"develov_be/models"
	"develov_be/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create
func (s *Server) CreatedPaymentController(c *gin.Context) {
	payment := models.Payment{}

	noRek := c.Request.PostFormValue("noRek")
	fmt.Println("value  : ", noRek)

	fileImage, err := helper.UploadImage(c)
	fmt.Println(fileImage)

	if fileImage == "" || err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	payment.NoRek = noRek
	payment.Struk = fileImage

	createdpPayment, err := payment.CreatedPayment(s.DB)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}
	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, createdpPayment.Id))
	response.JSON(c, http.StatusCreated, "Succes")
}

// read
func (s *Server) GetAllPaymentController(c *gin.Context) {
	payment := models.Payment{}

	getAllPayment, count, _ := payment.GetAllPayment(s.DB)

	response.GetJsonResponse(
		c, count, http.StatusOK, "Succes", getAllPayment)
}
