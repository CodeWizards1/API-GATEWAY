package handler

import (
	au "api-gateway/genproto/AuthentificationService"
	com "api-gateway/genproto/CommunityService"
	gar "api-gateway/genproto/GardenManagementService"
	sustain "api-gateway/genproto/SustainabilityService"
	us "api-gateway/genproto/UserManagementService"
)

type HandlerConfig struct {
}

func New(authentication au.AuthenticationServiceClient,
	usermanagement us.UserManagementServiceClient,
	gargardenmanagement gar.GardenManagementServiceClient,
	sustainability sustain.SustainabilityServiceClient,
	community com.CommunityServiceClient) (AuthenticationHandler,
	UserManagementHandler,
	GardenManagementHandler,
	SustainabilityHandler,
	CommunityHandler) {
	return NewAuthenticationHandler(authentication),
		NewUserManagementHandler(usermanagement),
		NewGardenManagementHandler(gargardenmanagement),
		NewSustainabilityHandler(sustainability),
		NewCommunityHandler(community)
}
