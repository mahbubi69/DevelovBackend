package controller

import (
	"develov_be/models"
	"develov_be/response"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
)

// chek email &n auto update otp
func (s *Server) ChekEmailUser(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	chekEmailuser, err := user.CekEmail(s.DB, user.UserName)
	fmt.Println(chekEmailuser)

	errEmail := errors.New("maaf Email anda tidak terdaftar")
	if err != nil {
		response.ErrorResponse(c, http.StatusUnauthorized, errEmail)
		return
	}

}

func (s *Server) SendEmailOtpUserController(c *gin.Context) {
	secretLength := 16
	otpAutoGenerate := gotp.RandomSecret(secretLength)
	newOtp := gotp.NewDefaultTOTP(otpAutoGenerate).Now()

	toEmailUserSendEmail := []string{"sari.142002@gmail.com"}
	subject := "PT.DEVELOV"
	otpMessageToEmail := "Hii\n OTP Anda Adalah: " + newOtp
	senderName := "PT.DEVELOV"

	s.SendOtpToEmail(toEmailUserSendEmail, senderName, subject, otpMessageToEmail)
	response.JSON(c, http.StatusCreated, otpMessageToEmail)
}
