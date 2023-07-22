package seeders

type PostSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *PostSeeder) Signature() string {
	return "PostSeeder"
}

// Run executes the seeder logic.
func (s *PostSeeder) Run() error {
	return nil
}
