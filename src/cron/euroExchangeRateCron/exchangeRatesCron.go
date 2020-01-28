package euroExchangeRateCron

import (

	"strings"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

import (
	"rohandvivedi.com/src/mailManager"
)

func attachCron() {
	
}

var sendMailCron = "0 * * * *";

func sendEuroExchangeRatesMail() {
	mail_to := []string{"rohandvivedi@gmail.com"};
	subject := "EUR --> INR Exchange rates";
	mail := "";

	currRates := getLatestExchangeRates("EUR", []string{"INR"});

	json, err := json.Marshal(*currRates);
	if(err == nil) {
		mail = mailManager.WritePlainEmail(mail_to, subject, (string)(json));
	} else {
		mail = mailManager.WritePlainEmail(mail_to, subject, "{}");
	}

	mailManager.SendMail(mail_to, subject, mail);
}

type CurrencyRates struct {
	err error
	base string
	date time.Time
	rates map[string]float64
};

func getLatestExchangeRates(from string, to []string) (*CurrencyRates) {
	baseUrl := "https://api.exchangeratesapi.io/latest"
	url := baseUrl
	if(from != "" && len(to) > 0) {
		url += "?"
		if(from != "") {
			url += "base=" + from
			if(len(to) > 0) {
				url += "&"
			}
		}
		if(len(to) > 0) {
			url += "symbols=" + strings.Join(to[:], ",")
		}
	}
	rates := CurrencyRates{};
	resp, err := http.Get(url);
	if(err != nil) {
		rates.err = err
		return &rates;
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rates.err = err
		return &rates;
	}
	err = json.Unmarshal(body, &rates);
	if(err != nil){
		rates.err = err
		return &rates;
	}
	return &rates;
}