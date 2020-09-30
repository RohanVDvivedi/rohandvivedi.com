package mails

import (
	"net/http"
	"encoding/json"
	"strconv"
	"time"
)

import (
	"rohandvivedi.com/src/config"
	"rohandvivedi.com/src/mailManager"
	"rohandvivedi.com/src/session"
	"rohandvivedi.com/src/stat"
)

// api handlers in this file
var SendAnonymousMail = http.HandlerFunc(sendAnonymousMail)

func SendDeploymentMail() bool {
	if(config.GetGlobalConfig().Auth_mail_client && config.GetGlobalConfig().Send_deployment_mail) {

		temp := config.GetGlobalConfig()
		temp.From_password = "********"

		jsonConfig, _ := json.MarshalIndent(temp, "", "    ")
		mailBody := "Deployment Successfull\nconfig:\n" + string(jsonConfig)
		mailBody += "\n\nrohandvivedi.com\n"
		
		msg := mailManager.WritePlainEmail([]string{config.GetGlobalConfig().From_mailid}, "Deployment Mail", mailBody);
		mailManager.SendMail([]string{config.GetGlobalConfig().From_mailid}, msg)

		return true
	}
	return false
}

func SendLoginCodeMail(loginCode string) bool {
	if(config.GetGlobalConfig().Auth_mail_client) {

		msg := mailManager.WritePlainEmail([]string{config.GetGlobalConfig().From_mailid},
			"Login code as requested", "\nYour owner login code : " + loginCode + "\n\nrohandvivedi.com\n");
		mailManager.SendMail([]string{config.GetGlobalConfig().From_mailid}, msg)

		return true
	}
	return false
}

func SendServerSystemStatsMail() {
	if(config.GetGlobalConfig().Auth_mail_client && config.GetGlobalConfig().Send_server_status_mail) {

		json, _ := json.MarshalIndent(stat.GetServerSystemStats(), "", "    ")
		mailBody := "Server System Stats:\n" + string(json)
		
		msg := mailManager.WritePlainEmail([]string{config.GetGlobalConfig().From_mailid}, "Server System Stats Mail", mailBody);
		mailManager.SendMail([]string{config.GetGlobalConfig().From_mailid}, msg)
	}
}

func sendAnonymousMail(w http.ResponseWriter, r *http.Request) {

	type RequestPayload struct {
		Subject 	string
		Body 		string
	}
	type ResponsePayload struct {
		Success 	bool
		Message 	string
	}

	if(!config.GetGlobalConfig().Auth_mail_client) {
		json, _ := json.Marshal(ResponsePayload{false,"services offline"});
		w.Write(json)
		return;
	}

	pld := RequestPayload{};

	err := json.NewDecoder(r.Body).Decode(&pld);
	if( err!=nil || len(pld.Subject) == 0 || len(pld.Body) == 0) {
		json, _ := json.Marshal(ResponsePayload{false,"Subject or Body contains inappropriate fields or are empty, please try again"});
		w.Write(json)
		return
	}

	s := session.GlobalSessionStore.GetExistingSession(r);
	userSessionId := s.SessionId

	userAnonMailCountIntr := s.ExecuteOnValues(func (values map[string]interface{}, additional_params interface{}) interface{} {
		userAnonMailCount := 0
		anonMailCountKey := "anon_mail_count"
		anonMailLastTimeKey := "anon_mail_last_sent"

		anonMailCount, anonMailCountExists := values[anonMailCountKey];
		if(anonMailCountExists) {
			valAnonMailCount, ok := anonMailCount.(int)
			if(ok){
				if(valAnonMailCount>0 && (valAnonMailCount%3)==0) {
					anonMailLastTimeIntr, lastTimeExists := values[anonMailLastTimeKey]
					if(!lastTimeExists){
						return nil
					}
					anonMailLastTime, isTime := anonMailLastTimeIntr.(time.Time)
					if(!isTime || time.Now().Sub(anonMailLastTime) < time.Hour * 48){
						return nil;
					}
				}
				userAnonMailCount = valAnonMailCount
			}
		}

		userAnonMailCount = userAnonMailCount + 1
		values[anonMailLastTimeKey] = time.Now()
		values[anonMailCountKey] = userAnonMailCount
		return userAnonMailCount
	}, nil);

	userAnonMailCount, ok := userAnonMailCountIntr.(int)

	if(userAnonMailCountIntr == nil || !ok){
		json, _ := json.Marshal(ResponsePayload{false,"anonymous mail request limit reached, please wait 48 hours and try again"});
		w.Write(json)
		return
	}

	msg := mailManager.WritePlainEmail([]string{config.GetGlobalConfig().From_mailid}, 
	"Anonymous User \"" + userSessionId + "\" : " + pld.Subject + " -> " + strconv.Itoa(userAnonMailCount), pld.Body);
	mailManager.SendMail([]string{config.GetGlobalConfig().From_mailid}, msg)

	json, _ := json.Marshal(ResponsePayload{true,"Thank you, for contacting me."});
	w.Write(json)
}