package middleware

import (
	"develov_be/auth"
	"develov_be/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetMiddlewareAuthencation(nxt gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			resp := response.ServerResponse{
				Status:  http.StatusUnauthorized,
				Message: "token tidak ada",
			}
			c.JSON(http.StatusUnauthorized, resp)
			return
		}
		nxt(c)
	}
}
