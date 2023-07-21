package models

import (
	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"

	"goravel/database/factories"
)

type Visibility string

const (
	VisibilityPublic  Visibility = "public"
	VisibilityPrivate Visibility = "private"
	VisibilityFriends Visibility = "friends"
	// Add more visibility options as needed
)

type Post struct {
	orm.Model
	UserID          uint            `json:"user_id"`
	User            *User           `json:"user,omitempty"`
	Title           string          `json:"title"`
	Body            string          `json:"body"`
	Likes           []Like          `json:"likes" gorm:"polymorphic:Likeable;"`
	LikeCount       uint            `json:"like_count"`
	Comments        []*Comment      `json:"comments" gorm:"foreignKey:PostID"`
	CommentCount    uint            `json:"comment_count"`
	ViewCount       uint            `json:"view_count"`
	ShareCount      uint            `json:"share_count"`
	Medias          []*Media        `json:"medias" gorm:"polymorphic:Owner;"`
	Hashtags        []*Tag          `json:"hashtags" gorm:"many2many:post_hashtags;"`
	Mentions        []*Mention      `json:"mentions" gorm:"polymorphic:Owner;"`
	IsPromoted      bool            `json:"is_promoted"`
	PromotionExpiry carbon.DateTime `json:"promotion_expiry"`
	IsFeatured      bool            `json:"is_featured"`
	FeaturedExpiry  carbon.DateTime `json:"featured_expiry"`
	Visibility      Visibility      `json:"visibility"`
	orm.SoftDeletes
}

func (p *Post) Factory() factory.Factory {
	return &factories.PostFactory{}
}
