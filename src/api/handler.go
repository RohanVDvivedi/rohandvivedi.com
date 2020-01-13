package api

import (
	"net/http"
	"encoding/json"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(struct {Name string; Skill string}{Name: "Rohan", Skill: "Coder"});
	if(err == nil) {
		w.Write(json);
	} else {
		w.Write([]byte("{}"));
	}
}