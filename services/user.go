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
type GetAllUsersResponse struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
}

type UserService interface {
	Register(RegisterBody) (*UUID, error)
	Login(LoginBody) (*LoginBody, error)
	GetAllUsers() ([]GetAllUsersResponse, error)
}
