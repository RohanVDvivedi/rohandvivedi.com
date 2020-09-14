package useractlogger

import (
	"os"
	"log"
)

var llog *log.Logger = nil

func init() {
	file, err := os.OpenFile("./log/user_activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	llog = log.New(file, "", log.Ldate|log.Ltime)
}

func LogUserActivity(sId string, path string, data string) {
	llog.Printf("%s %s %s\n", sId, path, data);
}