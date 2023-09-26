package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestJSONStreamEncoder(t *testing.T) {

	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Akhmad",
		MiddleName: "Budi",
		LastName:   "Kurniawan",
	}
	encoder.Encode(customer)

}
