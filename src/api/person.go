package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
	"strconv"
	"strings"
)

type Person struct {
	data.Person
	Socials []data.Social
}
 
func GetOwner(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(*data.GetOwner());
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
		requested_socials, exists_get_socials := parameters["get_socials"];
		if exists_get_socials {
			if(requested_socials[0] == "true") {
				p.Socials = p_db.FindSocials();
			}
		}
	}

	json, _ := json.Marshal(*p);
	w.Write(json);
}