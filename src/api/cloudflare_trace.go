package api

import (
	"net/http"
	"rohandvivedi.com/src/useractlogger"
	"rohandvivedi.com/src/session"
	"io/ioutil"
	"strings"
)

// api handlers in this file
var CloudflareTrace = http.HandlerFunc(cloudflareTrace)

func cloudflareTrace(w http.ResponseWriter, r *http.Request) {
	traceData, traceDataError := ioutil.ReadAll(r.Body)
	if(traceDataError == nil || len(traceData) == 0) {
		traceString := strings.Join(strings.Fields(strings.TrimSpace(string(traceData))), " ")

		// store the cloudflare trace in the session of the user
		s := session.GlobalSessionStore.GetExistingSession(r);
        sessionId := "UNKNOWN_OR_NEW_USER"
        if(s != nil) {
        	s.SetValue("CLOUDFLARE_TRACE", traceString)
            sessionId = s.SessionId
        }
        // also log the trace to user activity log
        go useractlogger.LogUserActivity(sessionId, r.URL.Path, traceString);
	}
}