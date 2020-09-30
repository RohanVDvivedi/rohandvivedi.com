package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
	"rohandvivedi.com/src/session"
	"rohandvivedi.com/src/randstring"
	"rohandvivedi.com/src/config"
)

// api handlers in this file
var IsOwner = http.HandlerFunc(isOwner)
var ReqLoginOwnerCode = http.HandlerFunc(reqLoginOwnerCode)
var LoginOwner = http.HandlerFunc(loginOwner)
var LogoutOwner = http.HandlerFunc(logoutOwner)

var successTrueJson = json.Marshal(struct{Success bool}{true})
var successFalseJson = json.Marshal(struct{Success bool}{false})

func isOwner(w http.ResponseWriter, r *http.Request) {
	ownerIntr, hasOwner := session.GlobalSessionStore.GetExistingSession(r).GetValue("owner")
	owner, isBool := ownerIntr.(bool)
	if(hasOwner && isBool && owner) {
		w.Write(successTrueJson);
	} else {
		w.Write(successFalseJson);
	}
}

func reqLoginOwnerCode(w http.ResponseWriter, r *http.Request) {
	logInCodeCanBeSent := false
	loginCode, _ := session.GlobalSessionStore.GetExistingSession(r).ExecuteOnValues(function(values map[string]interface{}, add_params interface{}){
		ownerIntr, hasOwner := values["owner"]
		owner, isBool := ownerIntr.(bool)
		if(!(hasOwner && isBool && owner)) {
			values["owner"] = false
			values["owner_login_code"] = randstring.GetRandomString(6)
			logInCodeCanBeSent = true
		}
		return nil
	}, nil).(string)

	if(logInCodeCanBeSent) {
		loginCodeSent := false
		messageLoginCodeString := "Your owner login code : " + loginCode

		if(false) {
			loginCodeSent = true
		}

		if(false) {
			loginCodeSent = true
		}

		if(!loginCodeSent) {
			fmt.Println(messageLoginCodeString)
		}

		w.Write(successTrueJson);
	} else {
		w.Write(successFalseJson);
	}
}

func loginOwner(w http.ResponseWriter, r *http.Request) {
	loggedInSuccessfull := false
	loginCodesInRequest, exists_login_code := r.URL.Query()["login_code"];

	if(exists_login_code) {
		session.GlobalSessionStore.GetExistingSession(r).ExecuteOnValues(function(values map[string]interface{}, add_params interface{}){
			loginCodeIntr, loginCodeFound := values["owner_login_code"]
			logincode, isString := loginCodeIntr.(string)
			if(loginCodeFound && isString) {
				if(loginCode == loginCodesInRequest[0]) {
					values["owner"] = true
					delete(values, "owner_login_code")
					loggedInSuccessfull = true
				} else {
					values["owner"] = false
					delete(values, "owner_login_code")
				}
			}
			return nil
		}, nil)
	}

	if(loggedInSuccessfull) {
		w.Write(successTrueJson);
	} else {
		w.Write(successFalseJson);
	}
}

func logoutOwner(w http.ResponseWriter, r *http.Request) {
	loggedOutSuccessfull := false
	session.GlobalSessionStore.GetExistingSession(r).ExecuteOnValues(function(values map[string]interface{}, add_params interface{}){
		ownerIntr, hasOwner := values["owner"]
		owner, isBool := ownerIntr.(bool)
		if(hasOwner && isBool && owner) {
			values["owner"] = false
			delete(values, "owner_login_code")
			loggedOutSuccessfull = true
		}
		return nil
	}, nil)

	if(loggedOutSuccessfull) {
		w.Write(successTrueJson);
	} else {
		w.Write(successFalseJson);
	}
}