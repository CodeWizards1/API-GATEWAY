package handler

import (
	user "api-gateway/genproto/UserManagementService"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret-key")

type Claims struct {
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}

type UserManagementHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
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

func (h *userManagementHandler) Login(c *gin.Context) {
	var req user.AutorizationRequest
	req.Email = c.GetHeader("email")
	req.Password = c.GetHeader("password")

	fmt.Println(req.Email, req.Password)
	res, err := h.usermanagement.Login(c, &req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "failed to login: " + err.Error()})
		return
	}
	fmt.Println(res.Message)
	if res.Message == "Login successful" {
		exprTime := time.Now().Add(time.Hour)
		claims := &Claims{
			Email: req.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: exprTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		strToken, err := token.SignedString(jwtKey)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"token": strToken})
	} else {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password not valid"})
	}
}

func (h *userManagementHandler) CreateUser(c *gin.Context) {
	var req user.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.usermanagement.CreateUser(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
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

	var req = user.UpdateUserProfileRequest{
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
