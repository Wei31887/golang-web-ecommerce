package routers

import (
	"SideProject/store/models"
	"fmt"
	"net/http"
)

// POST
// SignupAccount : sign up the new user
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	user := models.User{
		Name: 		r.PostFormValue("name"),
		Password: 	r.PostFormValue("password"),
		Email: 		r.PostFormValue("email"),
	}

	// check if the email has been used or not
	if _, ok := models.GetUserByEmail(user.Email); ok == nil {
		// email has been used 
		http.Error(w, "Email has been used", http.StatusForbidden)
		http.Redirect(w, r, "/signup", http.StatusFound)
	} else {
		if err := user.Create(); err != nil {
			panic(err)
		}
		// sucessfully create the account 
		http.Redirect(w, r, "/login", http.StatusFound)
	}	
}

// POST /Authenticate
// Authenticate the user given the email and password
func Authenticate(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	// is there the email in database?
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		http.Error(w, "Email is invalid", http.StatusForbidden)
		return
	}

	if r.PostFormValue("password") == user.Password {
		// create session for login user
		session, err := user.CreateSession()
		if err != nil {
			panic(err)
		}

		// set the cookie-based session for user
		c := &http.Cookie{
			Name: "session",
			Value: session.Uuid,
			HttpOnly: true,
		}
		
		http.SetCookie(w, c)
		
		// 
		if user.Role == 1 {
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else {
		http.Error(w, "Password is invalid", http.StatusForbidden)
		// http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// POST/
// Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		fmt.Println("no cookie")
	} else {
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUuid()
	}

	// remove the client cookie
	cookie = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	// renew the sessions 
	models.UpdateSession()

	// redirect the index page
	http.Redirect(w, r, "/", http.StatusFound)
}