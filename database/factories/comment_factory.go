package factories

import "github.com/brianvoe/gofakeit/v6"

type CommentFactory struct {
}

// Definition Define the model's default state.
func (f *CommentFactory) Definition() map[string]any {
	faker := gofakeit.New(0)
	return map[string]interface{}{
		"body":     faker.Paragraph(2, 20, 20, "."),
		"userID":   0,
		"postID":   0,
		"parentID": nil,
	}
}
