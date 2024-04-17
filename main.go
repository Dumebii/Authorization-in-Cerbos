package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	
	decisions, err := c.CheckResourceSet(
	  ctx,
	  client.NewPrincipal("user#1").
		WithRoles("user").
		WithAttributes(map[string]interface{}{
		}),
	  client.NewResourceSet("dashboardResource").
		AddResourceInstance("dashboardResource#1", map[string]interface{}{
		}), "create", "delete", "read", "update")
	  
	
	if err != nil {
	  log.Fatalf("Error while checking resource set: %v", err)
	}
	
	
	for _, action := range []string{"create", "delete", "read", "update"} {
	  fmt.Printf("\t%s -> %t\n", action, decisions.IsAllowed("dashboardResource#1", action))
	}
}

