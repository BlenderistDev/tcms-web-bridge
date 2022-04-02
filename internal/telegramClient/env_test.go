package telegramClient

import (
	"testing"

	"tcms-web-bridge/internal/dry"
)

func TestGetTelegramBridgeHost(t *testing.T) {
	dry.TestEnvString(t, "TELEGRAM_BRIDGE_HOST", getTelegramBridgeHost)
}
