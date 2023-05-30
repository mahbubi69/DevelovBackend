package controller

import (
	"develov_be/models"
	"develov_be/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// comment community
func (s *Server) CreatedCommentCommunityController(c *gin.Context) {
	comment := models.Comment{}

	if err := c.BindJSON(&comment); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	addComment, err := comment.CreatedComment(s.DB)

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, addComment.Id))
	response.JSON(c, http.StatusCreated, "Succes")
}
