package handler

import (
	au "api-gateway/genproto/AuthentificationService"
	us "api-gateway/genproto/UserManagementService"
	gar "api-gateway/genproto/GardenManagementService"
)

type HandlerConfig struct {
}

func New(authentication au.AuthenticationServiceClient, usermanagement us.UserManagementServiceClient, gargardenmanagement gar.GardenManagementServiceClient) (AuthenticationHandler, UserManagementHandler, GardenManagementHandler) {
	return NewAuthenticationHandler(authentication), NewUserManagementHandler(usermanagement), NewGardenManagementHandler(gargardenmanagement)
}
