package db

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewConfigDB(host string, port uint16, user, password, dbname string, sslMode bool) ConfigDB {
	return ConfigDB{host, port, user, password, dbname, sslMode}
}

type ConfigDB struct {
	host     string
	port     uint16
	user     string
	password string
	dbname   string
	sslMode  bool
}

func (conf *ConfigDB) GetConfigString() string {
	var sslMode string
	if conf.sslMode {
		sslMode = "enable"
	} else {
		sslMode = "disable"
	}
	configString := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=%s",
		conf.user,
		conf.password,
		conf.host,
		conf.port,
		conf.dbname,
		sslMode,
	)
	return configString
}

// Connect creates
func Connect(conf ConfigDB) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", conf.GetConfigString())
	if err != nil {
		return nil, err
	}
	return db, nil
}
