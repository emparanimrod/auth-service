package endpoints

import (
	"context"

	"auth/core"
	"auth/core/auth"
	"auth/core/errors"
	"auth/core/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
)

type UserServiceEndpoints struct {
	Register      endpoint.Endpoint
	SignInByEmail endpoint.Endpoint
	SignInByPhone endpoint.Endpoint

	ValidateToken endpoint.Endpoint
}

// Authentication acts as a go-kit endpoint middleware to create a jwt token for
// an authenticated user
func Authentication(tokenDuration uint, secretKey string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			resp, err := next(ctx, request)
			if err != nil {
				return nil, err
			}

			// generate an auth token string
			signInResp := resp.(userSigninResponse)
			token, err := auth.GetTokenString(signInResp.AuthID, tokenDuration, secretKey)
			if err != nil {
				return nil, err
			}

			return SignedUserResponse{
				AuthID: signInResp.AuthID.String(),
				Token:  token,
			}, nil
		}
	}
}

func MakeUserServiceEndpoints(service user.Service, config core.Config) UserServiceEndpoints {
	return UserServiceEndpoints{
		Register:      makeUserRegistrationEndpoint(service),
		SignInByEmail: Authentication(config.TokenDuration, config.Secret)(makeUserSignInByEmailEndpoint(service)),
		SignInByPhone: Authentication(config.TokenDuration, config.Secret)(makeUserSignInByPhoneEndpoint(service)),
		ValidateToken: makeTokenValidationEndpoint(config),
	}
}

func makeTokenValidationEndpoint(config core.Config) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TokenValidationRequest)

		var claims auth.TokenClaims
		token, err := auth.ParseToken(req.Token, config.Secret, &claims)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return nil, errors.Unauthorized{Message: "invalid signature on token"}
			}

			return nil, errors.Unauthorized{Message: "token has expired or is invalid"}
		}

		if valid := auth.ValidateToken(token); !valid {
			return nil, errors.Unauthorized{Message: "invalid token"}
		}

		return TokenValidationResponse{AuthID: claims.User.AuthID.String()}, nil
	}
}

func makeUserRegistrationEndpoint(userService user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRegistrationRequest)

		usr, err := userService.Register(req.Firstname, req.Lastname, req.Email, req.Phone, req.PIN)
		if err != nil {
			return nil, err
		}

		return UserRegistrationResponse{AuthID: usr.AuthID}, nil
	}
}

func makeUserSignInByEmailEndpoint(userService user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserEmailSigninRequest)

		usr, err := userService.AuthenticateByEmail(req.Email, req.PIN)
		if err != nil {
			return nil, err
		}

		return userSigninResponse{AuthID: usr.AuthID}, nil
	}
}

func makeUserSignInByPhoneEndpoint(userService user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserPhoneSigninRequest)

		usr, err := userService.AuthenticateByPhone(req.Phone, req.PIN)
		if err != nil {
			return nil, err
		}

		return userSigninResponse{AuthID: usr.AuthID}, nil
	}
}
