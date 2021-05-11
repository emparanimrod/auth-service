package grpc_service

import (
	"context"

	"auth/endpoints"
	pbauth "auth/pb/auth"
)

func decodeRegistrationRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pbauth.UserRegistrationRequest)
	return endpoints.UserRegistrationRequest{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Phone:     req.Phone,
		Email:     req.Email,
		PIN:       req.Pin,
		UserType:  req.Usertype,
	}, nil
}

func encodeRegistrationResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.UserRegistrationResponse)
	return &pbauth.UserRegistrationReply{AuthId: resp.AuthID.String()}, nil
}

func decodeEmailSignInRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pbauth.UserEmailSignInRequest)
	return endpoints.UserEmailSigninRequest{
		Email: req.Email,
		PIN:   req.Pin,
	}, nil
}

func decodePhoneSignInRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pbauth.UserPhoneSignInRequest)
	return endpoints.UserPhoneSigninRequest{
		Phone: req.Phone,
		PIN:   req.Pin,
	}, nil
}

func encodeSignInResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.SignedUserResponse)
	return &pbauth.UserSignInReply{AuthId: resp.AuthID, Token: resp.Token}, nil
}

func decodeTokenValidateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pbauth.UserTokenValidateRequest)
	return endpoints.TokenValidationRequest{Token: req.GetToken()}, nil
}

func encodeTokenValidateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.TokenValidationResponse)
	return &pbauth.UserTokenValidateReply{AuthId: resp.AuthID}, nil
}