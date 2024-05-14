package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/xvesiuk/go-microservices/svc-posts/api/docs"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/config"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/db"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/db/repository/postrepository"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/rest"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/rest/middleware"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/rest/posthandler"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/service/postservice"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8001
// @BasePath  /api/v1
func main() {
	// TODO: add config
	conf := config.New()

	// TODO: better setup logger
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	pool, err := db.NewPool(&conf.Database)
	if err != nil {
		panic("DB can't Connect")
	}

	r := chi.NewRouter()

	r.Use(middleware.NewLogger(log))
	r.Use(chimiddleware.Recoverer) // TODO: learn how Recoverer works
	r.Use(chimiddleware.URLFormat)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Ok"))
		})

		r.Get("/docs/*", httpSwagger.Handler(
			httpSwagger.URL(
				fmt.Sprintf("http://localhost:%d/api/v1/docs/doc.json", conf.Port),
			),
		))

		r.Route("/post", func(r chi.Router) {
			repo := postrepository.NewPostRepository(pool, log)
			service := postservice.NewService(repo, log)
			h := posthandler.NewHandler(service, log)

			r.Get("/{postID}", h.Get)
			r.Get("/", rest.NotImplementedHandler)

			r.Group(func(r chi.Router) {
				// TODO: protect with JWT
				r.Post("/", h.Post)
				r.Put("/{postID}", rest.NotImplementedHandler)
				r.Delete("/{postID}", rest.NotImplementedHandler)
			})
		})
	})

	// TODO: gracefull shutdown
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r)

}
