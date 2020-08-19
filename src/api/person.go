package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
	"strconv"
	"strings"
)

// api handlers in this file
var GetOwner = http.HandlerFunc(getOwner)
var GetPerson = http.HandlerFunc(getPerson)

type Person struct {
	data.Person
	Socials []data.Social
	Pasts []data.Past
}

func getOwner(w http.ResponseWriter, r *http.Request) {
	var p *Person = nil;

	p_db := data.GetOwner()

	if(p_db != nil) {
		p = &Person{};
		p.Person = *p_db

		// check if you need to send them socials
		requested_socials, exists_get_socials := r.URL.Query()["get_socials"];
		if exists_get_socials && (requested_socials[0] == "true") {
			p.Socials = p_db.FindSocials();
		}

		// check if you need to send them pasts
		requested_pasts, exists_get_pasts := r.URL.Query()["get_pasts"];
		if exists_get_pasts && (requested_pasts[0] == "true") {
			p.Pasts = p_db.FindPasts();
		}
	}

	json, _ := json.Marshal(p);
	w.Write(json);
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	var p *Person = nil;
	parameters := r.URL.Query();

	_, exists_id := parameters["id"];
	_, exists_name := parameters["name"];

	var p_db *data.Person = nil;

	if exists_id {
		id, err := strconv.Atoi(parameters["id"][0])
		if err == nil {
			p_db = data.GetPersonById(id);
		}
	}

	if p_db == nil && exists_name {
		fullNameFields := strings.Fields(parameters["name"][0])
		p_db = data.GetPersonByName(fullNameFields[0], fullNameFields[1]);
	}

	if(p_db != nil) {
		p = &Person{};
		p.Person = *p_db

		// check if you need to send them socials
		requested_socials, exists_get_socials := parameters["get_socials"];
		if exists_get_socials {
			if(requested_socials[0] == "true") {
				p.Socials = p_db.FindSocials();
			}
		}

		// check if you need to send them pasts
		requested_pasts, exists_get_pasts := r.URL.Query()["get_pasts"];
		if exists_get_pasts && (requested_pasts[0] == "true") {
			p.Pasts = p_db.FindPasts();
		}
	}

	json, _ := json.Marshal(p);
	w.Write(json);
}