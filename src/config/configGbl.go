package config

type Config struct {
	ssl_enabled									string
	create_user_sessions						string
	send_deployment_mail_to_owner				string
	send_server_monitor_status_mail_to_owner	string
}

var configGlobal *Config = nil;

func InitGlobalConfig(configFile string) {
	configGlobal = &Config{};
}

func getGlobalConfig() *Config {
	return configGlobal;
}