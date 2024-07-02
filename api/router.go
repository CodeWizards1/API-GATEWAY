package api

import (
	"api-gateway/api/handler"
	auth "api-gateway/genproto/AuthentificationService"
	user "api-gateway/genproto/UserManagementService"
	garden "api-gateway/genproto/GardenManagementService"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Option struct{}

func New(conn1, conn2, conn3 *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authenticationClient := auth.NewAuthenticationServiceClient(conn1)
	usermanagementClient := user.NewUserManagementServiceClient(conn2)
	gardenmanagementClient := garden.NewGardenManagementServiceClient(conn3)

	authHandler, userHandler, gardenHandler := handler.New(authenticationClient, usermanagementClient, gardenmanagementClient)

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
	}

	return router
}
