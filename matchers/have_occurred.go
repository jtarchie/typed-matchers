package matchers

import (
	"fmt"

	"github.com/onsi/gomega/format"
)

type HaveOccurredMatcher struct {
}

func (matcher *HaveOccurredMatcher) Match(actual error) (success bool, err error) {
	return actual != nil, nil
}

func (matcher *HaveOccurredMatcher) FailureMessage(actual error) (message string) {
	return fmt.Sprintf("Expected an error to have occurred.  Got:\n%s", format.Object(actual, 1))
}

func (matcher *HaveOccurredMatcher) NegatedFailureMessage(actual error) (message string) {
	return fmt.Sprintf("Unexpected error:\n%s\n%s", format.Object(actual, 1), "occurred")
}
