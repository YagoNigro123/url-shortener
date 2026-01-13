package store 

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/YagoNigro123/url-shortener/internal/core"
)

type PostgressStore struct {
	db *sql.DB
}

func NewPostgresStore(connString string) (*PostgressStore, error){
	db, err := sql.Open("postgres",connString)
	if err != nil{
		return nil, fmt.Errorf("unable to open db connection: %w", err)
	}

	if err := db.Ping(); err != nil{
		return nil, fmt.Errorf("uneable to connect to db: %w", err)
	}

	return &PostgressStore{db:db}, nil
}

func (s*PostgressStore) Save(link *core.Link) error {
	query := `
		INSERT INTO links (id, original_url, created_at, visits)
		VALUES ($1,$2,$3,$4)
	`
	_, err := s.db.Exec(
		query,
		link.ID,
		link.Original,
		link.CreatedAt,
		link.Visits,
	)

	if err != nil {
		return fmt.Errorf("error saving link: %w", err)
	}

	return nil
}

func (s*PostgressStore) Find(id string) (*core.Link, error){
	query := `SELECT id, original_url, created_at, visits FROM links WHERE id = $1`

	row := s.db.QueryRow(query, id)

	link := &core.Link{}

	err := row.Scan(
		&link.ID,
		&link.Original,
		&link.CreatedAt,
		&link.Visits,
	)

	if err != nil{
		if err == sql.ErrNoRows {
			return nil, core.ErrLinkNotFound
		}
		return nil, fmt.Errorf("error finding link: %w", err)
	}

	return link, nil
}
