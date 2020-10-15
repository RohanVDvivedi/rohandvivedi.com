package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/useractlogger"
	"rohandvivedi.com/src/session"
	"io/ioutil"
	"strings"
)

// api handlers in this file
var CloudflareTrace = http.HandlerFunc(cloudflareTrace)

func cloudflareTrace(w http.ResponseWriter, r *http.Request) {
	traceString, traceStringError := r.Body
	if(traceStringError == nil) {
		traceString = strings.Join(strings.Fields(strings.TrimSpace(traceString)), " ")

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