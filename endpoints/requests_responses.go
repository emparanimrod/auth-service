package endpoints

import "github.com/gofrs/uuid"

type UserRegistrationRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	PIN       string `json:"pin"`
	UserType  string `json:"usertype"`
}

type UserRegistrationResponse struct {
	AuthID uuid.UUID `json:"authId"`
}

type UserEmailSigninRequest struct {
	Email string `json:"email"`
	PIN   string `json:"pin"`
}

type UserPhoneSigninRequest struct {
	Phone string `json:"phone"`
	PIN   string `json:"pin"`
}

type userSigninResponse struct {
	AuthID uuid.UUID
}

type SignedUserResponse struct {
	AuthID string `json:"authId"`
	// UserType UserType `json:"userType"`
	Token string `json:"token"`
}

type TokenValidationRequest struct {
	Token string `json:"token"`
}

type TokenValidationResponse struct {
	AuthID   string `json:"authId"`
	// UserType string `json:"userType"`
}
