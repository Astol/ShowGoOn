package service

import (
	"ShowGoOn/backend/internal/repository"
)

type Slide struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	MediaPath string `json:"media_path"`
	MediaType int    `json:"media_type"`
}

type SlideService struct {
	repo repository.SlideRepository
}

func NewSlideService(repo repository.SlideRepository) *SlideService {
	return &SlideService{repo: repo}
}

func (s *SlideService) GetSlideByID(id int) (*Slide, error) {
	repoSlide, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toServiceSlide(repoSlide), nil
}

func (s *SlideService) CreateSlide(slide *Slide) error {
	repoSlide := toRepositorySlide(slide)
	return s.repo.Create(repoSlide)
}

func (s *SlideService) UpdateSlide(slide *Slide) error {
	repoSlide := toRepositorySlide(slide)
	return s.repo.Update(repoSlide)
}

func (s *SlideService) DeleteSlide(id int) error {
	return s.repo.Delete(id)
}

func (s *SlideService) GetAllSlides() ([]Slide, error) {
	repoSlides, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Convert repository slides to service slides
	serviceSlides := make([]Slide, len(repoSlides))
	for i, repoSlide := range repoSlides {
		serviceSlides[i] = *toServiceSlide(&repoSlide)
	}
	return serviceSlides, nil
}

func toRepositorySlide(slide *Slide) *repository.Slide {
	return &repository.Slide{
		Id:        slide.ID,
		Title:     slide.Title,
		Text:      slide.Text,
		MediaPath: slide.MediaPath,
		MediaType: slide.MediaType,
	}
}

func toServiceSlide(repoSlide *repository.Slide) *Slide {
	return &Slide{
		ID:        repoSlide.Id,
		Title:     repoSlide.Title,
		Text:      repoSlide.Text,
		MediaPath: repoSlide.MediaPath,
		MediaType: repoSlide.MediaType,
	}
}
