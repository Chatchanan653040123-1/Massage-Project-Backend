package repositories

import (
	"github.com/google/uuid"
)

type Users struct {
	UUID            uuid.UUID `gorm:"primaryKey"`
	Username        string    `gorm:"column:username"`
	Password        string    `gorm:"column:password"`
	Email           string    `gorm:"column:email;unique"`
	Role            string    `gorm:"column:role"`
	PermissionLevel int       `gorm:"column:permission_level"`
	Avatar          string    `gorm:"column:avatar"`
	CreatedAt       string    `gorm:"column:created_at"`
	UpdatedAt       string    `gorm:"column:updated_at"`
}
type Groups struct {
	UUID   uuid.UUID `gorm:"primaryKey"`
	Name   string    `gorm:"column:name;unique"`
	UserID uuid.UUID `gorm:"type:uuid;column:user_id;not null;"`
	User   Users     `gorm:"foreignKey:UserID"`
}

type UserRepository interface {
	Register(Users) (*Users, error)
	Login(Users) (*Users, error)
	GetAllUsers() ([]Users, error)
	GetUser(uuid.UUID) (*Users, error)
	UpdateAccount(uuid.UUID, Users) (*Users, error)
	DeleteAccount(uuid.UUID) (*Users, error)
	CreateGroup(Groups) (*Groups, error)
}
