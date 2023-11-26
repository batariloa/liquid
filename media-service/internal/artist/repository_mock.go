package artist

import "github.com/stretchr/testify/mock"

// MockRepository is a mock implementation of the Repository interface.
type MockRepository struct {
	mock.Mock
}

// Save mocks the Save method of the Repository interface.
func (m *MockRepository) Save(artist *Artist) (*Artist, error) {
	args := m.Called(artist)
	return args.Get(0).(*Artist), args.Error(1)
}

// GetById mocks the GetById method of the Repository interface.
func (m *MockRepository) GetById(artistId int) (*Artist, error) {
	args := m.Called(artistId)
	return args.Get(0).(*Artist), args.Error(1)
}
