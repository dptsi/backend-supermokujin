package go_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
	Third  T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Akhmad",
		Second: "Budi",
		Third:  "Kurniawan",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Akhmad",
		Second: "Budi",
		Third:  "Kurniawan",
	}

	assert.Equal(t, "Erlang", data.ChangeFirst("Erlang"))
	assert.Equal(t, "Hello Erlang", data.SayHello("Erlang"))
}
