package repository

import (
	"BakeMeets/api/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func (repository *CommentRepository) GetAllComments() ([]models.Comment, error) {
	var Comments []models.Comment
	result := repository.DB.Find(&Comments)
	return Comments, result.Error
}

func (repository *CommentRepository) GetCommentByBreadID(id uint) (models.Comment, error) {
	var comment models.Comment
	result := repository.DB.First(&comment, id)
	return comment, result.Error
}

func (repository *CommentRepository) CreateComment(comment models.Comment) (models.Comment, error) {
	result := repository.DB.Create(&comment)
	return comment, result.Error
}

func (repository *CommentRepository) UpdateComment(comment models.Comment) (models.Comment, error) {
	result := repository.DB.Save(&comment)
	return comment, result.Error
}

func (repository *CommentRepository) DeleteComment(comment models.Comment) error {
	result := repository.DB.Delete(&comment)
	return result.Error
}
