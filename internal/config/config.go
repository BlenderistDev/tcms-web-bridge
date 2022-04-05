package config

import "github.com/spf13/viper"

type Config struct {
	ApiHost            string `mapstructure:"API_HOST"`
	TelegramBridgeHost string `mapstructure:"TELEGRAM_BRIDGE_HOST"`
	TcmsHost           string `mapstructure:"TCMS_HOST"`
	KafkaHost          string `mapstructure:"KAFKA_HOST"`
	KafkaTopic         string `mapstructure:"KAFKA_TOPIC"`
	KafkaGroupId       string `mapstructure:"KAFKA_GROUP_ID"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
