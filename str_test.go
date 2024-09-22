package iter_test

import (
	"strings"
	"testing"

	"github.com/stygian91/iter-go"
)

func TestUtf8ReaderToRuneIter2(t *testing.T) {
	input := "lorem ipsum"
	r := strings.NewReader(input)
	expected := []rune{'l', 'o', 'r', 'e', 'm', ' ', 'i', 'p', 's', 'u', 'm'}

	for i, c := range iter.Utf8ReaderToRuneIter2(r, 5) {
		if expected[i] != c {
			t.Errorf("Expected %s, found %s at position %d in %s", string(expected[i]), string(c), i, input)
			return
		}
	}
}
