package forms

import (
	"net/http"
	"github.com/albrow/forms"
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Form{
		Access: router.GuestAuth,
		Path: "login",
		Validator: Login,
		After: router.ValidateLogin,
		Redirect: "/profile.html",
	}
	router.AddPost(b)
}

func Login(r *http.Request, p *message.Printer) []string {
	var errors []string
	data, err := forms.Parse(r)
	if err != nil {
		errors = append(errors, p.Sprint("Internal server error"))
		return errors
	}
	val := data.Validator()
	val.Require("email").Message(p.Sprint("Username required"))
	val.MatchEmail("email").Message(p.Sprint("User name must be a valid email"))
	val.Require("pass").Message(p.Sprint("Password required"))
	val.MinLength("pass", 8).Message(p.Sprint("Password must be at least 8 characters"))
	if val.HasErrors() {
		errors = append(errors, val.Messages()...)
	}
	return errors
}
