package config

import (
	"fmt"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

type MyUrlString string

type Specification struct {
	Port         int
	DB_url       MyUrlString
	Jaeger_url   MyUrlString
	Sentry_url   MyUrlString
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
	Username     string
}

func (adr *MyUrlString) Decode(value string) error {
	_, err := url.Parse(value)
	if err != nil {
		fmt.Errorf("Ошибка в формате адреса %w", err)
	}
	*adr = MyUrlString(value)
	return nil
}

func ReadConf() (spec Specification) {

	err := envconfig.Process("", &spec)
	if err != nil {
		fmt.Errorf("Ошибка чтения конфигурации %w", err)
	}
	return spec
}
