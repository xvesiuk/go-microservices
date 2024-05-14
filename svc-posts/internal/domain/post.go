package domain

import "time"

type Post struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  uint64    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
