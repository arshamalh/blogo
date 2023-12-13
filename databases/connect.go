// connect.go
package databases

import (
	"os"

	database "github.com/arshamalh/blogo/databases/bun"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/tools"
)

func ConnectDB() (Database, error) {
	var db Database

	if os.Getenv("USE_BUN") == "true" {
		// Use BunDB
		dsn := tools.DBConfig{
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASS"),
			DBName:   os.Getenv("DB_NAME"),
			Host:     os.Getenv("HOST"),
		}
		bunDB, err := database.Connect(dsn.String())
		if err != nil {
			log.Gl.Error(err.Error())
			return nil, err
		}
		db = bunDB
	} else {
		// Use GormDB (default)
		dsn := tools.DBConfig{
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASS"),
			DBName:   os.Getenv("DB_NAME"),
			Host:     os.Getenv("HOST"),
		}
		gormDB, err := database.Connect(dsn.String())
		if err != nil {
			log.Gl.Error(err.Error())
			return nil, err
		}
		db = gormDB
	}

	log.Gl.Info("Database connection successfully opened")

	return db, nil
}
