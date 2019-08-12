package pages

import (
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth,
		Css:    "",
		Path:   "signup",
		Data:   Signup,
		Extra:  0,
	}
	router.Add(b)
}

func Signup(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Sign Up"),
		"mainHeader": p.Sprintf("Sign Up for free"),
		"fname":      p.Sprintf("First Name:"),
		"fnameph":    p.Sprintf("Enter your first name"),
		"lnameph":    p.Sprintf("Enter your last name"),
		"lname":      p.Sprintf("Last Name:"),
		"email":      p.Sprintf("Email:"),
		"emailph":    p.Sprintf("Enter a valid email"),
		"pass":       p.Sprintf("Password:"),
		"passph":     p.Sprintf("Enter password (8+ chars)"),
		"gobutton":   p.Sprintf("Sign Up"),
	}
}
