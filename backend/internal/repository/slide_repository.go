package repository

import (
	"gorm.io/gorm"
)

type Slide struct {
	Id    int `gorm:"primary_key;AUTO_INCREMENT"`
	Title string
	Text  string
}

type SlideRepository interface {
	GetByID(id int) (*Slide, error)
	Create(slide *Slide) error
	Update(slide *Slide) error
	Delete(id int) error
}

type slideRepository struct {
	db *gorm.DB
}

func NewSlideRepository(db *gorm.DB) SlideRepository {
	return &slideRepository{db: db}
}

func (r *slideRepository) GetByID(id int) (*Slide, error) {
	var slide Slide
	if err := r.db.Where("id = ?", id).First(&slide).Error; err != nil {
		return nil, err
	}
	return &slide, nil
}

func (r *slideRepository) Create(slide *Slide) error {
	return r.db.Create(slide).Error
}

func (r *slideRepository) Update(slide *Slide) error {
	return r.db.Save(slide).Error
}

func (r *slideRepository) Delete(id int) error {
	return r.db.Where("id = ?", id).Delete(&Slide{}).Error
}
