package service

import (
	"ShowGoOn/backend/internal/repository"
)

type SlideService struct {
	repo repository.SlideRepository
}

func NewSlideService(repo repository.SlideRepository) *SlideService {
	return &SlideService{repo: repo}
}

func (s *SlideService) GetSlideByID(id int) (*repository.Slide, error) {
	return s.repo.GetByID(id)
}

func (s *SlideService) CreateSlide(slide *repository.Slide) error {
	return s.repo.Create(slide)
}

func (s *SlideService) UpdateSlide(slide *repository.Slide) error {
	return s.repo.Update(slide)
}

func (s *SlideService) DeleteSlide(id int) error {
	return s.repo.Delete(id)
}
