package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/stat"
)

// api handlers in this file
var GetServerSystemStats = http.HandlerFunc(getServerSystemStats)

func getServerSystemStats(w http.ResponseWriter, r *http.Request) {
	SysStat := stat.GetServerSystemStats();
	json, _ := json.Marshal(SysStat);
	w.Write(json);
}