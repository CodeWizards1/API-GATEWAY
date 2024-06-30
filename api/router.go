package api

import (
	"api-gateway/api/handler"
	"api-gateway/genproto/AuthenticationSevice/authentication"
	"api-gateway/genproto/UserManagementSevice/user"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Option struct{}

func New(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authenticationClient := authentication.NewAuthenticationServiceClient(conn)
	usermanagementClient := user.NewUserManagementServiceClient(conn)

	authHandler, userHandler := handler.New(authenticationClient, usermanagementClient)

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
	}

	return router
}
