package handler

import (
	pb "api-gateway/genproto/SustainabilityService"
	"log"

	"github.com/gin-gonic/gin"
)

type SustainabilityHandler interface {
	LogImpact(c *gin.Context)
	GetUserImpact(c *gin.Context)
	GetCommunityImpact(c *gin.Context)
	GetChallenges(c *gin.Context)
	JoinChallenge(c *gin.Context)
	UpdateChallengeProgress(c *gin.Context)
	GetUserChallenges(c *gin.Context)
	GetUserLeaderboard(c *gin.Context)
	GetCommunityLeaderboard(c *gin.Context)
	PostChallenges(c *gin.Context)
}

type sustainabilityService struct {
	sustainability pb.SustainabilityServiceClient
}

func NewSustainabilityHandler(sustainability pb.SustainabilityServiceClient) SustainabilityHandler {
	return &sustainabilityService{sustainability: sustainability}
}

// 1
func (s *sustainabilityService) LogImpact(c *gin.Context) {
	var req pb.LogImpactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	res, err := s.sustainability.LogImpact(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, res)
}

// 2
func (s *sustainabilityService) GetUserImpact(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetUserImpactRequest{
		UserId: id,
	}

	res, err := s.sustainability.GetUserImpact(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 3
func (s *sustainabilityService) GetCommunityImpact(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetCommunityImpactRequest{
		CommunityId: id,
	}
	res, err := s.sustainability.GetCommunityImpact(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, res)
}

// 4
func (s *sustainabilityService) GetChallenges(c *gin.Context) {
	res, err := s.sustainability.GetChallenges(c, &pb.GetChallengesRequest{})
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, res)
}

// 5
func (s *sustainabilityService) JoinChallenge(c *gin.Context) {
	challengeID := c.Param("id")
	req := pb.JoinChallengeRequest{
		ChallengeId: challengeID,
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := s.sustainability.JoinChallenge(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 6
func (s *sustainabilityService) UpdateChallengeProgress(c *gin.Context) {
	challengeID := c.Param("id")
	req := pb.UpdateChallengeProgressRequest{
		ChallengeId: challengeID,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := s.sustainability.UpdateChallengeProgress(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, res)
}

// 7
func (s *sustainabilityService) GetUserChallenges(c *gin.Context) {
	userID := c.Param("id")
	req := pb.GetUserChallengesRequest{
		UserId: userID,
	}

	res, err := s.sustainability.GetUserChallenges(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 8
func (s *sustainabilityService) GetUserLeaderboard(c *gin.Context) {
	res, err := s.sustainability.GetUserLeaderboard(c, &pb.GetUserLeaderboardRequest{})

	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 9
func (s *sustainabilityService) GetCommunityLeaderboard(c *gin.Context) {
	res, err := s.sustainability.GetCommunityLeaderboard(c, &pb.GetCommunityLeaderboardRequest{})
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

// 10
func (s *sustainabilityService) PostChallenges(c *gin.Context) {
	var req pb.PostChallengesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := s.sustainability.PostChallenges(c, &req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}
