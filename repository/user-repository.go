package repository

import (
	"alexpereap/pereaperformance-backend.git/entity"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.User)
	Delete(user entity.User) error
	FindAll() []entity.User
	FindOne(constanits map[string]interface{}) entity.User
	CloseDb()
}

type database struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	dsn := "host=localhost user=alex password=alex dbname=pereaperformance port=5432 sslmode=disable TimeZone=America/Bogota"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.User{})

	return &database{
		connection: db,
	}
}

func (db *database) Save(user *entity.User) {
	db.connection.Create(&user)
}

func (db *database) Delete(user entity.User) error {
	var existantUser entity.User
	db.connection.First(&existantUser, "id = ?", user.ID)

	if existantUser.ID == 0 {
		return errors.New("user not found")
	}

	db.connection.Delete(&user)

	return nil
}

func (db *database) FindAll() []entity.User {
	var users []entity.User
	db.connection.Find(&users)
	return users
}

func (db *database) FindOne(constraints map[string]interface{}) entity.User {
	var existantUser entity.User
	db.connection.First(&existantUser, constraints)

	return existantUser
}

func (db *database) CloseDb() {

}
