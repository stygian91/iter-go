package iter_test

import (
	"testing"
	stditer "iter"

	"github.com/stygian91/iter-go"
)

func TestSeqToSeq2(t *testing.T) {
	input := "asdf"
	strit := iter.SeqToSeq2(iter.StrRuneIter(input))
	next, _ := stditer.Pull2(strit)

	for i := uint(0); i < uint(len(input)); i++ {
		idx, value, exists := next()
		if !exists {
			t.Errorf("Expected value for i=%d to exist", i)
			return
		}

		if idx != i {
			t.Errorf("Expected i and idx to be equal: i=%d, idx=%d", i, idx)
			return
		}

		expectedRune := rune(input[idx])
		if value != expectedRune {
			t.Errorf("Expected values to be equal: value=%v, expected=%v", value, expectedRune)
			return
		}
	}
}
