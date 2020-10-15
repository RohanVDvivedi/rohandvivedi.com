package useractlogger

import (
	"os"
	"log"
	"rohandvivedi.com/src/config"
)

var llog *log.Logger = nil

func init() {
	file, err := os.OpenFile("./log/user_activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	llog = log.New(file, "", log.Ldate|log.Ltime)
}

func LogUserActivity(sId string, path string, data string) bool {
	if(config.GetGlobalConfig().Enable_user_activity_logging) {
		llog.Printf("%s %s %s\n", sId, path, data);
		return true
	}
	return false
}