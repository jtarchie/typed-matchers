package typedassert

import (
	"testing"
)

var t *testing.T

func Setup(init *testing.T) {
	t = init
}

func Expect[T any](expected T) *Assertion[T] {
	return &Assertion[T]{
		value: expected,
	}
}
