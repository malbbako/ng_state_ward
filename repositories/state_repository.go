package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/malbbako/ng_state_ward/models"
)

type StateRepository struct {
	db *gorm.DB
}

func NewStateRepository(db *gorm.DB) *StateRepository {
	return &StateRepository{db: db}
}

func (r *StateRepository) Save(state *models.State) RepositoryResult {
	err := r.db.Save(state).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: state}
}

func (r *StateRepository) FindAll() RepositoryResult {
	var states models.States
	err := r.db.Find(&states).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: states}
}
