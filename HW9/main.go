package main

import (
	"fmt"
	"my_mod/conf"
	"my_mod/data"
	"os"
)

func main() {

	os.Setenv("PORT", "8080")
	os.Setenv("DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	os.Setenv("JAEGER_URL", "http://jaeger:16686")
	os.Setenv("SENTRY_URL", "http://sentry:9000")
	os.Setenv("KAFKA_BROKER", "kafka:9092")
	os.Setenv("SOME_APP_ID", "testid")
	os.Setenv("SOME_APP_KEY", "testkey")

	spec := config.ReadConf()

	format := "Port: %d\nDB_url: %s\nJaeger_url: %s\nSentry_url: %s\nKafka_broker: %s\nSome_app_id: %s\nSome_app_key: %s\n"
	_, err := fmt.Printf(format, spec.Port, spec.DB_url, spec.Jaeger_url, spec.Sentry_url, spec.Kafka_broker, spec.Some_app_id, spec.Some_app_key)
	if err != nil {
		fmt.Errorf("ошибка печати %w", err)
	}

	userData := data.ReadUserData("./data/data.json")

	fmt.Printf("Данные пользователя: %+v\n", userData)

	spYaml := config.ReadYaml("./conf/config.yaml")

	fmt.Printf("конфигурация Yaml: %+v\n", spYaml)

	JsonData := config.ReadJsonData("./conf/config.json")

	fmt.Printf("конфигурация JSON: %+v\n", JsonData)
}
