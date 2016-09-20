package resources

import (
	"github.com/johnmcdnl/auth/auth"
	"github.com/johnmcdnl/darts/resources/board"
	"github.com/johnmcdnl/darts/resources/targets"
	"github.com/johnmcdnl/darts/utils/handlers"
	"github.com/pressly/chi"
)

func ResourcesRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(handlers.CORSHandler().Handler)
	r.Use(auth.ValidateHandler)
	r.Mount("/targets", targets.TargetsRouter())
	r.Mount("/board", board.BoardsRouter())
	return r
}
