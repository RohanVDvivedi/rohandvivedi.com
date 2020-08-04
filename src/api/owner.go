package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
)
 
func GetOwner(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(*data.GetOwner());
	w.Write(json);
}