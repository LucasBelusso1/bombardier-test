package chiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func Start() {
	r := chi.NewRouter()

	r.Get("/health", health)
	r.Post("/withBodyAndHeader", withBodyAndHeader)

	port := ":8083"
	log.Printf("Chi server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

type RequestBody struct {
	Message string `json:"message"`
}

func withBodyAndHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	apiKey := r.Header.Get("x-api-key")
	if apiKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"message": "Missing header x-xpi-key"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if _, err := uuid.Parse(apiKey); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"message": "Invalid x-api-key"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var body RequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"message": "Invalid JSON sent!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
}
