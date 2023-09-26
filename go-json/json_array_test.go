package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Akhmad",
		MiddleName: "Budi",
		LastName:   "Kurniawan",
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Akhmad","MiddleName":"Budi","LastName":"Kurniawan","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)

}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Budi",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "88888",
			},
			{
				Street:     "Jalan Lagi Dibangun",
				Country:    "Indonesia",
				PostalCode: "99999",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Print(string(bytes))

}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Budi","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"88888"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"99999"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)

}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "88888",
		},
		{
			Street:     "Jalan Lagi Dibangun",
			Country:    "Indonesia",
			PostalCode: "99999",
		},
	}

	bytes, _ := json.Marshal(addresses)

	fmt.Print(string(bytes))

}
