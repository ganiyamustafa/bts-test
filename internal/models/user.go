package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Username string    `json:"username" faker:"name"`
	Email    string    `json:"email" gorm:"unique" faker:"email"`
	Password string    `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

// hash user password function
func (u *User) HashPassword() error {
	// hash user password with bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// set user password to hash password
	u.Password = string(hashPassword)
	return nil
}
