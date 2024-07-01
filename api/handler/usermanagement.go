package handler

import (
	user "api-gateway/genproto/UserManagementService"
	"log"

	"github.com/gin-gonic/gin"
)

type UserManagementHandler interface {
	GetUserByID(c *gin.Context)
	UpdateUserByID(c *gin.Context)
	DeleteUserByID(c *gin.Context)
	GetUserProfileById(c *gin.Context)
	UpdateUserProfileById(c *gin.Context)
}

type userManagementHandler struct {
	usermanagement user.UserManagementServiceClient
}

func NewUserManagementHandler(usermanagement user.UserManagementServiceClient) UserManagementHandler {
	return &userManagementHandler{usermanagement: usermanagement}
}

func (h *userManagementHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var req = user.IdUserRequest{
		UserId: id,
	}

	res, err := h.usermanagement.GetUserById(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *userManagementHandler) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")

	var req = user.UpdateUserRequest{
		UserId: id,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.usermanagement.UpdateUserById(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *userManagementHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	var req = user.IdUserRequest{
		UserId: id,
	}

	res, err := h.usermanagement.DeleteUserById(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *userManagementHandler) GetUserProfileById(c *gin.Context) {
	id := c.Param("id")
	var req = user.IdUserRequest{
		UserId: id,
	}

	res, err := h.usermanagement.GetUserProfileById(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *userManagementHandler) UpdateUserProfileById(c *gin.Context) {
	id := c.Param("id")

	var req = user.UserProfileRequest{
		UserId: id,
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.usermanagement.UpdateUserProfileById(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}
