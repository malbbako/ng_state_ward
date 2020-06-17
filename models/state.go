package models

import "time"

type State struct {
	ID        string `gorm:"primary_key"`
	Abbr      string `gorm:"type:varchar(3);NOT NULL"`
	Name      string `gorm:"type:varchar(150);NOT NULL;UNIQUE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type States []State
