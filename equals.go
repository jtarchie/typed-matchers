package gomega

import (
	"fmt"
	"testing"
)

var t *testing.T

func Setup(init *testing.T) {
	t = init
}

type Matcher[T any] interface {
	Match(actual T) (success bool, err error)
	// FailureMessage(actual T) (message string)
	// NegatedFailureMessage(actual T) (message string)
}

type Assertion[T comparable] struct {
	value T
}

func (a *Assertion[T]) To(matcher Matcher[T]) bool {
	success, err := matcher.Match(a.value)
	if err != nil {
		t.Error(err.Error())
	}

	return success
}

type EqualMatcher[T comparable] struct {
	expected T
}

func (e *EqualMatcher[T]) Match(actual T) (bool, error) {
	if e.expected == actual {
		return true, nil
	}

	return false, fmt.Errorf("expected %v to equal %v", e.expected, actual)
}

func Equal[T comparable](value T) *EqualMatcher[T] {
	return &EqualMatcher[T]{
		expected: value,
	}
}

func Expect[T comparable](expected T) *Assertion[T] {
	return &Assertion[T]{
		value: expected,
	}
}
