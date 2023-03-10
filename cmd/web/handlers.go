package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	app.render(w, http.StatusOK, "home.tmpl", data)
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About KPK Accounting"))

}

func services(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Services offered from KPK Accounting"))

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact KPK Accounting"))
}

func contactformpost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Contact Form sending..."))
}
