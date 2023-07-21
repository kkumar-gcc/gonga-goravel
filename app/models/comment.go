package models

import (
	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"

	"goravel/database/factories"
)

type Comment struct {
	orm.Model
	UserID   uint       `json:"user_id"`
	User     *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	PostID   uint       `json:"post_id"`
	Post     *Post      `json:"post,omitempty" gorm:"foreignKey:PostID;"`
	Body     string     `json:"body"`
	Likes    []Like     `json:"likes" gorm:"polymorphic:Likeable;"`
	ParentID *uint      `json:"parent_id"`
	Parent   *Comment   `json:"parent,omitempty"`
	Children []*Comment `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Mentions []*Mention `json:"mentions" gorm:"polymorphic:Owner;"`
	orm.SoftDeletes
}

func (c *Comment) Factory() factory.Factory {
	return &factories.CommentFactory{}
}
