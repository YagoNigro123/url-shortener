package core_test

import (
	"testing"

	"github.com/YagoNigro123/url-shortener/internal/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Save(link *core.Link) error {
	args := m.Called(link)
	return args.Error(0)
}

func (m *MockStore) Find(id string) (*core.Link, error) {
	args := m.Called(id)
	return args.Get(0).(*core.Link), args.Error(1)
}

func TestShortenURL(t *testing.T) {
	mockStore := new(MockStore)
	service := core.NewService(mockStore)

	urlOriginal := "https://www.google.com"

	mockStore.On("Save", mock.Anything).Return(nil)

	result, err := service.Shorten(urlOriginal)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 6, len(result.ID))
	assert.Equal(t, urlOriginal, result.Original)

	mockStore.AssertExpectations(t)
}
