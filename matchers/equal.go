package matchers

import (
	"reflect"

	"github.com/onsi/gomega/format"
)

type EqualMatcher[T any] struct {
	Expected T
}

func (e *EqualMatcher[T]) Match(actual T) (bool, error) {
	return reflect.DeepEqual(actual, e.Expected), nil
}

func (matcher *EqualMatcher[T]) FailureMessage(actual T) (message string) {
	actualString, ok := any(actual).(string)
	if ok {
		expectedString := any(matcher.Expected).(string)
		return format.MessageWithDiff(actualString, "to equal", expectedString)
	}

	return format.Message(actual, "to equal", matcher.Expected)
}

func (matcher *EqualMatcher[T]) NegatedFailureMessage(actual T) (message string) {
	return format.Message(actual, "not to equal", matcher.Expected)
}