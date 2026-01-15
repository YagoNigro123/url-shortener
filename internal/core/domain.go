package core

import "time"

type Link struct {
	ID        string    `json:"id"`
	Original  string    `json:"original"`
	CreatedAt time.Time `json:"created_at"`
	Visits    int       `json:"visits"`
}

type LinkCache interface {
	Save(id string, original string) error
	Get(id string) (string, error)
}

type LinkStore interface {
	Save(link *Link) error
	Find(id string) (*Link, error)
}
