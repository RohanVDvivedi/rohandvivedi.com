package main

import (
	"github.com/robfig/cron/v3"
)

import (
	"rohandvivedi.com/src/mails"
	"rohandvivedi.com/src/config"
)

var c *cron.Cron = nil;

func initializeSystemCron() {
	if(config.GetGlobalConfig().Enable_all_cron) {
		c = cron.New()
		c.AddFunc("CRON_TZ=Asia/Kolkata 40 20 * * *", mails.SendServerSystemStatsMail)
		c.Start()
	}
}

func deinitializeSystemCron() {
	if(config.GetGlobalConfig().Enable_all_cron) {
		c.Stop()
	}
}