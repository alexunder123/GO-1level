package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type MyUrlString string

type Specification struct {
	Port         int         `json:"port"`
	DB_url       MyUrlString `json:"db_url"`
	Jaeger_url   MyUrlString `json:"jaeger_url"`
	Sentry_url   MyUrlString `json:"sentry_url"`
	Kafka_broker string      `json:"kafka_broker"`
	Some_app_id  string      `json:"some_app_id"`
	Some_app_key string      `json:"some_app_key"`
}

func (adr *MyUrlString) UnmarshalJSON(data []byte) error {
	value := strings.ReplaceAll(string(data), "\"", "")
	_, err := url.Parse(value)
	if err != nil {
		fmt.Errorf("ошибка в формате адреса %w", err)
		panic(err)
	}
	*adr = MyUrlString(value)
	return nil
}

func (adr *MyUrlString) Decode(value string) error {
	_, err := url.Parse(value)
	if err != nil {
		fmt.Errorf("ошибка в формате адреса %w", err)
		return err
	}
	*adr = MyUrlString(value)
	return nil
}

/* func (adr *Specification) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var testAdr string
	err := unmarshal(&testAdr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = url.Parse(testAdr)
	fmt.Println("unmarshaller111 ")
	if err != nil {
		fmt.Errorf("ошибка в формате адреса %w", err)
		return err
	}

	fmt.Println("unmarshaller ", adr)
	fmt.Println("unmarshaller ", testAdr)

	//adr.DB_url = testAdr
	return nil
} */

func ReadConf() (spec Specification) {

	err := envconfig.Process("", &spec)
	if err != nil {
		fmt.Errorf("ошибка чтения конфигурации %w", err)
	}
	return spec
}

func ReadYaml(file string) (spYA Specification) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Errorf("ошибка открытия файла %w", err)
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&spYA)
	if err != nil {
		fmt.Errorf("ошибка чтения данных конфигурации %w", err)
	}

	return spYA
}

func ReadJsonData(file string) (data Specification) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Errorf("ошибка открытия файла %w", err)
	}

	defer f.Close()

	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		fmt.Errorf("ошибка чтения данных конфигурации %w", err)
	}

	return data
}
