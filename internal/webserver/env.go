package webserver

import (
	"tcms-web-bridge/internal/dry"
)

func getApiHost() (string, error) {
	return dry.GetEnvStr("API_HOST")
}
