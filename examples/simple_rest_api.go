package examples

import (
	"encoding/json"
	"log"
	"net/http"
)

func RunSimpleRestApiExample() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(map[string]string{"message": "OK"})
	})
	log.Println("Running on port 3000")
	_ = http.ListenAndServe(":8080", nil)
}
