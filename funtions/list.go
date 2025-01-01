package funtions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ListCurrency() {

	file, err := os.ReadFile("../data/rates.json")
	if err != nil {
		log.Fatalf("Error to read the file, %v", err)
	}

	var rates map[string]float32
	err = json.Unmarshal(file, &rates)
	if err != nil {
		log.Fatalf("Error for decode the file in format JSON, %v", err)
	}

	fmt.Printf("List of currency available:\n")
	for currency, value := range rates {
		fmt.Printf("Currency: %s, Value: %.3f\n", currency, value)
	}
}
