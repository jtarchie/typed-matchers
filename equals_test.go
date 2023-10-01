package gomega_test

import (
	"testing"

	"github.com/jtarchie/gomega"
)

func TestEquals(t *testing.T) {
	gomega.Setup(t)

	gomega.Expect(1).To(gomega.Equal(1))
	gomega.Expect(1).To(gomega.Equal("asdf"))
}