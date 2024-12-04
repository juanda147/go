package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func GetFloatFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, _ := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteFloatToFile(accountBalance float64) {
	valueText := fmt.Sprint(accountBalance)
	os.WriteFile(accountBalanceFile, []byte(valueText), 0644)
}
