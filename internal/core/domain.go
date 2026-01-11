package core

import "time"

type Link struct {
	ID       string    `json:"id"`
	Original string    `json:original` // https://google.com
	CreateAt time.Time `json:create_at`
	Visits   int       `json:visits`
}

type LinkStore interface {
	Save(link *Link) error
	Find(id string) (*Link, error)
}
