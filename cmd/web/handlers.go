package main

import (
	"net/http"

	"github.com/go-mail/mail"
	"kaykodesigns.kpkaccounting.net/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "home.tmpl", data)

}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "about.tmpl", data)

}

func (app *application) services(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "services.tmpl", data)

}

type contactForm struct {
	Fullname            string `form:"fullname"`
	Cname               string `form:"cname"`
	Email               string `form:"email"`
	Phone               string `form:"phone"`
	Content             string `form:"content"`
	validator.Validator `form:"-"`
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "contact.tmpl", data)
}

func (form *contactForm) Deliver() error {
	email := mail.NewMessage()
	email.SetHeader("To", "kayleigh.schlupp@gmail.com")
	email.SetHeader("From", "server@example.com")
	email.SetHeader("Reply-To", form.Email)
	email.SetHeader("Subject", "New message via KPK Accounting Contact Form")
	email.SetBody("text/plain", form.Content)

	username := "your_username"
	password := "your_password"

	return mail.NewDialer("smtp.mailtrap.io", 25, username, password).DialAndSend(email)
}

func (app *application) contactFormPost(w http.ResponseWriter, r *http.Request) {
	var form contactForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.errorLog.Println("Error decoding form for contactFormPost")
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Fullname), "fullname", "Please enter your first and last name.")
	form.CheckField(validator.Matches(form.Email, validator.Emailrx), "email", "Please enter a valid email address ie. johndoe@abc.com.")
	form.CheckField(validator.Matches(form.Phone, validator.Phonerx), "phone", "Please enter a valid phone number format ie. 123-456-7890.")
	form.CheckField(validator.NotBlank(form.Content), "content", "Please describe the service we can assist with.")

	if !form.Valid() {
		app.infoLog.Println("Contact form not valid.")
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "contact.tmpl", data)
		return
	}

	if err := form.Deliver(); err != nil {
		app.errorLog.Println("There was an error submitting the form.")
		app.errorLog.Print(err)
		http.Error(w, "Sorry, something went wrong with submitting the form.", http.StatusInternalServerError)
		return
	}

	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "contact.tmpl", data)
}
