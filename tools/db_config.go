package tools

import "fmt"

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  bool
	TimeZone string
}

func (dbc DBConfig) String() string {
	if dbc.Host == "" {
		dbc.Host = "localhost"
	}

	if dbc.Port == 0 {
		dbc.Port = 5432
	}

	var sslmode string
	if dbc.SSLMode {
		sslmode = "enable"
	} else {
		sslmode = "disable"
	}

	if dbc.TimeZone == "" {
		dbc.TimeZone = "Asia/Tehran"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbc.Host,
		dbc.User,
		dbc.Password,
		dbc.DBName,
		dbc.Port,
		sslmode,
		dbc.TimeZone,
	)
}
