package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/cerbos/cerbos-sdk-go/cerbos"
)

func main() {
	// Initialize Cerbos client
	client, err := cerbos.New("unix:/var/sock/cerbos", cerbos.WithTLSCACert("/path/to/ca.crt"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a new Gorilla router
	r := mux.NewRouter()

	// Define a Gorilla handler function
	r.HandleFunc("/resource/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Extract parameters from the request
		vars := mux.Vars(r)
		resourceID := vars["id"]

		// Perform authorization check
		allowed, err := client.IsAllowed(context.TODO(), cerbos.NewPrincipal("sally").WithRoles("user"), cerbos.NewResource("album:object", resourceID), "view")
		if err != nil {
			http.Error(w, "Authorization check failed", http.StatusInternalServerError)
			return
		}

		if !allowed {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		// Authorized - Proceed with handling the request
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorized to view resource"))
	})

	// Start the HTTP server with the Gorilla router
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
