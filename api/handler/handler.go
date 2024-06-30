package handler

import "api-gateway/genproto/AuthenticationSevice/authentication"

type handler struct {
	authentication authentication.AuthenticationServiceClient
}

type HandlerConfig struct {
}

func New(authentication authentication.AuthenticationServiceClient) *handler {
	return &handler{authentication: authentication}
}
