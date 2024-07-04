package go_error

import (
	"errors"
	"go_practice/go_error/types"
	"testing"
)

func TestGenCusoError1(t *testing.T) {

	var err1 error = errors.Join(CustomError1, errors.New("custom error 1"))
	errors.Is(err1, CustomError1)

}

func main() {

	err := types.GenCusoError2("custom error 2")

	switch err.(type) {
	case types.CustomError1:
		println("CustomError1")
	case types.CustomError2:
		println("CustomError2")
	default:
		println("Unknown error")
	}
}
