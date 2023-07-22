package factories

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/goravel/framework/support/carbon"
)

type UserFactory struct {
}

// Definition Define the model's default state.
func (f *UserFactory) Definition() map[string]any {
	faker := gofakeit.New(0)
	return map[string]interface{}{
		"email":         faker.Email(),
		"password":      faker.Password(true, true, true, true, false, 12),
		"username":      faker.Username(),
		"firstName":     faker.FirstName(),
		"lastName":      faker.LastName(),
		"avatarUrl":     faker.Person().Image,
		"bio":           faker.Paragraph(2, 20, 20, "."),
		"gender":        faker.Gender(),
		"mobileNo":      faker.Phone(),
		"mobileNoCode":  "+91",
		"birthday":      carbon.Now(),
		"country":       faker.Country(),
		"city":          faker.City(),
		"websiteUrl":    faker.URL(),
		"occupation":    faker.JobTitle(),
		"education":     faker.BS(),
		"emailVerified": true,
		"createdAt":     carbon.Now(),
		"updatedAt":     carbon.Now(),
	}
}
