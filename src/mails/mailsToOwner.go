package mails

import (
	"encoding/json"
)

import (
	"rohandvivedi.com/src/config"
	"rohandvivedi.com/src/mailManager"
)

func SendDeploymentMail() {
	if(config.GetGlobalConfig().Auth_mail_client && config.GetGlobalConfig().Send_deployment_mail) {

		jsonConfig, _ := json.MarshalIndent(config.GetGlobalConfig(), "", "    ")
		mailBody := "Deployment Successfull\nconfig:\n" + string(jsonConfig)
		
		msg := mailManager.WritePlainEmail([]string{config.GetGlobalConfig().From_mailid}, "Deployment Mail", mailBody);
		mailManager.SendMail([]string{config.GetGlobalConfig().From_mailid}, msg)
	}
}