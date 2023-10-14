package typedassert_test

import (
	"errors"
	"os"
	"testing"

	. "github.com/jtarchie/typedassert"
)

type myCustomType struct {
	s   string
	n   int
	f   float32
	arr []string
}

func TestEquals(t *testing.T) {
	Setup(t)

	Expect(5).Should(Equal(5))
	Expect(5.0).Should(Equal(5.0))

	Expect(5).ShouldNot(Equal(3))

	Expect("5").Should(Equal("5"))
	Expect([]int{1, 2}).Should(Equal([]int{1, 2}))
	Expect([]int{1, 2}).ShouldNot(Equal([]int{2, 1}))
	Expect([]byte{'f', 'o', 'o'}).Should(Equal([]byte{'f', 'o', 'o'}))
	Expect([]byte{'f', 'o', 'o'}).ShouldNot(Equal([]byte{'b', 'a', 'r'}))
	Expect(map[string]string{"a": "b", "c": "d"}).Should(Equal(map[string]string{"a": "b", "c": "d"}))
	Expect(map[string]string{"a": "b", "c": "d"}).ShouldNot(Equal(map[string]string{"a": "b", "c": "e"}))
	Expect(errors.New("foo")).Should(Equal(errors.New("foo")))
	Expect(errors.New("foo")).ShouldNot(Equal(errors.New("bar")))

	Expect(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b"}}).Should(Equal(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b"}}))
	Expect(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b"}}).ShouldNot(Equal(myCustomType{s: "bar", n: 3, f: 2.0, arr: []string{"a", "b"}}))
	Expect(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b"}}).ShouldNot(Equal(myCustomType{s: "foo", n: 2, f: 2.0, arr: []string{"a", "b"}}))
	Expect(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b"}}).ShouldNot(Equal(myCustomType{s: "foo", n: 3, f: 3.0, arr: []string{"a", "b"}}))
	Expect(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b"}}).ShouldNot(Equal(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b", "c"}}))
}

func TestBeADirectory(t *testing.T) {
	Expect("/dne/test").NotTo(BeADirectory())
	Expect(".").To(BeADirectory())

	tmpFile, err := os.CreateTemp("", "gomega-test-tempfile")
	Expect(err).ShouldNot(HaveOccurred())
	defer os.Remove(tmpFile.Name())
	Expect(tmpFile.Name()).ShouldNot(BeADirectory())
}