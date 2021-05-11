package user

import (
	"auth/core/errors"
	"auth/storage"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type Repository interface {
	Add(User) (User, error)
	Delete(User) error
	GetByAuthID(uuid.UUID) (User, error)
	GetByEmail(string) (User, error)
	GetByPhone(string) (User, error)
	Update(User) error
}

// NewRepository creates and returns a new instance of admin repository
func NewRepository(database *storage.Database) Repository {
	return &repository{db: database}
}

type repository struct {
	db *storage.Database
}

func (r repository) searchBy(row User) (User, error) {
	var user User
	result := r.db.Where(row).First(&user)
	// check if no record found.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return User{}, errors.Error{Code: errors.ENOTFOUND}
	}
	if err := result.Error; err != nil {
		return User{}, errors.Error{Err: err, Code: errors.EINTERNAL}
	}

	return user, nil
}

// Add a user if already not in db.
func (r repository) Add(user User) (User, error) {
	// add new user to users table, if query return violation of unique key column,
	// we know that the user with given record exists and return that user instead
	result := r.db.Model(User{}).Create(&user)
	if err := result.Error; err != nil {
		// we check if the error is a postgres unique constraint violation
		if pgerr, ok := err.(*pgconn.PgError); ok && pgerr.Code == "23505" {
			return user, errors.Error{Code: errors.ECONFLICT, Message: errors.ErrUserExists}
		}
		return User{}, errors.Error{Err: result.Error, Code: errors.EINTERNAL}
	}

	return user, nil
}

// Delete an user
func (r repository) Delete(user User) error {
	result := r.db.Delete(&user)
	if result.Error != nil {
		return errors.Error{Err: result.Error, Code: errors.EINTERNAL}
	}
	return nil
}

// GetByAuthID searches user by provided authentication id
func (r repository) GetByAuthID(id uuid.UUID) (User, error) {
	user, err := r.searchBy(User{AuthID: id})
	return user, err
}

// GetByEmail searches user by provided email
func (r repository) GetByEmail(email string) (User, error) {
	user, err := r.searchBy(User{Email: email})
	return user, err
}

// GetByPhone searches user by provided phone
func (r repository) GetByPhone(phone string) (User, error) {
	user, err := r.searchBy(User{Phone: phone})
	return user, err
}

// Update details of a user
func (r repository) Update(user User) error {
	var u User
	result := r.db.Model(&u).Omit("id").Updates(user)
	if err := result.Error; err != nil {
		return errors.Error{Err: result.Error, Code: errors.EINTERNAL}
	}
	return nil
}
