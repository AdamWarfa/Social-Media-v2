package repositories

type MockPostRepository struct {
	PostRepository
}

func NewMockPostRepository() PostRepository {
	return &MockPostRepository{}
}
