package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := data.GetAllCategories();
	json, _ := json.Marshal(categories);
	w.Write(json);
}