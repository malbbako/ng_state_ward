package repositories

import (
	"math"

	"github.com/jinzhu/gorm"
	"github.com/malbbako/ng_state_ward/dtos"
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

// func (r *StateRepository) FindAll() RepositoryResult {
// 	var states models.States
// 	err := r.db.Find(&states).Error

// 	if err != nil {
// 		return RepositoryResult{Error: err}
// 	}

// 	return RepositoryResult{Result: &states}
// }
func (r *StateRepository) FindOneById(id int) RepositoryResult {
	var state models.State
	err := r.db.Where(&models.State{ID: id}).Take(&state).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: &state}

}

func (r *StateRepository) FindAll(pagination *dtos.Pagination) (RepositoryResult, int) {

	var states models.States
	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	offset := pagination.Page * pagination.Limit

	//get data with limit,offset & order
	errFind := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&states).Error
	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}
	pagination.Rows = states

	//count all rows
	errCount := r.db.Model(&models.State{}).Count(&totalRows).Error
	if errCount != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}
	pagination.TotalRows = totalRows

	//calculate total Pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1
	if pagination.Page == 0 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		fromRow = pagination.Page*pagination.Limit + 1
		toRow = (pagination.Page + 1) * pagination.Limit
	}

	if toRow > totalRows {
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.Torow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
