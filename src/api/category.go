package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
)

// api handlers in this file
var GetAllCategories = http.HandlerFunc(getAllCategories)

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := data.GetAllCategories();
	json, _ := json.Marshal(categories);
	w.Write(json);
}