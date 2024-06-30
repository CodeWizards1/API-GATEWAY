package api

import (
	"api-gateway/api/handler"
	"api-gateway/genproto/AuthenticationSevice/authentication"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Option struct{}

func New(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authenticationClient := authentication.NewAuthenticationServiceClient(conn)

	handler := handler.New(authenticationClient)

	crud := router.Group("/")
	{
		crud.POST("/api/auth/register", handler.CreateUser)
		crud.POST("/api/auth/login", handler.Login)
	}

	return router
}
