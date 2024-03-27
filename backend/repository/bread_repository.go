package repository

import (
	"BakeMeets/api/models"
	"gorm.io/gorm"
)

type BreadRepository struct {
	DB *gorm.DB
}

func (repository *BreadRepository) GetAllBreads() ([]models.Bread, error) {
	var breads []models.Bread
	result := repository.DB.Find(&breads)
	return breads, result.Error
}

func (repository *BreadRepository) GetBreadByID(id uint) (models.Bread, error) {
	var bread models.Bread
	result := repository.DB.First(&bread, id)
	return bread, result.Error
}

func (repository *BreadRepository) CreateBread(bread models.Bread) (models.Bread, error) {
	result := repository.DB.Create(&bread)
	return bread, result.Error
}

func (repository *BreadRepository) UpdateBread(bread models.Bread) (models.Bread, error) {
	result := repository.DB.Save(&bread)
	return bread, result.Error
}

func (repository *BreadRepository) DeleteBread(bread models.Bread) error {
	result := repository.DB.Delete(&bread)
	return result.Error
}
