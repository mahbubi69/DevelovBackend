package controller

import (
	"develov_be/models"
	"develov_be/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// created
func (s *Server) CreatedTicketController(c *gin.Context) {
	tiket := models.Ticket{}

	tiket.CodeTiket = "8565484I87Op"

	createdTools, err := tiket.CreatedTicket(s.DB)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, createdTools.Id))
	response.JSON(c, http.StatusCreated, "Succes")

}

// udpate
func (s *Server) UpdateTictketController(c *gin.Context) {
	id := c.Param("id")

	idTicket, _ := strconv.ParseInt(id, 10, 64)

	ticket := models.Ticket{}

	if err := c.BindJSON(&ticket); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	updateMentor, err := ticket.UpdateTicket(s.DB, uint32(idTicket))

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}
	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, updateMentor.Id))
	response.JSON(c, http.StatusCreated, "Succes Update")
}
