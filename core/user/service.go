package user

import (
	"auth/core/errors"
	"auth/helpers"

	"github.com/gofrs/uuid"
)

type Service interface {
	Register(firstname, lastname, email, phone, pin string) (User, error)

	AuthenticateByEmail(email, pin string) (User, error)
	AuthenticateByPhone(phone, pin string) (User, error)
}

func NewService(repository Repository) Service {
	return &service{repository}
}

type service struct {
	repository Repository
}

// Register adds a user to db if not already exist and returns this user
func (svc service) Register(firstname, lastname, email, phone, pin string) (User, error) {

	// hash user pin before adding to db.
	pinHash, err := helpers.HashPassword(pin)
	if err != nil { // if we get an error, it means our hashing func dint work
		return User{}, errors.Error{Err: err, Code: errors.EINTERNAL}
	}

	id, _ := uuid.NewV4()
	user := User{
		AuthID:    id,
		Phone:     phone,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		PIN:       pinHash,
	}
	usr, err := svc.repository.Add(user)
	if err != nil {
		return User{}, err
	}

	return usr, nil
}

// AuthenticateByEmail verifies a user by given email and pin
func (svc service) AuthenticateByEmail(email, pin string) (User, error) {
	// 	search for user by email
	user, err := svc.repository.GetByEmail(email)
	if errors.ErrorCode(err) == errors.ENOTFOUND {
		return User{}, errors.Error{Err: err, Message: errors.ErrUserNotFound}
	} else if err != nil {
		return User{}, err
	}

	// validate password
	if err := helpers.ComparePasswordToHash(user.PIN, pin); err != nil {
		return User{}, errors.Unauthorized{Message: errors.ErrorMessage(err)}
	}

	return user, nil
}

// AuthenticateByPhone verifies a user by given phone and pin
func (svc service) AuthenticateByPhone(phone, pin string) (User, error) {
	// 	search for user by phone
	user, err := svc.repository.GetByPhone(phone)
	if errors.ErrorCode(err) == errors.ENOTFOUND {
		return User{}, errors.Error{Err: err, Message: errors.ErrUserNotFound}
	} else if err != nil {
		return User{}, err
	}

	// validate password
	if err := helpers.ComparePasswordToHash(user.PIN, pin); err != nil {
		return User{}, errors.Unauthorized{Message: errors.ErrorMessage(err)}
	}

	return user, nil
}
