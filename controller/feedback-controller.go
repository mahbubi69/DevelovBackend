package controller

import (
	"develov_be/models"
	"develov_be/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// created
func (s *Server) CreatedFeedBackController(c *gin.Context) {

	feedBack := models.FeedBack{}

	if err := c.BindJSON(&feedBack); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	mentorCreated, err := feedBack.CreatedFeedBack(s.DB)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, mentorCreated.IdMentor))
	response.JSON(c, http.StatusCreated, "Succes")

}
