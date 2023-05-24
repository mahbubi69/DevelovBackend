package controller

func (s *Server) InitializeRoutes() {
	//how to version api group
	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/status", s.StatusServer)

		//user
		v1.POST("/register", s.CreatedUserController)
		v1.POST("/login", s.LoginUser)
		// v1.GET("/userm", s.GetAllUser)
		v1.GET("/user", s.GetAllUserMapping)

	}
}
