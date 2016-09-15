package resources

import (
	"github.com/pressly/chi"
	"github.com/johnmcdnl/darts/resources/targets"
	//"github.com/johnmcdnl/auth/auth"
)

func ResourcesRouter() chi.Router {
	r := chi.NewRouter()
	//r.Use(auth.ValidateHandler)
	r.Mount("/targets", targets.TargetsRouter())
	return r
}
