package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type userData struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Address     struct {
		Street     string `json:"street"`
		City       string `json:"city"`
		PostalCode int    `json:"postalCode"`
	} `json:"address"`
	HomePage string `json:"homePage"`
}

func ReadUserData(file string) (data userData) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Errorf("ошибка открытия файла %w", err)
	}

	defer f.Close()

	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		fmt.Errorf("ошибка чтения данных пользователя %w", err)
	}

	return data
}
