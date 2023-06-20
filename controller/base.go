package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"develov_be/middleware"
	"develov_be/models"
	"develov_be/response"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Server) ConfigOtpGmail(configAuthEmail, configAuthPassword, configSmtpHost, configSmtpPort string, toEmail []string, body []byte) error {
	auth := smtp.PlainAuth("", configAuthEmail, configAuthPassword, configSmtpHost)
	smtpAddr := fmt.Sprintf("%s:%v", configSmtpHost, configSmtpPort)

	err := smtp.SendMail(smtpAddr, auth, configAuthEmail, toEmail, body)
	if err != nil {
		return err
	}
	return nil
}

// config to db local
func (s *Server) InitializeServer(DbDriver, DbHost, DbUser, DbPassword, DbName, DbPort string) {
	var err error
	if DbDriver == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		s.DB, err = gorm.Open(DbDriver, dsn)
		if err != nil {
			fmt.Print("not connected to database", DbDriver)
			log.Fatal("This The Error :", err)
		} else {
			fmt.Printf("connected to the %v database", DbDriver)
		}
	}

	//auto migrate
	s.DB.AutoMigrate(
		models.User{},
		models.Mentor{},
		models.FeedBack{},
		models.Comment{},
		models.Community{},
		models.Tools{},
		models.MentorTools{},
		models.Schedule{},
		models.Payment{},
		models.Ticket{},
	)

	gin.SetMode(gin.ReleaseMode)
	// s.Router = gin.New()
	s.Router = gin.Default()
	s.InitializeRoutes()
	s.Router.Use(middleware.CORSMiddleware())
}

// status on server
func (s *Server) StatusServer(c *gin.Context) {
	response.GenericJsonResponse(c, http.StatusOK, "Server Develov Is Running")
}

func (s *Server) RunServer(addr string) {

	fmt.Println("Listen Of Port Server : " + os.Getenv("PORT"))
	handler := s.Router
	log.Fatal(http.ListenAndServe(addr, handler))
}
