package mails

import (
	"net/http"
	"encoding/json"
	"strconv"
)

import (
	"rohandvivedi.com/src/config"
	"rohandvivedi.com/src/mailManager"
	"rohandvivedi.com/src/session"
)

// api handlers in this file
var SendAnonymousMail = http.HandlerFunc(sendAnonymousMail)

func SendDeploymentMail(ownerSessionId string) {
	if(config.GetGlobalConfig().Auth_mail_client && config.GetGlobalConfig().Send_deployment_mail) {

		temp := config.GetGlobalConfig()
		temp.From_password = "********"

		jsonConfig, _ := json.MarshalIndent(temp, "", "    ")
		mailBody := "Deployment Successfull\nconfig:\n" + string(jsonConfig)
		mailBody += "\n\nSession Id : " + ownerSessionId + "\n"
		mailBody += "\nrohandvivedi.com\n"
		
		msg := mailManager.WritePlainEmail([]string{config.GetGlobalConfig().From_mailid}, "Deployment Mail", mailBody);
		mailManager.SendMail([]string{config.GetGlobalConfig().From_mailid}, msg)
	}
}

func sendAnonymousMail(w http.ResponseWriter, r *http.Request) {

	if(!config.GetGlobalConfig().Auth_mail_client || !config.GetGlobalConfig().Create_user_sessions) {
		w.Write([]byte("{'status':'failure','reason':'missing mail auth client or session store'}"))
		return;
	}

	s := session.GlobalSessionStore.GetOrCreateSession(w, r);

	userSessionId := ""
	userAnonMailCount := 1

	if(s != nil) {
		userSessionId = s.SessionId
		anonMailCountKey := "anon_mail_count"
		anonMailCount, anonMailCountExists := s.GetValue(anonMailCountKey);
		if(anonMailCountExists) {
			valanonMailCount, ok := anonMailCount.(int)
			if(ok){
				if(valanonMailCount >= 3) {
					w.Write([]byte("{'status':'failure','reason':'anonymous mail request limit reached'}"))
					return;
				}
				userAnonMailCount = valanonMailCount + 1
				s.SetValue(anonMailCountKey, userAnonMailCount);
			} else {
				s.SetValue(anonMailCountKey, 1);
			}
		} else {
			s.SetValue(anonMailCountKey, 1);
		}
	}

	subjects, exists_subjects := r.URL.Query()["anon_mail_subject"];
	subject := ""
	if(exists_subjects) {
		subject = subjects[0]
	}

	bodies, exists_bodies := r.URL.Query()["anon_mail_body"];
	body := ""
	if(exists_bodies) {
		body = bodies[0]
	}

	msg := mailManager.WritePlainEmail([]string{config.GetGlobalConfig().From_mailid}, 
	"Anonymous User \"" + userSessionId + "\" : " + subject + " -> " + strconv.Itoa(userAnonMailCount), body);
	mailManager.SendMail([]string{config.GetGlobalConfig().From_mailid}, msg)

	w.Write([]byte("{'status':'anonymous mail sent'}"))
}