package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Firebase SDK
		opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		authHeader := r.Header.Get("Authorization")
		fmt.Printf(r.Header.Get("Authrization"))
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		log.Printf("%v", authHeader)
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Printf("error veryfying ID Token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			
			w.Write([]byte("error veryfying ID token\n"))
			return
		}
		log.Printf("verified ID token: %v\n", token)
		next.ServeHTTP(w, r)

	}
}


func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello private"))
}

func main() {
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/public", public)
	r.HandleFunc("/private", authMiddleware(private))

	
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
