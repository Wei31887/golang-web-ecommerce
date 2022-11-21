package main

import (
	"SideProject/store/routers"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./template/style/"))))

	// handler to show the web page
	http.HandleFunc("/", routers.Index) 
	http.HandleFunc("/login", routers.Login)
	http.HandleFunc("/signup", routers.Signup)

	// handler for admin 
	http.HandleFunc("/admin", routers.AdminPage)
	http.HandleFunc("/admin/uploadproduct", routers.AdminAddProduct)
	// http.HandleFunc("/admin/editproduct", routers.AdminRemoveProduct)

	// handler to deal with user authenticate
	http.HandleFunc("/signupAccount", routers.SignupAccount)
	http.HandleFunc("/authenticate", routers.Authenticate)
	http.HandleFunc("/logout", routers.Logout)

	server.ListenAndServe()
}
