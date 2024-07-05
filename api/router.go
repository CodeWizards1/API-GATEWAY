package api

import (
	"api-gateway/api/handler"

	"github.com/gin-gonic/gin"
)

func New(server *handler.Server) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handeler := handler.NewHandlerConfig(server)

	user := router.Group("/api/user")
	{
		user.POST("/register", handeler.Usermanagement.CreateUser)
		user.POST("/login", handeler.Usermanagement.Login)
		user.GET("/:id", handeler.Usermanagement.GetUserByID)
		user.PUT("/:id", handeler.Usermanagement.UpdateUserByID)
		user.DELETE("/:id", handeler.Usermanagement.DeleteUserByID)
		user.GET("/profile/:id", handeler.Usermanagement.GetUserProfileById)
		user.PUT("/profile/:id", handeler.Usermanagement.UpdateUserProfileById)
	}

	garden := router.Group("/api/garden")
	{
		garden.POST("/create", handeler.Gargardenmanagement.CreateGarden)
		garden.GET("/:id", handeler.Gargardenmanagement.GetGardenByID)
		garden.PUT("/:id", handeler.Gargardenmanagement.UpdateGardenByID)
		garden.DELETE("/:id", handeler.Gargardenmanagement.DeleteGardenByID)
		garden.GET("/user/:id", handeler.Gargardenmanagement.GetGardensByUserID)
		garden.POST("/plant", handeler.Gargardenmanagement.CreatePlantByGardenID)
		garden.GET("/plant/:id", handeler.Gargardenmanagement.GetPlantsByGardenID)
		garden.PUT("/plant/:id", handeler.Gargardenmanagement.UpdatePlantByPlantsID)
		garden.DELETE("/plant/:id", handeler.Gargardenmanagement.DeletePlantByPlantsID)
		garden.POST("/plant/carelog", handeler.Gargardenmanagement.CreateCareLogByPlantID)
		garden.GET("/plant/carelog/:id", handeler.Gargardenmanagement.GetCareLogsByPlantID)
	}

	community := router.Group("/api")
	{
		community.POST("/communities", handeler.Community.CreateCommunity)
		community.GET("/communities/:id", handeler.Community.GetCommunityBy)
		community.PUT("/communities/:id", handeler.Community.UpdateCommunity)
		community.DELETE("/communities/:id", handeler.Community.DeleteCommunity)
		community.GET("/communities", handeler.Community.GetAllCommunity)
		community.POST("/communities/join/:id", handeler.Community.JoinCommunity)
		community.POST("/communities/leave/:id", handeler.Community.LeaveCommunity)
		community.POST("/communities/:id/events", handeler.Community.CreateCommunityEvent)
		community.GET("/communities/:id/events", handeler.Community.GetCommunityEventBy)
	}

	sustainability := router.Group("/api/sustainability")
	{
		sustainability.POST("/impact/log", handeler.Sustainability.LogImpact)
		sustainability.GET("/users/:id/impact", handeler.Sustainability.GetUserImpact)
		sustainability.GET("/communities/:id/impact", handeler.Sustainability.GetCommunityImpact)
		sustainability.GET("/challenges", handeler.Sustainability.GetChallenges)
		sustainability.POST("/challenges/:id/join", handeler.Sustainability.JoinChallenge)
		sustainability.PUT("/challenges/:id/progress", handeler.Sustainability.UpdateChallengeProgress)
		sustainability.GET("/users/:id/challenges", handeler.Sustainability.GetUserChallenges)
		sustainability.GET("/leaderboard/users", handeler.Sustainability.GetUserLeaderboard)
		sustainability.GET("/leaderboard/communities", handeler.Sustainability.GetCommunityLeaderboard)
		sustainability.POST("/post/challange", handeler.Sustainability.PostChallenges)
	}

	return router
}
