package entity

import (
	"gorm.io/gorm"
)

type TableStatus string

const (
	TableObserved  TableStatus = "observed"
	TableAvailable TableStatus = "available"
)

type Table struct {
	gorm.Model
	Number int         `gorm:"type:int;not null;unique"`
	Status TableStatus `gorm:"type:table_status;default:'available'"`
	Orders []Order     `gorm:"foreignKey:TableID"`
}
