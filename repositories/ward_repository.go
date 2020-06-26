package repositories

import (
	"math"

	"github.com/jinzhu/gorm"
	"github.com/malbbako/ng_state_ward/dtos"
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
func (r *WardRepository) FindAllWard(pagination *dtos.Pagination) (RepositoryResult, int) {

	var wards models.Wards
	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	offset := pagination.Page * pagination.Limit

	//get data with limit,offset & order
	errFind := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Preload("LocalGovernment").Find(&wards).Error
	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}
	pagination.Rows = wards

	//count all rows
	errCount := r.db.Model(&models.Ward{}).Count(&totalRows).Error
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

func (r *WardRepository) FindWardByLocalgovernment(localgovernmentID int, pagination *dtos.Pagination) (RepositoryResult, int) {

	var wards models.Wards
	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	offset := pagination.Page * pagination.Limit

	//get data with limit,offset & order
	errFind := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where(&models.Ward{LocalGovernmentID: localgovernmentID}).Preload("Localgovernment").Find(&wards).Error
	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}
	pagination.Rows = wards

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
