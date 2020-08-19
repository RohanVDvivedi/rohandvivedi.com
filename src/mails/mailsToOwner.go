package mails

import (
	"rohandvivedi.com/src/config"
	"rohandvivedi.com/src/mailManager"
)

func SendDeploymentMail() {
	if(config.GetGlobalConfig().Auth_mail_client && config.GetGlobalConfig().Send_deployment_mail) {

		mailBody := "Deplyment Successfull"
		
		msg := mailManager.WritePlainEmail([]string{config.GetGlobalConfig().From_mailid}, "Deployment Mail", mailBody);
		mailManager.SendMail([]string{config.GetGlobalConfig().From_mailid}, msg)
	}
}