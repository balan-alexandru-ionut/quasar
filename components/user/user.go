package user

import (
	"context"
	"quasar/database"
	"time"

	"gorm.io/gorm"
)

var db = database.DB

type User struct {
	ID           uint   `gorm:"autoIncrement"`
	FirstName    string `gorm:"not null,size:128"`
	LastName     string `gorm:"not null,size:128"`
	Email        string `gorm:"not null,unique"`
	PasswordHash []byte
	PasswordSalt []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) Create() error {
	ctx := context.Background()

	return gorm.G[User](db).Create(ctx, u)
}
