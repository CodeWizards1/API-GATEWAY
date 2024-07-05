package handler

import (
	com "api-gateway/genproto/CommunityService"
	gar "api-gateway/genproto/GardenManagementService"
	sustain "api-gateway/genproto/SustainabilityService"
	us "api-gateway/genproto/UserManagementService"

	"google.golang.org/grpc"
)

type Server struct {
	Usermanagement      *grpc.ClientConn
	Gargardenmanagement *grpc.ClientConn
	Sustainability      *grpc.ClientConn
	Community           *grpc.ClientConn
}

type HandlerConfig struct {
	Usermanagement      UserManagementHandler
	Gargardenmanagement GardenManagementHandler
	Sustainability      SustainabilityHandler
	Community           CommunityHandler
}

func NewHandlerConfig(conn *Server) *HandlerConfig {
	return &HandlerConfig{
		Usermanagement:      NewUserManagementHandler(us.NewUserManagementServiceClient(conn.Usermanagement)),
		Gargardenmanagement: NewGardenManagementHandler(gar.NewGardenManagementServiceClient(conn.Gargardenmanagement)),
		Sustainability:      NewSustainabilityHandler(sustain.NewSustainabilityServiceClient(conn.Sustainability)),
		Community:           NewCommunityHandler(com.NewCommunityServiceClient(conn.Community)),
	}
}
