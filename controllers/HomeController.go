package controllers

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(map[string]string{
		"message": "Working Successfully",
	})
}
