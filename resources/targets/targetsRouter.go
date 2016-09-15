package targets

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

const targetNameUrl = "/target-name/:target-name"
const usernameNameUrl = "/users/:username"

func TargetsRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", createTargetsMatch)

	r.Get("/", allTargetsHandler)

	r.Get(targetNameUrl, targetNameHandler)

	r.Get(usernameNameUrl, usernameHandler)
	r.Get(usernameNameUrl+targetNameUrl, usernameTargetNameHandler)

	return r
}

func usernameParam(r *http.Request) string {
	return chi.URLParam(r, "username")
}

func targetNameParam(r *http.Request) string {
	return chi.URLParam(r, "target-name")
}

func allTargetsHandler(w http.ResponseWriter, r *http.Request) {
	targets, err := retrieveAllTargets()

	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, targets)
}

func usernameHandler(w http.ResponseWriter, r *http.Request) {
	targets, err := retrieveAllTargetsByUsername(usernameParam(r))
	if err != nil {

		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, targets)
}

func targetNameHandler(w http.ResponseWriter, r *http.Request) {
	targets, err := retrieveAllTargetsByTargetName(targetNameParam(r))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, targets)
}

func usernameTargetNameHandler(w http.ResponseWriter, r *http.Request) {

	targets, err := retrieveAllTargetsByUsernameAndTargetName(usernameParam(r), targetNameParam(r))

	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, targets)
}

func createTargetsMatch(w http.ResponseWriter, r *http.Request) {
	t, err := bindTargetFromRequest(r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	if t, err = createTarget(t); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, t)
}
