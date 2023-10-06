package config

import (
	"credential-auth/helper"
)

var Database = struct {
	Connection string
	Host       string
	Port       string
	Name       string
	User       string
	Password   string
}{
	Connection: helper.Env.Get("DB_CONNECTION", ""),
	Host:       helper.Env.Get("DB_HOST", ""),
	Port:       helper.Env.Get("DB_PORT", ""),
	Name:       helper.Env.Get("DB_DATABASE", ""),
	User:       helper.Env.Get("DB_USERNAME", ""),
	Password:   helper.Env.Get("DB_PASSWORD", ""),
}
