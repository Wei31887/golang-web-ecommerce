package routers

import (
	"SideProject/store/models"
	"net/http"
)

// GET /
// Index : show the login page to client
func Index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		// There is no session: show the public version index
		tpl := parseTemplateFile("index", "public-nav")
		tpl.ExecuteTemplate(w, "index", nil)
	} else {
		uuid := cookie.Value
		_, err := models.GetUserByUuid(uuid)
		if err != nil {
			// user has not log in yet: show the public version index
			tpl := parseTemplateFile("index", "public-nav")
			tpl.ExecuteTemplate(w, "index", nil)

		} else {
			// user has log: show the private version index
			tpl := parseTemplateFile("index", "private-nav")
			tpl.ExecuteTemplate(w, "index", nil)
		}
	}
}

// GET /
// Login : show the login page to client
func Login(w http.ResponseWriter, r *http.Request) {
	tpl := parseTemplateFile("login-layout", "login", "public-nav")
	tpl.ExecuteTemplate(w, "login-layout", nil)
}

// GET /
// Signup : show the signup page to client
func Signup(w http.ResponseWriter, r *http.Request) {
	tpl := parseTemplateFile("signup-layout", "signup", "public-nav")
	tpl.ExecuteTemplate(w, "signup-layout", nil)
}

