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

func (db *userRepository) Save(user *entity.User) {
	db.connection.Create(&user)
}

func (db *userRepository) Delete(user entity.User) error {
	var existantUser entity.User
	db.connection.First(&existantUser, "id = ?", user.ID)

	if existantUser.ID == 0 {
		return errors.New("user not found")
	}

	db.connection.Delete(&user)

	return nil
}

func (db *userRepository) FindAll() []entity.User {
	var users []entity.User
	db.connection.Find(&users)
	return users
}

func (db *userRepository) FindOne(constraints map[string]interface{}) entity.User {
	var existantUser entity.User
	db.connection.First(&existantUser, constraints)

	return existantUser
}

func (db *userRepository) CloseDb() {

}
