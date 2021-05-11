package grpc_service

import (
	"auth/endpoints"
	pbauth "auth/pb/auth"

	"google.golang.org/grpc"
)

func NewServer(serviceEndpoints *endpoints.ServiceEndpoints) *grpc.Server {
	// create a grpc server
	server := grpc.NewServer()

	// init grpc services
	registrationSvc := NewRegistrationService(serviceEndpoints.UserService)
	authSvc := NewAuthenticationService(serviceEndpoints.UserService)

	// 	register grpc services to our server
	pbauth.RegisterRegistrationServiceServer(server, registrationSvc)
	pbauth.RegisterAuthenticationServiceServer(server, authSvc)

	return server
}
