package models

import (
	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"

	"goravel/database/factories"
)

type Tag struct {
	orm.Model
	Title        string  `json:"title" gorm:"unique" faker:"unique"`
	CoverImage   string  `json:"cover_image"`
	BackendImage string  `json:"backend_image"`
	Description  string  `json:"description"`
	Color        string  `json:"color"`
	Slug         string  `json:"slug"`
	UserID       uint    `json:"user_id"`
	User         User    `json:"user" gorm:"foreignKey:UserID"`
	Posts        []*Post `json:"posts" gorm:"many2many:post_hashtags;"`
	orm.SoftDeletes
}

func (t *Tag) Factory() factory.Factory {
	return &factories.TagFactory{}
}
