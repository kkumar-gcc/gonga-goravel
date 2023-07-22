package seeders

type TagSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *TagSeeder) Signature() string {
	return "TagSeeder"
}

// Run executes the seeder logic.
func (s *TagSeeder) Run() error {
	return nil
}
