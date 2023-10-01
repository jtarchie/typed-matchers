package typedassert

type Assertion[T any] struct {
	value T
}

func (a *Assertion[T]) To(matcher Matcher[T]) bool {
	success, err := matcher.Match(a.value)
	if err != nil {
		t.Error(err.Error())
	}
	if !success {
		t.Error(matcher.FailureMessage(a.value))
	}

	return success
}

func (a *Assertion[T]) Should(matcher Matcher[T]) bool {
	return a.To(matcher)
}

func (a *Assertion[T]) NotTo(matcher Matcher[T]) bool {
	success, err := matcher.Match(a.value)
	if err != nil {
		t.Error(err.Error())
	}

	if success {
		t.Error(matcher.NegatedFailureMessage(a.value))
	}

	return success
}

func (a *Assertion[T]) ShouldNot(matcher Matcher[T]) bool {
	return a.NotTo(matcher)
}