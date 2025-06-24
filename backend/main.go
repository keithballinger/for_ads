package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Sonnet struct {
	Title string   `json:"title"`
	Lines []string `json:"lines"`
}

func main() {
	// Serve the frontend
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// API endpoint
	http.HandleFunc("/api/sonnet", sonnetHandler)

	// Determine port for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func sonnetHandler(w http.ResponseWriter, r *http.Request) {
	sonnet := Sonnet{
		Title: "Sonnet 116",
		Lines: []string{
			"Let me not to the marriage of true minds",
			"Admit impediments. Love is not love",
			"Which alters when it alteration finds,",
			"Or bends with the remover to remove:",
			"O no! it is an ever-fixed mark",
			"That looks on tempests and is never shaken;",
			"It is the star to every wandering bark,",
			"Whose worth's unknown, although his height be taken.",
			"Love's not Time's fool, though rosy lips and cheeks",
			"Within his bending sickle's compass come:",
			"Love alters not with his brief hours and weeks,",
			"But bears it out even to the edge of doom.",
			"If this be error and upon me proved,",
			"I never writ, nor no man ever loved.",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sonnet)
}
