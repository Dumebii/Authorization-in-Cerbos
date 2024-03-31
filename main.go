package main

import (
	"fmt"
	"net/http"
	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))
var c, err = cerbos.New("localhost:3593", cerbos.WithPlaintext())
	


func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/dashboard", dashboardHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// Check credentials
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validate credentials
	if username == "user" && password == "password" {
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		fmt.Fprintln(w, "Invalid credentials. Please try again.")
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {

	
	session, _ := store.Get(r, "session-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Check access using Cerbos
	principal := cerbos.NewPrincipal("admin", "user")
	resource := cerbos.NewResource("dashboard", "user_profile")
	decision, err := c.IsAllowed(r.Context(), principal, resource, "update")
	if err != nil {
		http.Error(w, "Error checking access", http.StatusInternalServerError)
		return
	}

	if decision {
		fmt.Fprintln(w, "Welcome to the dashboard!")
	} else {
		http.Error(w, "Access denied", http.StatusForbidden)
	}
}

