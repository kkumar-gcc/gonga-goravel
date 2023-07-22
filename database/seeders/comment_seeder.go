package seeders

type CommentSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *CommentSeeder) Signature() string {
	return "CommentSeeder"
}

// Run executes the seeder logic.
func (s *CommentSeeder) Run() error {
	return nil
}
