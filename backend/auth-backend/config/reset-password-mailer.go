package config

import (
	"credential-auth/helper"
)

var ResetPasswordMailer = struct {
	ResetPasswordUrl string
	SenderName       string
	ExpirationTime   string
	RebecaSenderName string
}{
	ResetPasswordUrl: helper.Env.Get("RESET_PASSWORD_URL", ""),
	SenderName:       helper.Env.Get("RESET_PASSWORD_SENDER_NAME", ""),
	ExpirationTime:   helper.Env.Get("RESET_PASSWORD_EXPIRATION_TIME", ""),
	RebecaSenderName: helper.Env.Get("REBECA_EMAIL_SENDER_NAME", ""),
}
