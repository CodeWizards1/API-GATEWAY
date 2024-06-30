package handler

import (
	"api-gateway/genproto/AuthenticationSevice/authentication"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}

type authenticationHandler struct {
	authentication authentication.AuthenticationServiceClient
}

func NewAuthenticationHandler(authentication authentication.AuthenticationServiceClient) AuthenticationHandler {
	return &authenticationHandler{authentication: authentication}
}

func (h *authenticationHandler) CreateUser(c *gin.Context) {
	var req authentication.UserRequest

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
	var req authentication.AutorizationRequest

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
