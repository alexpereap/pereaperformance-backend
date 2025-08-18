package repository

import (
	"alexpereap/pereaperformance-backend.git/entity"

	"gorm.io/gorm"
)

type SlideRepository interface {
	FindAll() []entity.Slide
}

type slideRepository struct {
	connection *gorm.DB
}

func NewSlideRepository(dbConn *gorm.DB) SlideRepository {
	dbConn.AutoMigrate(&entity.Slide{})

	return &slideRepository{
		connection: dbConn,
	}
}

func (s *slideRepository) FindAll() []entity.Slide {
	var slides []entity.Slide

	s.connection.Find(&slides)

	return slides
}
