package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, status int, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(resp))
}

func SendErr(w http.ResponseWriter, errcode int, message string) {
	SendResponse(w, errcode, map[string]string{"error": message})
}
