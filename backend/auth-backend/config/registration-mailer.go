package config

import (
	"credential-auth/helper"
)

var RegistrationMailer = struct {
	VerificationUrl string
	SenderName      string
	ExpirationTime  string
}{
	VerificationUrl: helper.Env.Get("VERIFICATION_URL", ""),
	SenderName:      helper.Env.Get("VERIFICATION_SENDER_NAME", ""),
	ExpirationTime:  helper.Env.Get("VERIFICATION_EXPIRATION_TIME", ""),
}
