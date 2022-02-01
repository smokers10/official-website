package libraries

import (
	"encoding/json"
	"io/ioutil"
)

type configuration struct {
	Database    database    `json:"database"`
	Application application `json:"application"`
	SMTP        smtp        `json:"smtp"`
}

type database struct {
	Host                      string `json:"mysql_host"`
	Port                      int    `json:"mysql_port"`
	Username                  string `json:"mysql_username"`
	Database                  string `json:"mysql_database"`
	Password                  string `json:"mysql_password"`
	SessionTable              string `json:"mysql_session_table"`
	URI                       string `json:"mysql_uri"`
	MaxConnection             int    `json:"mysql_max_connection"`
	MaxConnectionLifeTime     int    `json:"mysql_max_connection_life_time"`
	MaxIdleConnection         int    `json:"mysql_max_idle_connection"`
	MaxIdleConnectionLifeTime int    `json:"mysql_max_idle_connection_life_time"`
}

type application struct {
	Secret         string `json:"app_secret"`
	Port           string `json:"app_port"`
	ProductionMode string `json:"app_production_mode"`
	BaseURL        string `json:"app_base_url"`
}

type smtp struct {
	Host         string `json:"SMTP_host"`
	Port         string `json:"SMTP_port"`
	SenderName   string `json:"SMTP_sender_name"`
	AuthUsername string `json:"SMTP_auth_username"`
	AuthPassword string `json:"SMTP_auth_password"`
}

func ReadConfig() *configuration {
	result := configuration{}

	// on unit testing change file dir to ../../app.config.json
	raw, err := ioutil.ReadFile("app.config.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(raw, &result)

	return &result
}
