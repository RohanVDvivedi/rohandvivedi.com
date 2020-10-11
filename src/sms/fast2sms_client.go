package sms

import (
	"net/http"
	"rohandvivedi.com/src/config"
	"strings"
	"encoding/json"
)

const Fast2SMS_URL = "https://www.fast2sms.com/dev/bulk"

// returns true if sent
func Send(dest []string, isFlash bool, message string) bool {

	if(len(dest) == 0 || config.GetGlobalConfig().Fast2SMS_auth == "") {
		return false
	}

	type requestPayload struct {
		sender_id	string
		message		string
		language	string
		route		string
		numbers		[]string
		flash 		string
	};

	reqP := postSmsPayload {
		sender_id: "FSTSMS",
		message: message,
		language: "english",
		route: "t",
		numbers: dest,
		flash: "0"
	}
	if(isFlash) {
		reqP.flash = "1"
	}

	type responsePayload struct {
		success bool `json:"return"`
		request_id string
		message []string
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://example.com", 
			"sender_id=" + reqP.sender_id + "&" +
			"message=" + reqP.message + "&" +
			"language=" + reqP.language + "&" +
			"route=" + reqP.route + "&" +
			"numbers=" + strings.Join(reqP.numbers, ",") + "&" +
			"flash=" + reqP.flash)
	req.Header.Add("authorization", config.GetGlobalConfig().Fast2SMS_auth)
	resp, err := client.Do(req)

	if(err != nil || resp.StatusCode != 200) {
		return false
	}
	respBody, errRead := ioutil.ReadAll(resp.Body)
	if(errRead != nil) {
		return false
	}

	respP := &responsePayload{}
	errUnmarshal := json.Unmarshal(respBody, respP);
	if(errUnmarshal != nil || !respP.success) {
		return false
	}

	fmt.Println(reqBody)
	fmt.Println(respP)

	return true
}