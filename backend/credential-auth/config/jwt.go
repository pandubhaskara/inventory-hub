package config

import (
	"credential-auth/helper"
)

var Jwt = struct {
	ATDuration string
	RTDuration string
	Signature  string
	Issuer     string
}{
	ATDuration: helper.Env.Get("ACCESS_TOKEN_DURATION", ""),
	RTDuration: helper.Env.Get("REFRESH_TOKEN_DURATION", ""),
	Signature:  helper.Env.Get("JWT_SIGNATURE_KEY", ""),
	Issuer:     helper.Env.Get("JWT_ISSUER", ""),
}
