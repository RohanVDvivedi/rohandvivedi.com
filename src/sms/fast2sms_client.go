package sms

import (
	"fmt"
	"io/ioutil"
	//"net/url"
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
		Sender_id	string 		`json:"sender_id"`
		Message		string 		`json:"message"`
		Language	string  	`json:"language"`
		Route		string  	`json:"route"`
		Numbers		[]string  	`json:"numbers"`
		Flash 		string  	`json:"flash"`
	};

	reqP := requestPayload {
		Sender_id: "FSTSMS",
		"37740",
		loginCode
		Message: message,
		Language: "english",
		Route: "p",
		Numbers: dest,
		Flash: "1",
	}

	fmt.Println(reqP)

	client := &http.Client{}
	reqBody, err := json.Marshal(&reqP)/*"sender_id=" + reqP.sender_id + "&" +
				"message=" + strings.ReplaceAll(url.QueryEscape(reqP.message), "+", "%20") + "&" +
				"language=" + reqP.language + "&" +
				"route=" + reqP.route + "&" +
				"numbers=" + strings.Join(reqP.numbers, ",") + "&" +
				"flash=" + reqP.flash*/
	if(err != nil) {
		fmt.Println(reqBody, err)
	}
	req, err := http.NewRequest("POST", "http://example.com", strings.NewReader(string(reqBody)))
	req.Header.Add("authorization", config.GetGlobalConfig().Fast2SMS_auth)
	req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Cache-Control", "no-cache")
	resp, err := client.Do(req)

	if(err != nil || resp.StatusCode != 200) {
		return false
	}
	respBody, errRead := ioutil.ReadAll(resp.Body)
	if(errRead != nil) {
		return false
	}

	type responsePayload struct {
		Success bool 		`json:"return"`
		Request_id string 	`json:"request_id"`
		Message []string 	`json:"message"`
	}

	respP := &responsePayload{}
	errUnmarshal := json.Unmarshal(respBody, respP);
	fmt.Println(string(reqBody))
	fmt.Println(resp)
	fmt.Println(string(respBody))
	fmt.Println(respP)
	if(errUnmarshal != nil || !respP.Success) {
		return false
	}

	return true
}