package repositories

import (
	"math"

	"github.com/jinzhu/gorm"
	"github.com/malbbako/ng_state_ward/dtos"
	"github.com/malbbako/ng_state_ward/models"
)

type LocalGovernmentRepository struct {
	db *gorm.DB
}

func NewLocalGovernmentRepository(db *gorm.DB) *LocalGovernmentRepository {
	return &LocalGovernmentRepository{db: db}
}

func (r *LocalGovernmentRepository) CreateLocalGovernment(localGovernment *models.LocalGovernment) RepositoryResult {
	err := r.db.Save(localGovernment).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: localGovernment}
}

func (r *LocalGovernmentRepository) FindLocalGovernmentById(id int) RepositoryResult {
	var localGovernment models.LocalGovernment
	err := r.db.Where(&models.LocalGovernment{ID: id}).Take(&localGovernment).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: &localGovernment}

}

func (r *LocalGovernmentRepository) FindLocalGovernmentByState(stateID int, pagination *dtos.Pagination) (RepositoryResult, int) {

	var localGovernments models.LocalGovernments
	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	offset := pagination.Page * pagination.Limit

	//get data with limit,offset & order
	errFind := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where(&models.LocalGovernment{StateID: stateID}).Take(&localGovernments).Error
	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}
	pagination.Rows = localGovernments

	//count all rows
	errCount := r.db.Model(&models.LocalGovernment{}).Count(&totalRows).Error
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

func (r *LocalGovernmentRepository) FindAllLocalGovernment(pagination *dtos.Pagination) (RepositoryResult, int) {

	var localGovernments models.LocalGovernments
	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	offset := pagination.Page * pagination.Limit

	//get data with limit,offset & order
	errFind := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&localGovernments).Error
	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}
	pagination.Rows = localGovernments

	//count all rows
	errCount := r.db.Model(&models.LocalGovernment{}).Count(&totalRows).Error
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
