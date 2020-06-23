package models

import "time"

type Ward struct {
	ID                int             `gorm:"primary_key" json:"id"`
	Abbr              string          `gorm:"type:varchar(3);NOT NULL" json:"abbr" binding:"required"`
	Name              string          `gorm:"type:varchar(150);NOT NULL;UNIQUE" json:"name" binding:"required"`
	LocalGovernment   LocalGovernment `json:"localgovernment" gorm:"foreignkey:LocalGovernmentID"`
	LocalGovernmentID int             `json:"-"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
type Wards []Ward
