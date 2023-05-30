package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func TestServer_SendEmailOtpUserController(t *testing.T) {
	type fields struct {
		DB     *gorm.DB
		Router *gin.Engine
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				DB:     tt.fields.DB,
				Router: tt.fields.Router,
			}
			s.SendEmailOtpUserController(tt.args.c)
		})
	}
}
