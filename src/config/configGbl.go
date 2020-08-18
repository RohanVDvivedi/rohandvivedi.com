package config

import (
	"strings"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	SSL_enabled				bool
	Create_user_sessions	bool
	Send_deployment_mail	bool
	Send_server_status_mail	bool
}

var configGlobal *Config = nil;

func InitGlobalConfig(configFileToUse string) {
	configGlobal = &Config{};
	data := []byte{}
	if(strings.Contains(configFileToUse, "prod")) {
		data, _ = ioutil.ReadFile("./prod_config.json")
	} else {
		data, _ = ioutil.ReadFile("./dev_config.json")
	}
	_ = json.Unmarshal(data, &configGlobal);
}

func GetGlobalConfig() Config {
	return *configGlobal;
}