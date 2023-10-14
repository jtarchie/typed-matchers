package typedassert

import "github.com/jtarchie/typedassert/matchers"

type Matcher[T any] interface {
	Match(actual T) (success bool, err error)
	FailureMessage(actual T) (message string)
	NegatedFailureMessage(actual T) (message string)
}

func Equal[T any](value T) *matchers.EqualMatcher[T] {
	return &matchers.EqualMatcher[T]{
		Expected: value,
	}
}

func BeADirectory() *matchers.BeADirectoryMatcher {
	return &matchers.BeADirectoryMatcher{}
}

func HaveOccurred() *matchers.HaveOccurredMatcher {
	return &matchers.HaveOccurredMatcher{}
}