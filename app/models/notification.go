package models

import (
	"github.com/goravel/framework/database/orm"
)

type Notification struct {
	orm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	orm.SoftDeletes
}
