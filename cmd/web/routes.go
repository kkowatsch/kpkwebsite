package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/kpk-about", dynamic.ThenFunc(app.about))
	router.Handler(http.MethodGet, "/kpk-services", dynamic.ThenFunc(app.services))
	router.Handler(http.MethodGet, "/kpk-contact", dynamic.ThenFunc(app.contact))
	router.Handler(http.MethodGet, "/kpk-contactform", dynamic.ThenFunc(app.contactFormPost))

	standard := alice.New(app.recoverPanic, app.logRequest)
	return standard.Then(router)

}
