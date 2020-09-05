package config

import (
	"strings"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Environment				string

	SSL_enabled				bool
	Create_user_sessions	bool

	Auth_mail_client 		bool

	Send_deployment_mail	bool
	Send_server_status_mail	bool

	Recreate_search_index	bool

	From_mailid 			string
	From_password 			string

	SQLite3_database_file	string
	Bleve_search_index_file	string
}

var configGlobal *Config = nil;

func InitGlobalConfig(configFileToUse string) {
	configGlobal = &Config{};
	data := []byte{}
	if(strings.Contains(configFileToUse, "prod")) {
		data, _ = ioutil.ReadFile("./prod_config.json")
		_ = json.Unmarshal(data, &configGlobal);
		configGlobal.Environment = "prod"
	} else {
		data, _ = ioutil.ReadFile("./dev_config.json")
		_ = json.Unmarshal(data, &configGlobal);
		configGlobal.Environment = "dev"
	}
}

func GetGlobalConfig() Config {
	return *configGlobal;
}