package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/malbbako/ng_state_ward/models"
)

type WardRepository struct {
	db *gorm.DB
}

func NewWardRepository(db *gorm.DB) *WardRepository {
	return &WardRepository{db: db}
}

func (r *WardRepository) CreateWard(ward *models.Ward) RepositoryResult {
	err := r.db.Save(ward).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: ward}
}
