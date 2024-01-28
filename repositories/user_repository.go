package repositories

import (
	"massage/logs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRespositoryDB(db *gorm.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}
func (r userRepositoryDB) Register(user Users) (*Users, error) {
	err := r.db.Create(&user)
	if err.Error != nil {
		logs.Error(err.Error)
		return nil, err.Error
	}

	return &user, nil
}

func (r userRepositoryDB) Login(user Users) (*Users, error) {
	if user.Username != "" {
		err := r.db.First(&user, "username = ?", user.Username)
		if err.Error != nil {
			logs.Error(err.Error)
			return nil, err.Error
		}
	}
	if user.Email != "" {
		err := r.db.First(&user, "email = ?", user.Email)
		if err.Error != nil {
			logs.Error(err.Error)
			return nil, err.Error
		}
	}

	return &user, nil
}

// to get all users
func (r userRepositoryDB) GetAllUsers() ([]Users, error) {
	var users []Users
	err := r.db.Find(&users)
	if err.Error != nil {
		logs.Error(err.Error)
		return nil, err.Error
	}

	return users, nil
}
func (r userRepositoryDB) GetUser(uuid uuid.UUID) (*Users, error) {
	var user Users
	err := r.db.First(&user, "uuid = ?", uuid)
	if err.Error != nil {
		logs.Error(err.Error)
		return nil, err.Error
	}

	return &user, nil
}
func (r userRepositoryDB) UpdateAccount(uuid uuid.UUID, user Users) (*Users, error) {
	err := r.db.Model(&user).Where("uuid = ?", uuid).Updates(Users{
		UUID:            uuid,
		Username:        user.Username,
		Password:        user.Password,
		Role:            user.Role,
		Email:           user.Email,
		PermissionLevel: user.PermissionLevel,
		Avatar:          user.Avatar,
		UpdatedAt:       user.UpdatedAt,
	})
	if err.Error != nil {
		logs.Error(err.Error)
		return nil, err.Error
	}

	return &user, nil
}
func (r userRepositoryDB) DeleteAccount(uuid uuid.UUID) (*Users, error) {
	var user Users
	err := r.db.Where("uuid = ?", uuid).Delete(&user)
	if err.Error != nil {
		logs.Error(err.Error)
		return nil, err.Error
	}

	return &user, nil
}
