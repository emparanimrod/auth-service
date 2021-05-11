package core

import (
	"auth/core/user"
	"auth/storage"
)

type Core struct {
	User      user.Service
}

func New(database *storage.Database) *Core {
	userRepo := user.NewRepository(database)
	// shopRepo := shop.NewRepository(database)
	// shopkeeperRepo := shopkeeper.NewRepository(database)

	return &Core{
		User:      user.NewService(userRepo),
	}
}
