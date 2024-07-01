package handler

import (
	au "api-gateway/genproto/AuthentificationService"
	us "api-gateway/genproto/UserManagementService"
)

type HandlerConfig struct {
}

func New(authentication au.AuthenticationServiceClient, usermanagement us.UserManagementServiceClient) (AuthenticationHandler, UserManagementHandler) {
	return NewAuthenticationHandler(authentication), NewUserManagementHandler(usermanagement)
}
