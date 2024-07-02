package handler

import (
	garden "api-gateway/genproto/GardenManagementService"
	"log"

	"github.com/gin-gonic/gin"
)

type GardenManagementHandler interface {
	CreateGarden(c *gin.Context)
	GetGardenByID(c *gin.Context)
	UpdateGardenByID(c *gin.Context)
	DeleteGardenByID(c *gin.Context)
	GetGardensByUserID(c *gin.Context)
	CreatePlantByGardenID(c *gin.Context)
	GetPlantsByGardenID(c *gin.Context)
	UpdatePlantByPlantsID(c *gin.Context)
	DeletePlantByPlantsID(c *gin.Context)
	CreateCareLogByPlantID(c *gin.Context)
	GetCareLogsByPlantID(c *gin.Context)
}

type gardenManagementHandler struct {
	gardenManagemen garden.GardenManagementServiceClient
}

func NewGardenManagementHandler(gardenManagemen garden.GardenManagementServiceClient) GardenManagementHandler {
	return &gardenManagementHandler{gardenManagemen: gardenManagemen}
}

// 1
func (h *gardenManagementHandler) CreateGarden(c *gin.Context) {
	var req garden.GardenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	res, err := h.gardenManagemen.CreateGarden(c, &req)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 2
func (h *gardenManagementHandler) GetGardenByID(c *gin.Context) {
	id := c.Param("id")
	var req = garden.IdRequest{
		Id: id,
	}

	res, err := h.gardenManagemen.GetGardenByID(c, &req)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 3
func (h *gardenManagementHandler) UpdateGardenByID(c *gin.Context) {
	id := c.Param("id")
	var req = garden.UpdateGardenRequest{
		Id: id,
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.gardenManagemen.UpdateGardenByID(c, &req)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 4
func (h *gardenManagementHandler) DeleteGardenByID(c *gin.Context) {
	id := c.Param("id")
	var req = garden.IdRequest{
		Id: id,
	}

	res, err := h.gardenManagemen.DeleteGardenByID(c, &req)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 5
func (h *gardenManagementHandler) GetGardensByUserID(c *gin.Context) {
	id := c.Param("id")
	var req = garden.IdRequest{
		Id: id,
	}

	res, err := h.gardenManagemen.GetGardensByUserID(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 6
func (h *gardenManagementHandler) CreatePlantByGardenID(c *gin.Context) {
	var req garden.PlantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.gardenManagemen.CreatePlantByGardenID(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 7
func (h *gardenManagementHandler) GetPlantsByGardenID(c *gin.Context) {
	id := c.Param("id")
	var req = garden.IdRequest{
		Id: id,
	}

	res, err := h.gardenManagemen.GetPlantsByGardenID(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 8
func (h *gardenManagementHandler) UpdatePlantByPlantsID(c *gin.Context) {
	id := c.Param("id")
	var req = garden.PlantRequest{
		GardenId: id,
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.gardenManagemen.UpdatePlantByPlantsID(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 9
func (h *gardenManagementHandler) DeletePlantByPlantsID(c *gin.Context) {
	id := c.Param("id")
	var req = garden.IdRequest{
		Id: id,
	}

	res, err := h.gardenManagemen.DeletePlantByPlantsID(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 10
func (h *gardenManagementHandler) CreateCareLogByPlantID(c *gin.Context) {
	var req garden.CareLogs
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.gardenManagemen.CreateCareLogByPlantID(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 11
func (h *gardenManagementHandler) GetCareLogsByPlantID(c *gin.Context) {
	id := c.Param("id")
	var req = garden.IdRequest{
		Id: id,
	}

	res, err := h.gardenManagemen.GetCareLogsByPlantID(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}
