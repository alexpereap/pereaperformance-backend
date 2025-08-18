package repository

import (
	"alexpereap/pereaperformance-backend.git/entity"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.User)
	Delete(user entity.User) error
	FindAll() []entity.User
	FindOne(constanits map[string]interface{}) entity.User
	CloseDb()
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {

	dbConn.AutoMigrate(&entity.User{})

	return &userRepository{
		connection: dbConn,
	}
}

func (u *userRepository) Save(user *entity.User) {
	u.connection.Create(&user)
}

func (u *userRepository) Delete(user entity.User) error {
	var existantUser entity.User
	u.connection.First(&existantUser, "id = ?", user.ID)

	if existantUser.ID == 0 {
		return errors.New("user not found")
	}

	u.connection.Delete(&user)

	return nil
}

func (u *userRepository) FindAll() []entity.User {
	var users []entity.User
	u.connection.Find(&users)
	return users
}

func (u *userRepository) FindOne(constraints map[string]interface{}) entity.User {
	var existantUser entity.User
	u.connection.First(&existantUser, constraints)

	return existantUser
}

func (u *userRepository) CloseDb() {

}
