package data

import (
	"database/sql"
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
	return nil
}

func (m PostModel) Get(id int64) (*Post, error) {
	return nil, nil
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
