package main

import (
	"net/http"

	"kaykodesigns.kpkaccounting.net/internal/validator"
)

type PageParams struct {
	Title string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Page = PageParams{Title: "Home"}

	app.render(w, http.StatusOK, "home.tmpl", data)

}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Page = PageParams{Title: "About"}

	app.render(w, http.StatusOK, "about.tmpl", data)

}

func (app *application) services(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Page = PageParams{Title: "Services"}

	app.render(w, http.StatusOK, "services.tmpl", data)

}

type contactForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Phone               string `form:"phone"`
	Content             string `form:"content"`
	validator.Validator `form:"-"`
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Page = PageParams{Title: "Contact"}
	data.Form = contactForm{}

	app.render(w, http.StatusOK, "contact.tmpl", data)
}

func (app *application) ContactFormPost(w http.ResponseWriter, r *http.Request) {
	var form contactForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank.")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank.")
	form.CheckField(validator.Matches(form.Email, validator.Emailrx), "email", "This field must be a valid email address.")
	form.CheckField(validator.NotBlank(form.Phone), "phone", "This field cannot be blank.")
	form.CheckField(validator.Matches(form.Phone, validator.Phonerx), "phone", "This field must be a valid email address.")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank.")

	if !form.Valid() {
		app.infoLog.Println("form not valid.")
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "contact.tmpl", data)
		return
	}

	data := app.newTemplateData(r)
	data.Form = form
	data.Page = PageParams{Title: "Confirmation"}
	app.render(w, http.StatusOK, "confirmation.tmpl", data)

}

// func (form *contactForm) Deliver() error {
// 	email := mail.NewMessage()
// 	email.SetHeader("To", "kayleigh.schlupp@gmail.com")
// 	email.SetHeader("From", "kayleighkat_2@hotmail.com")
// 	email.SetHeader("Reply-To", form.Email)
// 	email.SetHeader("Subject", "New message via KPK Accounting Contact Form")
// 	email.SetBody("text/plain", form.Content)

// 	username := "kayleighkat_2@hotmail.com"
// 	password := "2!BX6699!VickyRose"

// 	return mail.NewDialer("smtp-mail.outlook.com", 587, username, password).DialAndSend(email)
// }

// func (app *application) contactFormPost(w http.ResponseWriter, r *http.Request) {
// 	var form contactForm
// 	err := app.decodePostForm(r, &form)
// 	if err != nil {
// 		app.errorLog.Println("Error decoding form for contactFormPost")
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}

// 	form.CheckField(validator.NotBlank(form.Fullname), "fullname", "Please enter your first and last name.")
// 	form.CheckField(validator.Matches(form.Email, validator.Emailrx), "email", "Please enter a valid email address ie. johndoe@abc.com.")
// 	form.CheckField(validator.Matches(form.Phone, validator.Phonerx), "phone", "Please enter a valid phone number format ie. 123-456-7890.")
// 	form.CheckField(validator.NotBlank(form.Content), "content", "Please describe the service we can assist with.")

// 	// if !form.Valid() {
// 	// 	app.infoLog.Println("Contact form not valid.")
// 	// 	data := app.newTemplateData(r)
// 	// 	data.Form = form
// 	// 	data.Page = PageParams{Title: "Contact"}
// 	// 	app.render(w, http.StatusOK, "contact.tmpl", data)
// 	// 	return
// 	// }

// 	data := app.newTemplateData(r)
// 	data.Form = form
// 	data.Page = PageParams{Title: "Contact"}
// 	app.render(w, http.StatusOK, "contact.tmpl", data)

// 	if err := form.Deliver(); err != nil {
// 		app.errorLog.Println("There was an error submitting the form.")
// 		app.errorLog.Print(err)
// 		http.Error(w, "Sorry, something went wrong with submitting the form.", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, "confirmationTest.tmpl", http.StatusSeeOther)

// 	app.infoLog.Println("Submitted")
// }
