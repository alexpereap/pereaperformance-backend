package controllers

import (
	"alexpereap/pereaperformance-backend.git/entity"
	"alexpereap/pereaperformance-backend.git/service"
)

type SlideController interface {
	FindAll() []entity.Slide
}

type slideController struct {
	service service.SlideService
}

func NewSlideController(service service.SlideService) SlideController {
	return &slideController{
		service: service,
	}
}

func (s *slideController) FindAll() []entity.Slide {
	return s.service.FindAll()
}
