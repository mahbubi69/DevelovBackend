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
func (s *Server) CreatedScheduleController(c *gin.Context) {
	schedule := models.Schedule{}

	if err := c.BindJSON(&schedule); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	scheduleCreated, err := schedule.CreatedSchedule(s.DB)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, scheduleCreated.Id))
	response.JSON(c, http.StatusCreated, "Succes")
}

// get
func (s *Server) GetAllScheduleController(c *gin.Context) {
	schedule := models.Schedule{}

	getAllSchedule, count, _ := schedule.GetSchedule(s.DB)
	response.GetJsonResponse(
		c, count, http.StatusOK, "Succes", getAllSchedule)
}

// udpate
func (s *Server) UpdateScheduleController(c *gin.Context) {
	id := c.Param("id")
	idSchedule, _ := strconv.ParseInt(id, 10, 64)

	schedule := models.Schedule{}

	if err := c.BindJSON(&schedule); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	updateSchedule, err := schedule.UpdateSchedule(s.DB, uint32(idSchedule))

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, updateSchedule.Id))
	response.JSON(c, http.StatusCreated, "Succes")
}
