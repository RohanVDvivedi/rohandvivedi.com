package mails

import (
	"encoding/json"
)

import (
	"rohandvivedi.com/src/config"
	"rohandvivedi.com/src/mailManager"
)

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