package seeders

import (
	"github.com/goravel/framework/contracts/database/seeder"
	"github.com/goravel/framework/facades"
)

type DatabaseSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *DatabaseSeeder) Signature() string {
	return "DatabaseSeeder"
}

// Run executes the seeder logic.
func (s *DatabaseSeeder) Run() error {
	err := facades.Seeder().CallOnce([]seeder.Seeder{
		&TagSeeder{},
		&UserSeeder{},
		&PostSeeder{},
		&CommentSeeder{},
	})
	if err != nil {
		return err
	}
	return nil
}
