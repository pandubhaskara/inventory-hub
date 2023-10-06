package config

import (
	"account-backend/helper"
)

var Aws = struct {
	Bucket string
}{
	Bucket: helper.Env.Get("AWS_BUCKET", ""),
}
