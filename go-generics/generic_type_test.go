package go_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[X any](bag Bag[X]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Akhmad", "Budi", "Kurniawan"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 32, 4}
	PrintBag(numbers)
}
