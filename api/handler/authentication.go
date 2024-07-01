package handler

import (
	auth "api-gateway/genproto/AuthentificationService"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}

type authenticationHandler struct {
	authentication auth.AuthenticationServiceClient
}

func NewAuthenticationHandler(authentication auth.AuthenticationServiceClient) AuthenticationHandler {
	return &authenticationHandler{authentication: authentication}
}

func (h *authenticationHandler) CreateUser(c *gin.Context) {
	var req auth.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authentication.CreateUser(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *authenticationHandler) Login(c *gin.Context) {
	var req auth.AutorizationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authentication.Login(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}
