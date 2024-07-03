package handler

import (
	com "api-gateway/genproto/CommunityService"

	"github.com/gin-gonic/gin"
)

type CommunityHandler interface {
	CreateCommunity(c *gin.Context)
	GetCommunityBy(c *gin.Context)
	UpdateCommunity(c *gin.Context)
	DeleteCommunity(c *gin.Context)
	GetAllCommunity(c *gin.Context)
	JoinCommunity(c *gin.Context)
	LeaveCommunity(c *gin.Context)
	CreateCommunityEvent(c *gin.Context)
	GetCommunityEventBy(c *gin.Context)
}

type communityHandler struct {
	community com.CommunityServiceClient
}

func NewCommunityHandler(community com.CommunityServiceClient) CommunityHandler {
	return &communityHandler{community: community}
}

// 1
func (h *communityHandler) CreateCommunity(c *gin.Context) {
	var req com.Community
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.community.CreateCommunity(c, &com.CreateCommunityRequest{Community: &req})
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 2
func (h *communityHandler) GetCommunityBy(c *gin.Context) {
	id := c.Param("id")

	req := com.GetCommunityRequest{Id: id}

	res, err := h.community.GetCommunityBy(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 3
func (h *communityHandler) UpdateCommunity(c *gin.Context) {
	var req com.Community
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.community.UpdateCommunity(c, &com.UpdateCommunityRequest{Community: &req})
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 4
func (h *communityHandler) DeleteCommunity(c *gin.Context) {
	id := c.Param("id")
	req := com.DeleteCommunityRequest{Id: id}

	res, err := h.community.DeleteCommunity(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 5
func (h *communityHandler) GetAllCommunity(c *gin.Context) {
	var req com.GetAllCommunityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.community.GetAllCommunity(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 6
func (h *communityHandler) JoinCommunity(c *gin.Context) {
	var req com.JoinCommunityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.community.JoinCommunity(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 7
func (h *communityHandler) LeaveCommunity(c *gin.Context) {
	var req com.LeaveCommunityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.community.LeaveCommunity(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 8
func (h *communityHandler) CreateCommunityEvent(c *gin.Context) {
	var req com.Event
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.community.CreateCommunityEvent(c, &com.CreateCommunityEventRequest{Event: &req})
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 9
func (h *communityHandler) GetCommunityEventBy(c *gin.Context) {
	id := c.Param("id")
	req := com.GetCommunityEventRequest{Id: id}

	res, err := h.community.GetCommunityEvent(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}
