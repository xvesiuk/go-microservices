package postservice

import (
	"log/slog"

	"github.com/xvesiuk/go-microservices/svc-posts/internal/domain"
)

type PostRepository interface {
	Create(post domain.Post) (uint64, error)
	Update(id uint64, post domain.Post) (uint64, error)
	Delete(id uint64) (uint64, error)
	Get(id uint64) (domain.Post, error)
	GetBatch(lastPostId uint64, limit int) ([]domain.Post, error)
}

type Service struct {
	postRepository PostRepository
	log            *slog.Logger
}

func NewService(p PostRepository, l *slog.Logger) *Service {
	return &Service{p, l}
}

func (s *Service) Create(post domain.Post) (uint64, error) {
	// notify subscribers
	return s.postRepository.Create(post)
}

func (s *Service) Delete(id uint64) (uint64, error) {
	return s.postRepository.Delete(id)
}

func (s *Service) GetBatch(lastPostId uint64, limit int) ([]domain.Post, error) {
	return s.postRepository.GetBatch(lastPostId, limit)
}

func (s *Service) Update(id uint64, post domain.Post) (uint64, error) {
	return s.postRepository.Update(id, post)
}

func (s *Service) Get(id uint64) (domain.Post, error) {
	return s.postRepository.Get(id)
}
