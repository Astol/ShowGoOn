package repository

import (
	"time"

	"gorm.io/gorm"
)

type Settings struct {
	Id              int `gorm:"primary_key;AUTO_INCREMENT"`
	SlideDuration   int
	Slidetransition int
	UpdateDatetime  time.Time
}

const (
	TransitionTypeSlide = 1
	TransitionTypeFade  = 2
)

type SettingsRepository interface {
	GetByID(id int) (*Settings, error)
	Create(settings *Settings) error
	Update(settings *Settings) error
	Delete(id int) error
}

type settingsRepository struct {
	db *gorm.DB
}

func NewSettingsRepository(db *gorm.DB) SettingsRepository {
	return &settingsRepository{db: db}
}

func (r *settingsRepository) GetByID(id int) (*Settings, error) {
	var settings Settings
	if err := r.db.Where("id = ?", id).First(&settings).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}

func (r *settingsRepository) Create(settings *Settings) error {
	return r.db.Create(settings).Error
}

func (r *settingsRepository) Update(settings *Settings) error {
	return r.db.Save(settings).Error
}

func (r *settingsRepository) Delete(id int) error {
	return r.db.Where("id = ?", id).Delete(&Settings{}).Error
}
