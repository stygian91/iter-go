package iter

import "iter"

func Map[A, B any](seq iter.Seq[A], fn func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

func Map2[K, A, B any](seq iter.Seq2[K, A], fn func(A) B) iter.Seq2[K, B] {
	return func(yield func(K, B) bool) {
		for k, v := range seq {
			if !yield(k, fn(v)) {
				return
			}
		}
	}
}
