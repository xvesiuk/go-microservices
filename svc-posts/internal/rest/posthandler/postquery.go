package posthandler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/domain"
	_ "github.com/xvesiuk/go-microservices/svc-posts/internal/domain"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/rest"
)

type postQueryResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  uint64    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newPostQueryResponse(post domain.Post) postQueryResponse {
	return postQueryResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorId:  post.AuthorId,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

// @Tags         post
// @Produce      json
// @Param        postID path int true "post id"
// @Success      200  {object}  rest.Response[postQueryResponse]
// @Failure      400  {object}  rest.ResponseDefault
// @Failure      401  {object}  rest.ResponseDefault
// @Failure      500  {object}  rest.ResponseDefault
// @Router       /post/{postID} [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		rest.WriteError(w, 400, err)
		return
	}

	post, err := h.postService.Get(uint64(postID))
	if err != nil {
		rest.WriteError(w, 500, err)
		return
	}

	rest.WriteOk(w, 200, newPostQueryResponse(post))
}
