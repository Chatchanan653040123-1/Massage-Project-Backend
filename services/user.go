package services

import (
	"time"

	"github.com/google/uuid"
)

type LoginBody struct {
	UUID       uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	Password   string    `json:"password"`
}

type RegisterBody struct {
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
	Avatar   string    `json:"avatar"`
}
type UUID struct {
	UUID uuid.UUID `json:"id"`
}
type GetUsersResponse struct {
	UUID     uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
}
type EntityPermission struct {
	Role            string `json:"role"`
	PermissionLevel int    `json:"permission_level"`
}
type UpdateUserRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	Role            string `json:"role"`
	PermissionLevel int    `json:"permission_level"`
}
type UserService interface {
	Register(RegisterBody) (*UUID, error)
	Login(LoginBody) (*LoginBody, error)
	GetAllUsers() ([]GetUsersResponse, error)
	GetUser(uuid uuid.UUID) (*GetUsersResponse, error)
	UpdateMyAccount(uuid.UUID, RegisterBody) (*UUID, error)
	UpdateAccount(uuid.UUID, UpdateUserRequest) (*UUID, error)
	DeleteAccount(uuid.UUID) (*UUID, error)
	GetEntityPermission(uuid.UUID) (*EntityPermission, error)
}
