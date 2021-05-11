package grpc_service

import (
	"context"

	"auth/endpoints"
	pbauth "auth/pb/auth"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type registrationService struct {
	register grpctransport.Handler
}

func (svc registrationService) Register(ctx context.Context, req *pbauth.UserRegistrationRequest) (*pbauth.UserRegistrationReply, error) {
	_, rep, err := svc.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pbauth.UserRegistrationReply), nil
}

type authenticationService struct {
	emailSignIn   grpctransport.Handler
	phoneSignIn   grpctransport.Handler
	tokenValidate grpctransport.Handler
}

func (svc authenticationService) AuthenticateToken(ctx context.Context, req *pbauth.UserTokenValidateRequest) (*pbauth.UserTokenValidateReply, error) {
	_, rep, err := svc.tokenValidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pbauth.UserTokenValidateReply), nil
}

func (svc authenticationService) EmailSignIn(ctx context.Context, req *pbauth.UserEmailSignInRequest) (*pbauth.UserSignInReply, error) {
	_, rep, err := svc.emailSignIn.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pbauth.UserSignInReply), nil
}

func (svc authenticationService) PhoneSignIn(ctx context.Context, req *pbauth.UserPhoneSignInRequest) (*pbauth.UserSignInReply, error) {
	_, rep, err := svc.phoneSignIn.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pbauth.UserSignInReply), nil
}

func NewRegistrationService(userSvcEndpoints endpoints.UserServiceEndpoints) pbauth.RegistrationServiceServer {
	return &registrationService{
		register: grpctransport.NewServer(
			userSvcEndpoints.Register,
			decodeRegistrationRequest,
			encodeRegistrationResponse,
		)}
}

func NewAuthenticationService(userSvcEndpoints endpoints.UserServiceEndpoints) pbauth.AuthenticationServiceServer {
	return &authenticationService{
		emailSignIn: grpctransport.NewServer(
			userSvcEndpoints.SignInByEmail,
			decodeEmailSignInRequest,
			encodeSignInResponse,
		),
		phoneSignIn: grpctransport.NewServer(
			userSvcEndpoints.SignInByPhone,
			decodePhoneSignInRequest,
			encodeSignInResponse,
		),
		tokenValidate: grpctransport.NewServer(
			userSvcEndpoints.ValidateToken,
			decodeTokenValidateRequest,
			encodeTokenValidateResponse,
		),
	}
}
