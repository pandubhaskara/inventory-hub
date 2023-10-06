package config

import (
	"account-backend/helper"
)

var Amqp = struct {
	Host               string
	Port               string
	User               string
	Password           string
	Exchange           string
	RoutingKeyRegister string
	QueueRegister      string
}{
	Host:               helper.Env.Get("AMQP_HOST", ""),
	Port:               helper.Env.Get("AMQP_PORT", ""),
	User:               helper.Env.Get("AMQP_USERNAME", ""),
	Password:           helper.Env.Get("AMQP_PASSWORD", ""),
	Exchange:           helper.Env.Get("AMQP_EXCHANGE", ""),
	RoutingKeyRegister: helper.Env.Get("AMQP_ROUTING_KEY_REGISTER", ""),
	QueueRegister:      helper.Env.Get("AMQP_QUEUE_REGISTER", ""),
}
