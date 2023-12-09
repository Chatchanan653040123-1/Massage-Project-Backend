package services

import (
	errs "massage/errors"
	"massage/logs"
	"massage/repositories"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) Register(req RegisterBody) (*UUID, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if req.Username == "" || req.Password == "" || req.Email == "" {
		logs.Error("Username, Password, and Email cannot be empty")
		return nil, errs.NewValidationError("(Username or Email) and password cannot be empty")
	}

	if err != nil {
		return nil, errs.NewUnexpectedError("Failed to hash password")
	}

	user := repositories.Users{
		UUID:      uuid.New(),
		Username:  req.Username,
		Password:  string(password),
		Email:     req.Email,
		IsAdmin:   false,
		CreatedAt: time.Now().Format("2006-1-2 15:04:05"),
	}

	newUser, err := s.userRepo.Register(user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("Failed to register user")
	}

	userResponse := UUID{
		UUID: newUser.UUID,
	}
	return &userResponse, nil
}

func (s userService) Login(req LoginBody) (*LoginBody, error) {
	user := repositories.Users{
		Username: req.Identifier,
		Password: req.Password,
	}
	if strings.Contains(req.Identifier, "@") {
		user = repositories.Users{
			Email:    req.Identifier,
			Password: req.Password,
		}
	}

	loginUser, err := s.userRepo.Login(user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("Invalid username or password")
	}

	allUserResponse := LoginBody{
		UUID:     loginUser.UUID,
		Password: loginUser.Password,
	}
	return &allUserResponse, nil
}
func (s userService) GetAllUsers() ([]GetAllUsersResponse, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("Failed to get all users")
	}

	var allUserResponse []GetAllUsersResponse
	for _, user := range users {
		allUserResponse = append(allUserResponse, GetAllUsersResponse{
			UUID:     user.UUID,
			Username: user.Username,
			Avatar:   user.Avatar,
		})
	}
	return allUserResponse, nil
}
