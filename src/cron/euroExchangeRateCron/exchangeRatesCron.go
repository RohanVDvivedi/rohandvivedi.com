package euroExchangeRateCron

import (
	"fmt"
	"strings"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"github.com/robfig/cron"
)

import (
	"rohandvivedi.com/src/mailManager"
)

func AttachCron() {
	//sendEuroExchangeRatesMail()
	c := cron.New()
    c.AddFunc(sendMailCron, sendEuroExchangeRatesMail)
    c.Start()
}

var sendMailCron = "0 * ? * ?";

func sendEuroExchangeRatesMail() {
	mail_to := []string{"rohandvivedi@gmail.com"};
	subject := "EUR --> INR Exchange rates";
	mail := "";

	currRates := getLatestExchangeRates("EUR", []string{"INR"});

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
	Err error
	Base string
	Date time.Time
	Rates map[string]float64
};

/*func (r *CurrencyRates) UnmarshalJSON(b []byte) error {
    return json.Unmarshal(b, r);
}

func (r CurrencyRates) MarshalJSON() ([]byte, error) {
    return json.Marshal(r);
}*/

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
		rates.Err = err
		return &rates;
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rates.Err = err
		return &rates;
	}
	err = json.Unmarshal(body, &rates);
	if(err != nil) {
		rates.Err = err
		return &rates;
	}
	return &rates;
}