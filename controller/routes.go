package controller

func (s *Server) InitializeRoutes() {
	//how to version api group
	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/status", s.StatusServer)

		//user
		v1.POST("/register", s.CreatedUserController)
		v1.POST("/login", s.LoginUserController)
		v1.GET("/user", s.GetAllUserMappingController)
		v1.GET("/profile", s.GetUserByTokenController)
		v1.GET("assets/images/:file", s.ReadImagesController)
		v1.PUT("/updateImage", s.UpdateImageController)
		v1.PUT("/deleteImage", s.DeleteImageUserController)
		v1.DELETE("/delete/:id", s.DeleteUserController)
		// forgot password
		v1.POST("/chekEmailUser", s.ChekEmailUserController)
		v1.PUT("/changePasswordUser", s.ChangePasswordController)
		v1.POST("/checkOtp", s.CheckOtpController)

		//mentor
		v1.POST("/createdMentor", s.CreatedMentorController)
		v1.GET("/feedBack", s.GetAllMentorController)
		v1.PUT("/updateMentor/:id", s.UpdateMentorController)
		v1.DELETE("/deleteMentor/:id", s.DeleteMentorController)
		//feedback mentor
		v1.POST("/createdFeedBack", s.CreatedFeedBackController)

		//community
		v1.POST("/createdCommunity", s.CreatedCommunityController)
		v1.GET("/community", s.GetAllCommunityController)
		//comment community
		v1.POST("/createdComment", s.CreatedCommentCommunityController)
		v1.POST("/searchCommunity", s.SearchCommunityByTitleController)

		//tools
		v1.POST("/createdTools", s.CreatedToolsController)
		v1.GET("/tools", s.ReadToolsController)
		v1.PUT("/updateTool/:id", s.UpdateToolsController)
		//mentor tools
		v1.POST("/createdMentorTools/:idMentor", s.CeatedMentorToolsController)

		// schedule
		v1.POST("/createdSchedule", s.CreatedScheduleController)
		v1.GET("/schedule", s.GetAllScheduleController)
		v1.PUT("/updateSchedule/:id", s.UpdateScheduleController)

		// payement
		v1.POST("/createdPayment", s.CreatedPaymentController)
		v1.GET("/payment", s.GetAllPaymentController)

		// tiket
		v1.POST("/createdTicket/:iduser", s.CreatedTicketController)
		v1.PUT("/updateTicket/:id", s.UpdateTictketController)
	}

}
