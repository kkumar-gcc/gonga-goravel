package seeders

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	"math/rand"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	var users []models.User
	err := facades.Orm().Factory().Times(500).Create(&users)
	if err != nil {
		return err
	} // 500 users

	var tags []models.Tag
	if err := facades.Orm().Factory().Times(20).Make(&tags); err != nil {
		return err
	} // 20 tags
	for i := range tags {
		tags[i].UserID = users[rand.Intn(len(users))].ID
	}
	// Save changes to the tags (updating the user_id)
	if err := facades.Orm().Query().Create(&tags); err != nil {
		return err
	}

	for _, user := range users {
		var posts []models.Post
		err := facades.Orm().Factory().Times(rand.Intn(4)).Make(&posts)
		if err != nil {
			return err
		} // max 4 posts

		for _, post := range posts {
			post.UserID = user.ID
			if err := facades.Orm().Query().Create(&post); err != nil {
				return err
			}

			var comments []models.Comment
			err := facades.Orm().Factory().Times(rand.Intn(5)).Make(&comments)
			if err != nil {
				return err
			} // max 5 comments
			for i := range comments {
				comments[i].UserID = user.ID
				comments[i].PostID = post.ID
			}
			if err := facades.Orm().Query().Create(&comments); err != nil {
				return err
			}
		}
	}
	return nil
}
