package posthandler

import (
	"log/slog"
	"net/http"

	"github.com/xvesiuk/go-microservices/svc-posts/internal/domain"
)

type PostService interface {
	Create(post domain.Post) (uint64, error)
	Update(id uint64, post domain.Post) (uint64, error)
	Delete(id uint64) (uint64, error)
	Get(id uint64) (domain.Post, error)
	GetBatch(lastPostId uint64, limit int) ([]domain.Post, error)
}

type Handler struct {
	postService PostService
	log         *slog.Logger
}

func NewHandler(p PostService, l *slog.Logger) *Handler {
	return &Handler{p, l}
}

func (h *Handler) GetBatch(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
