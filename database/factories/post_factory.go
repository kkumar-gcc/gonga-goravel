package factories

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/goravel/framework/support/carbon"
)

type PostFactory struct {
}

// Definition Define the model's default state.
func (f *PostFactory) Definition() map[string]any {
	faker := gofakeit.New(0)
	return map[string]interface{}{
		"title":           faker.Sentence(30),
		"body":            faker.Paragraph(4, 20, 20, "."),
		"userID":          0, // Set the appropriate user ID here
		"likeCount":       0,
		"commentCount":    0,
		"viewCount":       0,
		"shareCount":      0,
		"isPromoted":      false,
		"promotionExpiry": carbon.Now(),
		"isFeatured":      false,
		"featuredExpiry":  carbon.Now(),
		"visibility":      "public",
	}
}
