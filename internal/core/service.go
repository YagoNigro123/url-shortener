package core

import (
	"errors"
	"math/rand"
	"time"
)

var (
	ErrLinkNotFound = errors.New("Link not found")
	ErrLinkExists   = errors.New("Link already exists")
)

type Service struct {
	store LinkStore
}

func NewService(store LinkStore) *Service {
	return &Service{store: store}
}

func (s *Service) Shorten(originalURL string) (*Link, error) {
	id := generateShortID()

	link := &Link{
		ID:        id,
		Original:  originalURL,
		CreatedAt: time.Now(),
		Visits:    0,
	}

	err := s.store.Save(link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func generateShortID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *Service) GetOriginal(id string) (*Link, error) {
	return s.store.Find(id)
}
