package service

import (
	"alexpereap/pereaperformance-backend.git/entity"
	"alexpereap/pereaperformance-backend.git/repository"
)

type SlideService interface {
	FindAll() []entity.Slide
	Save(slide *entity.Slide) *entity.Slide
}

type slideService struct {
	SlideRepository repository.SlideRepository
}

func NewSlideService(repo repository.SlideRepository) SlideService {
	return &slideService{
		SlideRepository: repo,
	}
}

func (service *slideService) FindAll() []entity.Slide {
	return service.SlideRepository.FindAll()
}

func (service *slideService) Save(slide *entity.Slide) *entity.Slide {
	service.SlideRepository.Save(slide)
	return slide
}
