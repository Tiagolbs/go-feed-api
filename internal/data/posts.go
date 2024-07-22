package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/tiagolbs/go-feed-api/internal/validator"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostModel struct {
	DB *sql.DB
}

func (m PostModel) Insert(post *Post) error {
	query := `
		INSERT INTO posts (content)
		VALUES ($1)
		RETURNING id, created_at, updated_at
	`

	args := []interface{}{post.Content}

	return m.DB.QueryRow(query, args...).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}

func (m PostModel) Get(id int64) (*Post, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, updated_at, content
		FROM posts
		WHERE id = $1
	`

	var post Post

	err := m.DB.QueryRow(query, id).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Content,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (m PostModel) Update(post *Post) error {
	return nil
}

func (m PostModel) Delete(id int64) error {
	return nil
}

func ValidatePost(v *validator.Validator, post *Post) {
	v.Check(post.Content != "", "content", "must be provided")
}
