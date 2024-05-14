package postrepository

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/domain"
)

type postRepository struct {
	pool *pgxpool.Pool
	log  *slog.Logger
}

func NewPostRepository(p *pgxpool.Pool, l *slog.Logger) *postRepository {
	return &postRepository{p, l}
}

func (p *postRepository) Create(post domain.Post) (uint64, error) {
	query := `
	INSERT INTO posts (title, content, author_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	var id uint64
	err := p.pool.QueryRow(
		context.Background(),
		query,
		post.Title,
		post.Content,
		post.AuthorId,
		post.CreatedAt,
		post.UpdatedAt,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *postRepository) Delete(id uint64) (uint64, error) {
	panic("unimplemented")
}

func (p *postRepository) Get(id uint64) (domain.Post, error) {
	query := `SELECT * FROM posts WHERE id = $1`

	var post domain.Post
	err := p.pool.QueryRow(context.Background(), query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorId,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

func (p *postRepository) GetBatch(lastPostId uint64, limit int) ([]domain.Post, error) {
	panic("unimplemented")
}

func (p *postRepository) Update(id uint64, post domain.Post) (uint64, error) {
	panic("unimplemented")
}
