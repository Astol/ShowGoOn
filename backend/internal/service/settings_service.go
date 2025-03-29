package service

import (
	"ShowGoOn/backend/internal/repository"
	"time"
)

type Settings struct {
	ID              int
	SlideDuration   int
	SlideTransition int
	UpdateDatetime  time.Time
}

type SettingsService struct {
	repo repository.SettingsRepository
}

func NewSettingsService(repo repository.SettingsRepository) *SettingsService {
	return &SettingsService{repo: repo}
}

func (s *SettingsService) GetSettingsByID(id int) (*Settings, error) {
	repositorySettings, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toServiceSettings(repositorySettings), nil
}

func (s *SettingsService) CreateSettings(settings *Settings) error {
	repositorySettings := toRepositorySettings(settings)
	return s.repo.Create(repositorySettings)
}

func (s *SettingsService) UpdateSettings(settings *Settings) error {
	repositorySettings := toRepositorySettings(settings)
	return s.repo.Update(repositorySettings)
}

func (s *SettingsService) DeleteSettings(id int) error {
	return s.repo.Delete(id)
}

func toRepositorySettings(settings *Settings) *repository.Settings {
	return &repository.Settings{
		Id:              settings.ID,
		SlideDuration:   settings.SlideDuration,
		Slidetransition: settings.SlideTransition,
		UpdateDatetime:  settings.UpdateDatetime,
	}
}

func toServiceSettings(settings *repository.Settings) *Settings {
	return &Settings{
		ID:              settings.Id,
		SlideDuration:   settings.SlideDuration,
		SlideTransition: settings.Slidetransition,
		UpdateDatetime:  settings.UpdateDatetime,
	}
}
