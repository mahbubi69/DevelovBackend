package controller

import "develov_be/middleware"

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
		v1.DELETE("/delete/:id", s.DeleteUserController)
		// forgot password
		v1.POST("/chekEmailUser", s.ChekEmailUserController)
		v1.PUT("/changePasswordUser", s.ChangePasswordController)
		v1.POST("/checkOtp", s.CheckOtpController)

		//mentor
		v1.POST("/createdMentor", s.CreatedMentorController)
		v1.GET("/feedBack", s.GetAllMentorController)
		v1.DELETE("/deleteMentor/:id", s.DeleteMentorController)
		//feedback mentor
		v1.POST("/createdFeedBack", s.CreatedFeedBackController)

		//community
		v1.POST("/createdCommunity", s.CreatedCommunityController)
		v1.GET("/community", middleware.SetMiddlewareAuthencation(s.GetAllCommunityController))
		//comment community
		v1.POST("/createdComment", s.CreatedCommentCommunityController)
		v1.POST("/searchCommunity", s.SearchCommunityByTitleController)

		//tools
		v1.POST("/createdTools", s.CreatedToolsController)
		v1.GET("/tools", s.ReadToolsController)
		v1.PUT("/updateTools", s.UpdateToolsController)

		v1.POST("/createdMentorTools/:idMentor", s.CeatedMentorToolsController)

	}

}
