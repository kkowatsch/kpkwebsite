package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/joho/godotenv"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	sessionManager *scs.SessionManager
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // Default port if not specified
	}

	addr := flag.String("addr", ":"+port, "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		sessionManager: sessionManager,
		templateCache:  templateCache,
		formDecoder:    formDecoder,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
