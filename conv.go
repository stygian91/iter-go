package iter

import stditer "iter"

func SeqToSeq2[T any](it stditer.Seq[T]) stditer.Seq2[uint, T] {
	return func(yield func(uint, T) bool) {
		i := uint(0)
		for value := range it {
			if !yield(i, value) {
				return
			}
			i++
		}
	}
}
