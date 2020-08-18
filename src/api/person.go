package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
	"strconv"
	"strings"
)

import (
	"rohandvivedi.com/src/session"
)

type Person struct {
	data.Person
	Socials []data.Social
	Pasts []data.Past
}

func CountApiHitsInSessionValues(SessionVals map[string]interface{}, add_par interface{}) interface{} {
	count, exists := SessionVals["GetOwner_count"]
	if(exists) {
		if int_count, ok := count.(int); ok {
			SessionVals["GetOwner_count"] = ((int)(int_count)) + 1
			return nil
		}
	}
	SessionVals["GetOwner_count"] = 1
	return nil
}
 
func GetOwner(w http.ResponseWriter, r *http.Request) {

	s := session.GetOrCreateSession(w, r);
	if(s!=nil) {
		_ = s.ExecuteOnValues(CountApiHitsInSessionValues, nil);
	}
	session.PrintSessionStore()


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

func GetPerson(w http.ResponseWriter, r *http.Request) {
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