package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
}

func main() {
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Println("Error creating template")
		errorLog.Fatal(err)
	}

	sessionManager := scs.New()

	http.HandleFunc("/about", about)
	http.HandleFunc("/services", services)
	http.HandleFunc("/contact", contact)
	// sends email to kpkaccounting@___.com
	http.HandleFunc("/contact/form", contactformpost)

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		templateCache:  templateCache,
		sessionManager: sessionManager,
	}

	srv := &http.Server{
		Addr:     "localhost:4000",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	log.Printf("Starting server on 4000")
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
