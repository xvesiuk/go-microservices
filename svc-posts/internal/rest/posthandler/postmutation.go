package posthandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/xvesiuk/go-microservices/svc-posts/internal/domain"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/rest"
)

type postMutationBody struct {
	Title   string `json:"title" validate:"required,lte=50"`
	Content string `json:"content"  validate:"required,lte=300"`
}

type postMutationResponse struct {
	ID uint64 `json:"id"`
}

// @Tags         post
// @Accept       json
// @Produce      json
// @Param        data body postMutationBody true "post payload"
// @Success      201  {object}  rest.Response[postMutationResponse]
// @Failure      400  {object}  rest.ResponseDefault
// @Failure      401  {object}  rest.ResponseDefault
// @Failure      500  {object}  rest.ResponseDefault
// @Router       /post [post]
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		rest.WriteError(w, 400, rest.ErrDecodeBody)
		return
	}

	var payload postMutationBody
	err = json.Unmarshal(body, &payload)
	if err != nil {
		rest.WriteError(w, 400, rest.ErrParseBody)
		return
	}

	// validator.New()
	id, err := h.postService.Create(domain.Post{
		Title:    payload.Title,
		Content:  payload.Content,
		AuthorId: 1,
	})
	if err != nil {
		h.log.Error(err.Error())
		rest.WriteError(w, 500, rest.ErrInternal)
		return
	}

	err = rest.WriteOk(w, 201, postMutationResponse{id})
	if err != nil {
		rest.WriteError(w, 500, rest.ErrInternal)
		return
	}
}
