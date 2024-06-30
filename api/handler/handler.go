package handler

import (
	au "api-gateway/genproto/AuthenticationSevice/authentication"
	us "api-gateway/genproto/UserManagementSevice/user"
)

type HandlerConfig struct {
}

func New(authentication au.AuthenticationServiceClient, usermanagement us.UserManagementServiceClient) (AuthenticationHandler, UserManagementHandler) {
	return NewAuthenticationHandler(authentication), NewUserManagementHandler(usermanagement)
}
