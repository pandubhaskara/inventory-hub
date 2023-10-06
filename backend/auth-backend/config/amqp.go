package config

import (
	"credential-auth/helper"
)

var Amqp = struct {
	Host           string
	Port           string
	User           string
	Password       string
	Exchange       string
	RoutingKey     string
	ExchangeMail   string
	RoutingKeyMail string
}{
	Host:           helper.Env.Get("AMQP_HOST", ""),
	Port:           helper.Env.Get("AMQP_PORT", ""),
	User:           helper.Env.Get("AMQP_USERNAME", ""),
	Password:       helper.Env.Get("AMQP_PASSWORD", ""),
	Exchange:       helper.Env.Get("AMQP_EXCHANGE", ""),
	RoutingKey:     helper.Env.Get("AMQP_ROUTING_KEY", ""),
	ExchangeMail:   helper.Env.Get("AMQP_EXCHANGE_MAIL", ""),
	RoutingKeyMail: helper.Env.Get("AMQP_ROUTING_KEY_MAIL", ""),
}
