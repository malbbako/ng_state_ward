package models

import "time"

type State struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	Abbr      string `gorm:"type:varchar(3);NOT NULL" json:"abbr" binding:"required"`
	Name      string `gorm:"type:varchar(150);NOT NULL;UNIQUE" json:"name" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type States []State
