package repository

import (
	"template/model"

	"gorm.io/gorm"
)

type CardRepository interface {
	Create(card *model.Card) error
	GetByID(id string) (*model.Card, error)
	Getall(page, limit int) ([]model.Card, int64, error)
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db}
}

func (r *cardRepository) Create(card *model.Card) error {
	return r.db.Save(card).Error
}

func (r *cardRepository) GetByID(id string) (*model.Card, error) {
	var card model.Card
	err := r.db.Where("id = ?", id).First(&card).Error
	return &card, err
}

func (r *cardRepository) Getall(page, limit int) ([]model.Card, int64, error) {
	var cards []model.Card
	var total int64
	offset := (page - 1) * limit

	err := r.db.Model(&model.Card{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(limit).Find(&cards).Error
	return cards, total, err
}
