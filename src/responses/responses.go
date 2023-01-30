package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)

	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}

}

func Error(w http.ResponseWriter, status int, error error) {
	JSON(w, status, struct {
		Error string `json:"error"`
	}{
		Error: error.Error(),
	})

}
