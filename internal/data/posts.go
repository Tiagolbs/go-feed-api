package data

import (
	"time"

	"github.com/tiagolbs/go-feed-api/internal/validator"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ValidatePost(v *validator.Validator, post *Post) {
	v.Check(post.Content != "", "content", "must be provided")
}
