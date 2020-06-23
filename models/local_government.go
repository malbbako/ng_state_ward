package models

import "time"

type LocalGovernment struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Abbr      string `gorm:"type:varchar(3);NOT NULL" json:"abbr" binding:"required"`
	Name      string `gorm:"type:varchar(150);NOT NULL;UNIQUE" json:"name" binding:"required"`
	State     State  `json:"state" gorm:"foreignkey:StateID"`
	StateID   int    `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type LocalGovernments []LocalGovernment
