package service

import (
	"alexpereap/pereaperformance-backend.git/entity"
	"alexpereap/pereaperformance-backend.git/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Save(user *entity.User) *entity.User
	Delete(user entity.User) error
	FindAll() []entity.User
	FindOne(constraints map[string]interface{}) entity.User
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		UserRepository: repo,
	}
}

func (service *userService) Save(user *entity.User) *entity.User {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Can't hash password %v", err)
	}

	user.Password = string(hash)

	service.UserRepository.Save(user)
	return user
}

func (service *userService) Delete(user entity.User) error {
	var err error = nil
	err = service.UserRepository.Delete(user)

	return err
}

func (service *userService) FindAll() []entity.User {
	return service.UserRepository.FindAll()
}

func (service *userService) FindOne(constraints map[string]interface{}) entity.User {
	return service.UserRepository.FindOne(constraints)
}
