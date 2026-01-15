package core

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	ErrLinkNotFound = errors.New("Link not found")
	ErrLinkExists   = errors.New("Link already exists")
)

type Service struct {
	store LinkStore
	cache LinkCache
}

func NewService(store LinkStore, cache LinkCache) *Service {
	return &Service{
		store: store,
		cache: cache,
	}
}

func (s *Service) Shorten(originalURL string) (*Link, error) {
	id := generateShortID()

	link := &Link{
		ID:        id,
		Original:  originalURL,
		CreatedAt: time.Now(),
		Visits:    0,
	}

	if err := s.store.Save(link); err != nil {
		return nil, fmt.Errorf("Error saving to db: %w", err)
	}

	go func() {
		if err := s.cache.Save(link.ID, link.Original); err != nil {
			log.Printf("Error saving to cache: %v", err)
		}
	}()

	return link, nil
}

func (s *Service) GetOriginal(id string) (*Link, error) {
	cacheURL, err := s.cache.Get(id)

	if err == nil {
		return &Link{ID: id, Original: cacheURL}, nil
	}

	link, err := s.store.Find(id)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := s.cache.Save(link.ID, link.Original); err != nil {
			log.Printf("Error updating cache: %v", err)
		}
	}()

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
