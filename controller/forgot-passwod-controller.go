package controller

import (
	"develov_be/auth"
	"develov_be/models"
	"develov_be/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
)

// chek email &n send auto migrate otp to email
func (s *Server) ChekEmailUserController(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	_, err := user.CekEmail(s.DB, user.Email)

	if err != nil {
		response.ErrorResponseAuth(c, http.StatusUnprocessableEntity, "maaf email anda tidak terdaftar")
		return
	} else {
		otp := s.SendOtpUser(user.Email)

		token := auth.ExtractToken(c)
		_, err := user.UpdateOtp(s.DB, token, otp)

		if err != nil {
			response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
			return
		}

		response.JSON(c, http.StatusCreated, otp)
	}
}

// chek otp
func (s *Server) CheckOtpController(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	_, err := user.CekOtp(s.DB, user.Otp)

	if err != nil {
		response.ErrorResponseAuth(c, http.StatusUnprocessableEntity, "Otp Tidak Cocok")
		return
	}

	response.JSON(c, http.StatusCreated, "Otp Cocok")
}

// change password
func (s *Server) ChangePasswordController(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}
	token := auth.ExtractToken(c)

	hashPassword := models.HashPasswordToSha256(user.Password)

	_, err := user.UpdatePassword(s.DB, token, hashPassword)

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(c, http.StatusCreated, hashPassword)
}

// send otp
func (s *Server) SendOtpUser(email string) string {

	secretLength := 16
	otpAutoGenerate := gotp.RandomSecret(secretLength)
	newOtp := gotp.NewDefaultTOTP(otpAutoGenerate).Now()

	toEmailUserSendEmail := []string{email}
	subject := "PT.DEVELOV"
	otpMessageToEmail := "Hii\n OTP Anda Adalah: " + newOtp
	senderName := "PT.DEVELOV"

	s.SendOtpToEmail(toEmailUserSendEmail, senderName, subject, otpMessageToEmail)

	return newOtp
}
