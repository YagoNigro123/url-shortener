package core

import "time"

type Link struct {
	ID        string    `json:"id"`
	Original  string    `json:"original"`   // <--- Agregué comillas
	CreatedAt time.Time `json:"created_at"` // <--- Corregí el nombre (Agregué la 'd') y las comillas
	Visits    int       `json:"visits"`     // <--- Agregué comillas
}

type LinkStore interface {
	Save(link *Link) error
	Find(id string) (*Link, error)
}
