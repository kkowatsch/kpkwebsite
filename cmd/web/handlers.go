package main

import (
	"fmt"
	"net/http"

	"github.com/go-mail/mail"
	"kaykodesigns.kpkaccounting.net/internal/validator"
)

type PageParams struct {
	Title       string
	Description string
}

type contactForm struct {
	Name                string `form:"name"`
	Company             string `form:"company"`
	Email               string `form:"email"`
	Phone               string `form:"phone"`
	Content             string `form:"content"`
	validator.Validator `form:"-"`
}

func (form *contactForm) Deliver() error {
	username := "ken@kpkaccounting.com"
	password := "20Springhill23$"

	// Email Header Details
	email := mail.NewMessage()
	email.SetHeader("To", username)
	email.SetHeader("From", "ken@kpkaccounting.com")
	email.SetHeader("Reply-To", form.Email)
	email.SetHeader("Subject", "New message via KPK Accounting Contact Form")

	// Email Body
	body := fmt.Sprintf("Hello,\n\nA new contact form has been submitted on your website.\n\nName: %s\nCompany Name: %s\nEmail Address: %s\nPhone Number: %s\nMessage: %s\n\nThank you,\nYour Website", form.Name, form.Company, form.Email, form.Phone, form.Content)
	email.SetBody("text/plain", body)

	return mail.NewDialer("smtp.office365.com", 587, username, password).DialAndSend(email)
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Page = PageParams{Title: "Home", Description: "The KPK Accounting Inc. home page with quick highlights of the services provided, the story of KPK's CEO and quick access for submitting inquiries through the contact form."}
	data.Form = contactForm{}

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) homeContactFormPost(w http.ResponseWriter, r *http.Request) {
	var form contactForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "This field requires at least one non-space character.")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field requires at least one non-space character.")
	form.CheckField(validator.Matches(form.Email, validator.Emailrx), "email", "Please enter a valid email address ie. jdoe@abc.com")
	form.CheckField(validator.NotBlank(form.Phone), "phone", "This field requires at least one non-space character.")
	form.CheckField(validator.Matches(form.Phone, validator.Phonerx), "phone", "Please enter a valid phone number format ie. 123-456-7890")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field requires at least one non-space character.")

	if !form.Valid() {
		app.infoLog.Println("form not valid.")
		data := app.newTemplateData(r)
		data.Form = form
		data.Page = PageParams{Title: "Home", Description: "The KPK Accounting Inc. home page with quick highlights of the services provided, the story of KPK's CEO and quick access for submitting inquiries through the contact form."}
		app.render(w, http.StatusUnprocessableEntity, "home.tmpl", data)
		return
	}

	data := app.newTemplateData(r)
	data.Form = form
	data.Page = PageParams{Title: "Confirmation", Description: "This page is visible because an inquiry through the contact form has been submitted to KPK Accounting Inc."}
	app.render(w, http.StatusOK, "confirmation.tmpl", data)

	if err := form.Deliver(); err != nil {
		app.errorLog.Println("There was an error submitting the form.")
		app.errorLog.Print(err)
		http.Error(w, "Sorry, something went wrong with submitting the form.", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Submitted")
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Page = PageParams{Title: "About", Description: "The story of the KPK Accounting Inc. CEO, Ken Kowatsch, and his approach to providing business owners, entrepreneurs and start-ups with the tools for long-term financial health."}

	app.render(w, http.StatusOK, "about.tmpl", data)

}

func (app *application) services(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Page = PageParams{Title: "Services", Description: "The services provided by KPK Accounting Inc. to provide long-term financial health. Some services offered are Cash Flow Optimization, Audit Leadership, Risk Identification, KPI Reporting Implementation, etc."}

	app.render(w, http.StatusOK, "services.tmpl", data)

}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Page = PageParams{Title: "Contact", Description: "The KPK Accounting Inc. contact form where clients can easily connect with KPK Accounting regarding financial inquiries."}
	data.Form = contactForm{}

	app.render(w, http.StatusOK, "contact.tmpl", data)
}

func (app *application) contactFormPost(w http.ResponseWriter, r *http.Request) {
	var form contactForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "This field requires at least one non-space character.")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field requires at least one non-space character.")
	form.CheckField(validator.Matches(form.Email, validator.Emailrx), "email", "Please enter a valid email address ie. jdoe@abc.com")
	form.CheckField(validator.NotBlank(form.Phone), "phone", "This field requires at least one non-space character.")
	form.CheckField(validator.Matches(form.Phone, validator.Phonerx), "phone", "Please enter a valid phone number format ie. 123-456-7890")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field requires at least one non-space character.")

	if !form.Valid() {
		app.infoLog.Println("form not valid.")
		data := app.newTemplateData(r)
		data.Form = form
		data.Page = PageParams{Title: "Contact", Description: "The KPK Accounting Inc. contact form where clients can easily connect with KPK Accounting regarding financial inquiries."}
		app.render(w, http.StatusUnprocessableEntity, "contact.tmpl", data)
		return
	}

	data := app.newTemplateData(r)
	data.Form = form
	data.Page = PageParams{Title: "Confirmation", Description: "This page is visible because an inquiry through the contact form has been submitted to KPK Accounting Inc."}
	app.render(w, http.StatusOK, "confirmation.tmpl", data)

	if err := form.Deliver(); err != nil {
		app.errorLog.Println("There was an error submitting the form.")
		app.errorLog.Print(err)
		http.Error(w, "Sorry, something went wrong with submitting the form.", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Submitted")
}
