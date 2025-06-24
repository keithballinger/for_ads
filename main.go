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
	http.HandleFunc("/api/sonnet", sonnetHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func sonnetHandler(w http.ResponseWriter, r *http.Request) {
	// Allow cross-origin requests from the frontend
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

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
	json.NewEncoder(w).Encode(sonnet)
}
