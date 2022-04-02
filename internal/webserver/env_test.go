package webserver

import (
	"testing"

	"tcms-web-bridge/internal/dry"
)

func TestGetApiHost(t *testing.T) {
	dry.TestEnvString(t, "API_HOST", getApiHost)
}
