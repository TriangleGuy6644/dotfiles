package handlers
import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"goapi/internal/middleware"

)


func Handler(r *chi.Mux) {
	//Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account",	func(router chi.Router) {
		router.Use(middleware.Authorization)
	}
}