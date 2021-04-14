package sysCron

import (
	"github.com/robfig/cron/v3"
)

import (
	"rohandvivedi.com/src/mails"
	"rohandvivedi.com/src/config"
)

var c *cron.Cron = nil;

func InitializeSystemCron() {
	if(config.GetGlobalConfig().Enable_all_cron) {
		c = cron.New()
		c.AddFunc("CRON_TZ=Asia/Kolkata 40 20 * * *", func(){mails.SendServerSystemStatsMail()})
		c.Start()
	}
}

func DeinitializeSystemCron() {
	if(config.GetGlobalConfig().Enable_all_cron) {
		c.Stop()
	}
}