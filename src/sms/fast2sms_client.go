package sms

import (
	"io/ioutil"
	"net/http"
	"rohandvivedi.com/src/config"
	"strings"
	"encoding/json"
)

const Fast2SMS_URL = "https://www.fast2sms.com/dev/bulk"

// returns true if sent
func SendLoginCode(dest string, loginCode string) bool {

	if(len(dest) == 0 || config.GetGlobalConfig().Fast2SMS_auth == "") {
		return false
	}

	type requestPayload struct {
		Sender_id			string 		`json:"sender_id"`
		Language			string  	`json:"language"`
		Route				string  	`json:"route"`
		Numbers				[]string  	`json:"numbers"`
		Flash 				string  	`json:"flash"`
		Message				string 		`json:"message"`
		Variables			string 		`json:"variables"`
		Variables_values	string 		`json:"variables_values"`
	};

	reqP := requestPayload {
		Sender_id: "FSTSMS",
		Language: "english",
		Route: "qt",
		Numbers: []string{dest},
		Flash: "1",
		Message: "37740",
		Variables: "{#FF#}",
		Variables_values: loginCode,
	}

	client := &http.Client{}
	reqBody := 	"sender_id=" + reqP.Sender_id +
				"&language=" + reqP.Language +
				"&route=" + reqP.Route +
				"&numbers=" + strings.Join(reqP.Numbers, ",") +
				"&flash=" + reqP.Flash +
				"&message=" + reqP.Message +
				"&variables=" + reqP.Variables +
				"&variables_values=" + reqP.Variables_values
	req, err := http.NewRequest("POST", Fast2SMS_URL, strings.NewReader(string(reqBody)))
	req.Header.Add("authorization", config.GetGlobalConfig().Fast2SMS_auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Cache-Control", "no-cache")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if(err != nil || resp.StatusCode != 200) {
		return false
	}

	respBody, errRead := ioutil.ReadAll(resp.Body)
	if(errRead != nil) {
		return false
	}

	type responsePayload struct {
		Success bool 		`json:"return"`
		Status_code string 	`json:"Status_code"`
		Request_id string 	`json:"request_id"`
		Message []string 	`json:"message"`
	}

	respP := &responsePayload{}
	errUnmarshal := json.Unmarshal(respBody, respP);
	if(errUnmarshal != nil || !respP.Success) {
		return false
	}

	return true
}