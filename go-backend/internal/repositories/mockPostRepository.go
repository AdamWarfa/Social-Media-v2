package repositories

type MockPostRepository struct {
	PostRepositoryI
}

func NewMockPostRepository() PostRepositoryI {
	return &MockPostRepository{}
}
