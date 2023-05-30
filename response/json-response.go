package response

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}
type JsonGetResponse struct {
	Count   uint64      `json:"count"`
	Status  uint16      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ServerResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type JsonLoginResponse struct {
	Status  uint16      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GenericJsonResponse(c *gin.Context, statusCode int, message string) {
	res := ServerResponse{
		Status:  statusCode,
		Message: message,
	}
	c.JSON(statusCode, res)
}

func GetJsonResponse(c *gin.Context, count uint64, statusCode uint16, message string, data interface{}) {
	res := JsonGetResponse{
		Count:   count,
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	c.JSON(int(statusCode), res)
}

func JSONLOGIN(c *gin.Context, statusCode int, message string, token string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"token":   token,
	})

}

func JSON(c *gin.Context, statusCode uint16, message string) {
	res := JsonResponse{
		Status:  statusCode,
		Message: message,
	}
	c.JSON(int(statusCode), res)
}

func ErrorResponseAuth(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": err,
	})
}

func ErrorSigInResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
	})
}
