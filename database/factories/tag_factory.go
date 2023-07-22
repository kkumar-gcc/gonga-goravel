package factories

import "github.com/brianvoe/gofakeit/v6"

type TagFactory struct {
}

// Definition Define the model's default state.
func (f *TagFactory) Definition() map[string]any {
	faker := gofakeit.New(0)
	return map[string]interface{}{
		"title":        faker.Username(),
		"coverImage":   faker.ImageURL(300, 400),
		"backendImage": faker.ImageURL(800, 600),
		"description":  faker.Paragraph(1, 5, 15, "."),
		"color":        faker.HexColor(),
		"slug":         faker.Sentence(10),
		"userID":       0,
	}
}
