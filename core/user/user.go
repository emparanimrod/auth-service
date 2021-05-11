package user

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	AuthID    uuid.UUID `gorm:"primaryKey"`
	Phone     string
	Firstname string
	Lastname  string
	Email     string
	PIN       string
	UserType  string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
