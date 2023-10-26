package go_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length[string]("Akhmad")
	assert.Equal(t, "Akhmad", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
