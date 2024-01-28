package repositories

import "github.com/google/uuid"

type Users struct {
	UUID            uuid.UUID `gorm:"column:uuid"`
	Username        string    `gorm:"column:username"`
	Password        string    `gorm:"column:password"`
	Email           string    `gorm:"column:email;unique"`
	Role            string    `gorm:"column:role"`
	PermissionLevel int       `gorm:"column:permission_level"`
	CreatedAt       string    `gorm:"column:created_at"`
	UpdatedAt       string    `gorm:"column:updated_at"`
	Avatar          string    `gorm:"column:avatar"`
}

type UserRepository interface {
	Register(Users) (*Users, error)
	Login(Users) (*Users, error)
	GetAllUsers() ([]Users, error)
	GetUser(uuid.UUID) (*Users, error)
}
