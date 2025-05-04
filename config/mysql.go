package config

import (
	"time"

	"github.com/SarkiMudboy/meeet/pkg/env"
)

type DBConfig struct {
	addr            string
	maxIdleConn     int
	maxOpenConn     int
	maxConnLifetime int
}

func loadDBConfig() *DBConfig {
	return &DBConfig{
		addr:            env.GetString("DB_ADDR", "admin:1234@/admin?parseTime=true"),
		maxIdleConn:     env.GetInt("DB_MAX_IDLE_CONN", 10),
		maxOpenConn:     env.GetInt("DB_MAX_OPEN_CONN", 10),
		maxConnLifetime: env.GetInt("DB_MAX_CONN_LIFETIME", 10),
	}

}

func (d *DBConfig) Addr() string {
	return d.addr
}
func (d *DBConfig) MaxIdleConn() int {
	return d.maxIdleConn
}
func (d *DBConfig) MaxOpenConn() int {
	return d.maxOpenConn
}
func (d *DBConfig) MaxConnLifetime() time.Duration {
	return time.Duration(d.maxConnLifetime)
}
