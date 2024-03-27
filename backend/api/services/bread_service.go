package services

import (
	"BakeMeets/api/models"
	"BakeMeets/repository"
)

type BreadService struct {
	Repo *repository.BreadRepository
}

func NewBreadService(repo *repository.BreadRepository) *BreadService {
	return &BreadService{
		Repo: repo,
	}
}

func (bs *BreadService) GetAllBreads() ([]models.Bread, error) {
	return bs.Repo.GetAllBreads()
}

func (bs *BreadService) GetBreadByID(id uint) (models.Bread, error) {
	return bs.Repo.GetBreadByID(id)
}

func (bs *BreadService) CreateBread(bread *models.Bread) error {
	return bs.Repo.CreateBread(*bread)
}

func (bs *BreadService) UpdateBread(bread *models.Bread) error {
	return bs.Repo.UpdateBread(*bread)
}

func (bs *BreadService) DeleteBread(id uint) error {
	return bs.Repo.DeleteBread(id)
}
