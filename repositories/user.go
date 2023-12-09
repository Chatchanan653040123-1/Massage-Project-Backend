package repositories

import "github.com/google/uuid"

type Users struct {
	UUID      uuid.UUID `gorm:"column:id"`
	Username  string    `gorm:"username"`
	Password  string    `gorm:"password"`
	Email     string    `gorm:"unique"`
	IsAdmin   bool      `gorm:"is_admin"`
	CreatedAt string    `gorm:"created_at"`
	UpdateAt  string    `gorm:"update_at"`
	Avatar    string    `gorm:"avatar"`
}

type UserRepository interface {
	Register(Users) (*Users, error)
	Login(Users) (*Users, error)
	GetAllUsers() ([]Users, error)
}
