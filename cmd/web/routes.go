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
	router.Handler(http.MethodGet, "/kpk-about", dynamic.ThenFunc(app.about))
	router.Handler(http.MethodGet, "/kpk-services", dynamic.ThenFunc(app.services))
	router.Handler(http.MethodGet, "/kpk-contact", dynamic.ThenFunc(app.contact))
	router.Handler(http.MethodGet, "/kpk-contactform", dynamic.ThenFunc(app.contactFormPost))

	standard := alice.New(app.recoverPanic, app.logRequest)
	return standard.Then(router)
}
