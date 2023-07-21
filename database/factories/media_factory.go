package factories

import "github.com/brianvoe/gofakeit/v6"

type MediaFactory struct {
}

// Definition Define the model's default state.
func (f *MediaFactory) Definition() map[string]any {
	faker := gofakeit.New(0)
	return map[string]interface{}{
		"URL":       faker.ImageURL(300, 400),
		"type":      faker.RandomString([]string{"jpeg", "png", "gif"}),
		"ownerID":   0,                                                          // Set the appropriate owner ID here
		"ownerType": faker.RandomString([]string{"posts", "comments", "users"}), // Set the appropriate owner type here
	}
}
