package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"kaykodesigns.kpkaccounting.net/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	dynamic := alice.New(app.sessionManager.LoadAndSave)
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodPost, "/", dynamic.ThenFunc(app.homeContactFormPost))
	router.Handler(http.MethodGet, "/about", dynamic.ThenFunc(app.about))
	router.Handler(http.MethodGet, "/services", dynamic.ThenFunc(app.services))
	router.Handler(http.MethodGet, "/contact", dynamic.ThenFunc(app.contact))
	router.Handler(http.MethodPost, "/contact", dynamic.ThenFunc(app.contactFormPost))

	standard := alice.New(app.recoverPanic, app.logRequest)
	return standard.Then(router)
}
