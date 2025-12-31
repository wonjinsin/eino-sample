package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	custommiddleware "github.com/wonjinsin/eino-sample/internal/handler/http/middleware"
	"github.com/wonjinsin/eino-sample/internal/usecase"
)

// NewRouter creates and configures a new chi router
func NewRouter(
	basicChatSvc usecase.BasicChatService,
) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(custommiddleware.TrID())
	r.Use(custommiddleware.CORS())
	r.Use(middleware.RealIP)
	r.Use(custommiddleware.HTTPLogger())
	r.Use(middleware.Recoverer)

	// Controllers
	healthCtrl := NewHealthController()
	basicChatCtrl := NewBasicChatController(basicChatSvc)

	// Routes
	r.Get("/healthz", healthCtrl.Check)

	// Basic chat routes
	r.Route("/basic-chat", func(r chi.Router) {
		r.Post("/", basicChatCtrl.AskBasicChat)
		r.Post("/prompt-template", basicChatCtrl.AskBasicPromptTemplateChat)
	})

	return r
}
