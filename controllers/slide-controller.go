package controllers

import (
	"alexpereap/pereaperformance-backend.git/entity"
	"alexpereap/pereaperformance-backend.git/service"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type SlideController interface {
	FindAll() []entity.Slide
	Dashboard(ctx *gin.Context)
	CreateForm(ctx *gin.Context)
	Save(ctx *gin.Context) (*entity.Slide, error)
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

func (s *slideController) Save(ctx *gin.Context) (*entity.Slide, error) {
	/*var slide *entity.Slide
	err := ctx.ShouldBind(&slide)*/

	slide := new(entity.Slide)
	err := ctx.ShouldBind(slide)

	if err != nil {
		return slide, err
	}

	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 10<<20) // 10 MB
	ctx.Request.ParseMultipartForm(12 << 20)                                     // 12 MB

	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required", "detail": err.Error()})
		return slide, err
	}

	destDir := "uploads/slides"
	if err := os.MkdirAll(destDir, 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create folder", "detail": err.Error()})
		return slide, err
	}

	// generates unique name
	ext := filepath.Ext(file.Filename)
	base := filepath.Base(file.Filename[:len(file.Filename)-len(ext)]) // no ext
	if base == "." || base == "/" || base == "" {
		base = "file"
	}
	filename := fmt.Sprintf("%s_%d%s", base, time.Now().UnixNano(), ext)

	// final path
	dst := filepath.Join(destDir, filename)

	// Save file on disc
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file", "detail": err.Error()})
		return slide, err
	}

	slide.Image = filename

	slide = s.service.Save(slide)
	return slide, nil
}
