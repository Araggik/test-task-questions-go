package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

func receiveIdFromPath(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")

	rawId := parts[2]

	return strconv.Atoi(rawId)
}

func writeJSON(w http.ResponseWriter, result []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
