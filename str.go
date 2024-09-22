package iter

import (
	"io"
	stditer "iter"
	"slices"
	"unicode/utf8"
)

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

func Utf8ReaderToRuneIter2(r io.Reader, size int) stditer.Seq2[int, rune] {
	i := 0
	buf := make([]byte, size)
	prev := []byte{}

	return func(yield func(int, rune) bool) {
		for {
			n, err := r.Read(buf[:])
			vbuf, ibuf := splitValidUtf8(buf[:n])
			vbuf = slices.Concat(prev, vbuf)

			for _, char := range StrRuneIter2(string(vbuf)) {
				if !yield(i, char) {
					return
				}
				i++
			}

			// check for error AFTER processing the buf, see https://pkg.go.dev/io#Reader.Read
			if err != nil {
				return
			}

			prev = ibuf
		}
	}
}

func splitValidUtf8(buf []byte) ([]byte, []byte) {
	for i := len(buf); i > 0; i-- {
		vbuf := buf[:i]

		if utf8.Valid(vbuf) {
			return vbuf, buf[i:]
		}
	}

	return []byte{}, buf
}
