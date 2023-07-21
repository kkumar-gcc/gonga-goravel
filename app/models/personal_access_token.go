package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
)

type PersonalAccessToken struct {
	orm.Model
	Name       string `gorm:"not null"`
	Token      string `gorm:"unique;not null"`
	LastUsedAt carbon.DateTime
	ExpiresAt  carbon.DateTime
}
