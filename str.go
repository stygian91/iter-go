package iter

import stditer "iter"

func StrRuneIter(input string) stditer.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, char := range input {
			if !yield(char) {
				return
			}
		}
	}
}

func StrRuneIter2(input string) stditer.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, char := range input {
			if !yield(i, char) {
				return
			}
		}
	}
}
