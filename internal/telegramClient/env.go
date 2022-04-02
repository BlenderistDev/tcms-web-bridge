package telegramClient

import (
	"tcms-web-bridge/internal/dry"
)

func getTelegramBridgeHost() (string, error) {
	return dry.GetEnvStr("TELEGRAM_BRIDGE_HOST")
}
