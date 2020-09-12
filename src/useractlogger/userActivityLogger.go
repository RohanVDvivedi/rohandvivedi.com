package useractlogger

import (
	"os"
	"log"
	"sync"
	"rohandvivedi.com/src/config"
)

var onceInitLoggerOnce sync.Once
var llog *log.Logger = nil

func initUserActivityLogger() {
	file, err := os.OpenFile("./log/user_activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	llog = log.New(file, "", log.Ldate|log.Ltime)
}

func LogUserActivity(sId string, path string, data string) {
	// for user activity logging, you need 
	if(!config.GetGlobalConfig().Enable_user_activity_logging || 
		!config.GetGlobalConfig().Create_user_sessions) {
		return
	}

	onceInitLoggerOnce.Do(initUserActivityLogger)
	llog.Printf("%s %s %s\n", sId, path, data);
}