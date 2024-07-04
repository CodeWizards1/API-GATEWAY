package api

import (
	"api-gateway/api/handler"
	auth "api-gateway/genproto/AuthentificationService"
	garden "api-gateway/genproto/GardenManagementService"
	sustain "api-gateway/genproto/SustainabilityService"
	user "api-gateway/genproto/UserManagementService"
	com "api-gateway/genproto/CommunityService"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Option struct{}

func New(conn1, conn2, conn3, conn4, conn5 *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authenticationClient := auth.NewAuthenticationServiceClient(conn1)
	usermanagementClient := user.NewUserManagementServiceClient(conn2)
	gardenmanagementClient := garden.NewGardenManagementServiceClient(conn3)
	sustainabilityClient := sustain.NewSustainabilityServiceClient(conn4)
	communityClient := com.NewCommunityServiceClient(conn5)

	authHandler, userHandler, gardenHandler, sustainabilityHandler, communityHandler := handler.New(authenticationClient, usermanagementClient, gardenmanagementClient, sustainabilityClient, communityClient)

	crud := router.Group("/")
	{
		// Authentication
		crud.POST("/api/auth/register", authHandler.CreateUser)
		crud.POST("/api/auth/login", authHandler.Login)

		// User Management
		crud.GET("/api/user/:id", userHandler.GetUserByID)
		crud.PUT("/api/user/:id", userHandler.UpdateUserByID)
		crud.DELETE("/api/user/:id", userHandler.DeleteUserByID)
		crud.GET("/api/user/profile/:id", userHandler.GetUserProfileById)
		crud.PUT("/api/user/profile/:id", userHandler.UpdateUserProfileById)

		// Garden Management
		crud.POST("/api/garden", gardenHandler.CreateGarden)
		crud.GET("/api/garden/:id", gardenHandler.GetGardenByID)
		crud.PUT("/api/garden/:id", gardenHandler.UpdateGardenByID)
		crud.DELETE("/api/garden/:id", gardenHandler.DeleteGardenByID)
		crud.GET("/api/garden/user/:id", gardenHandler.GetGardensByUserID)
		crud.POST("/api/garden/plant", gardenHandler.CreatePlantByGardenID)
		crud.GET("/api/garden/plant/:id", gardenHandler.GetPlantsByGardenID)
		crud.PUT("/api/garden/plant/:id", gardenHandler.UpdatePlantByPlantsID)
		crud.DELETE("/api/garden/plant/:id", gardenHandler.DeletePlantByPlantsID)
		crud.POST("/api/garden/carelog", gardenHandler.CreateCareLogByPlantID)
		crud.GET("/api/garden/carelog/:id", gardenHandler.GetCareLogsByPlantID)

		// Community Management
		crud.POST("/api/communities", communityHandler.CreateCommunity)
		crud.GET("/api/communities/:id", communityHandler.GetCommunityBy)
		crud.PUT("/api/communities/:id", communityHandler.UpdateCommunity)
		crud.DELETE("/api/communities/:id", communityHandler.DeleteCommunity)
		crud.GET("/api/communities", communityHandler.GetAllCommunity)
		crud.POST("/api/communities/join/:id", communityHandler.JoinCommunity)
		crud.POST("/api/communities/leave/:id", communityHandler.LeaveCommunity)
		crud.POST("/api/communities/:id/events", communityHandler.CreateCommunityEvent)
		crud.GET("/api/communities/:id/events", communityHandler.GetCommunityEventBy)

		// Sustainability
		crud.POST("/api/impact/log", sustainabilityHandler.LogImpact)
		crud.GET("/api/users/:id/impact", sustainabilityHandler.GetUserImpact)
		crud.GET("/api/communities/:id/impact", sustainabilityHandler.GetCommunityImpact)
		crud.GET("/api/challenges", sustainabilityHandler.GetChallenges)
		crud.POST("/api/challenges/:id/join", sustainabilityHandler.JoinChallenge)
		crud.PUT("/api/challenges/:id/progress", sustainabilityHandler.UpdateChallengeProgress)
		crud.GET("/api/users/:id/challenges", sustainabilityHandler.GetUserChallenges)
		crud.GET("/api/leaderboard/users", sustainabilityHandler.GetUserLeaderboard)
		crud.GET("/api/leaderboard/communities", sustainabilityHandler.GetCommunityLeaderboard)
		crud.POST("/api/post/challange", sustainabilityHandler.PostChallenges)

	}

	return router
}
