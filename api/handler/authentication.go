package handler

import (
	"api-gateway/genproto/AuthenticationSevice/authentication"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateUser(c *gin.Context) {
	var req authentication.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authentication.CreateUser(c, &req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) Login(c *gin.Context) {
	var req authentication.AutorizationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authentication.Login(c, &req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}
