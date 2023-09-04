package common

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type HandlerDependencies struct {
	DB       *sql.DB
	StrDummy string
	IntDummy int
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func WriteErrorJSONResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	response := ErrorResponse{
		ErrorMessage: errorMessage,
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
