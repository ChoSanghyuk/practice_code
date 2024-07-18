package go_error

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestErrorJoin(t *testing.T) {

	var err1 error = errors.Join(Err1, Err2)

	fmt.Println(err1.Error())
	assert.IsEqual(errors.Is(err1, Err1), true)

}

func TestErrorJoin2(t *testing.T) {

	var err1 error = errors.Join(Err1, Err2)

	fmt.Println(err1.Error())
	assert.IsEqual(errors.Is(err1, Err2), true)

}

func TestErrorWrap(t *testing.T) {

	var err1 error = fmt.Errorf("Wrapped Error %w %w", Err1, ErrExistCheck)

	fmt.Println(err1.Error())
	assert.IsEqual(errors.Is(err1, Err1), true)
	assert.IsEqual(errors.Is(err1, ErrExistCheck), true)

	fmt.Println(errors.Unwrap(err1))
}
