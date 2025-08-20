package controllers

import (
	"alexpereap/pereaperformance-backend.git/entity"
	"alexpereap/pereaperformance-backend.git/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SlideController interface {
	FindAll() []entity.Slide
	Dashboard(ctx *gin.Context)
	CreateForm(ctx *gin.Context)
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

func (s *slideController) Dashboard(ctx *gin.Context) {
	slides := s.service.FindAll()
	ctx.HTML(http.StatusOK, "cms/slides/dashboard.html", gin.H{
		"slides": slides,
	})
}

func (s *slideController) CreateForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "cms/slides/createForm.html", gin.H{})
}
