package euroExchangeRateCron

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/robfig/cron"
)

import (
	"rohandvivedi.com/src/mailManager"
)

func AttachCron() {
	sendEuroExchangeRatesMail()
	c := cron.New()
	c.AddFunc(sendMailCron, sendEuroExchangeRatesMail)
	c.Start()
}

var sendMailCron = "*/20 * * * *";

func sendEuroExchangeRatesMail() {
	mail_to := []string{"rohandvivedi@gmail.com"};
	subject := "EUR --> INR Exchange rates";
	mail := "";

	currRates := getLatestEUR_INR_USDExchangeRates();

	json, err := json.MarshalIndent(currRates, "", "   ");
	if(err == nil) {
		mail = mailManager.WritePlainEmail(mail_to, subject, (string)(json));
	} else {
		mail = mailManager.WritePlainEmail(mail_to, subject, "{}");
	}

	err = mailManager.SendMail(mail_to, subject, mail);
	if(err != nil) {
		fmt.Println(err.Error())
	}
}

type CurrencyRates struct {
	Rates map[string]float64
	Err error
};

func getLatestEUR_INR_USDExchangeRates() (*CurrencyRates) {
	url := "https://free.currconv.com/api/v7/convert?q=EUR_INR,EUR_USD,USD_INR,USD_EUR,INR_USD,INR_EUR&compact=ultra&apiKey=752975347b77fe583aa3"
	rates := CurrencyRates{};
	resp, err := http.Get(url);
	if(err != nil) {
		rates.Err = err
		return &rates;
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rates.Err = err
		return &rates;
	}
	err = json.Unmarshal(body, &(rates.Rates));
	if(err != nil) {
		rates.Err = err
		return &rates;
	}
	return &rates;
}