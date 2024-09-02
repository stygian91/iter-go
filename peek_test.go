package iter_test

import (
	stditer "iter"
	"testing"

	"github.com/stygian91/iter-go"
)

func strIter(input string) stditer.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}
}

func strIter2(input string) stditer.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, v := range input {
			if !yield(i, v) {
				return
			}
		}
	}
}

type peekT struct {
	actualV       rune
	actualValid   bool
	expectedV     rune
	expectedValid bool
}

func checkV(t *testing.T, tv peekT) {
	if tv.actualV != tv.expectedV {
		t.Fatalf("Expected value: %c, got: %c", tv.expectedV, tv.actualV)
	}

	if tv.actualValid != tv.expectedValid {
		t.Fatalf("Expected valid: %t, got: %t", tv.expectedValid, tv.actualValid)
	}
}

func checkPeek(t *testing.T, next, peek iter.GetFn[rune], expectedCurrV, expectedPeekV rune, expectedCurrValid, expectedPeekValid bool) {
	currV, currValid := next()
	peekV, peekValid := peek()

	checkKV(t, peek2T{
		actualV:       currV,
		actualValid:   currValid,
		expectedV:     expectedCurrV,
		expectedValid: expectedCurrValid,
	})

	checkKV(t, peek2T{
		actualV:       peekV,
		actualValid:   peekValid,
		expectedV:     expectedPeekV,
		expectedValid: expectedPeekValid,
	})
}

func TestPeekable(t *testing.T) {
	next, peek, _ := iter.Peek(strIter("abc"))

	checkNext := func(expectedCurrV, expectedPeekV rune, expectedCurrValid, expectedPeekValid bool) {
		checkPeek(t, next, peek,  expectedCurrV, expectedPeekV, expectedCurrValid, expectedPeekValid)
	}

	checkNext('a', 'b', true, true)
	checkNext('b', 'c', true, true)
	checkNext('c', 0, true, false)
	checkNext(0, 0, false, false)
}

func TestPeekableStop(t *testing.T) {
	next, peek, stop := iter.Peek(strIter("abc"))

	checkNext := func(expectedCurrV, expectedPeekV rune, expectedCurrValid, expectedPeekValid bool) {
		checkPeek(t, next, peek,  expectedCurrV, expectedPeekV, expectedCurrValid, expectedPeekValid)
	}

	checkNext('a', 'b', true, true)
	checkNext('b', 'c', true, true)
	stop()
	checkNext(0, 0, false, false)
}

func TestPeekablePeekRepeat(t *testing.T) {
	next, peek, _ := iter.Peek(strIter("abc"))

	currV, currValid := next()
	peekV, peekValid := peek()

	checkV(t, peekT{
		actualV:       currV,
		actualValid:   currValid,
		expectedV:     'a',
		expectedValid: true,
	})

	checkV(t, peekT{
		actualV:       peekV,
		actualValid:   peekValid,
		expectedV:     'b',
		expectedValid: true,
	})

	peekV, peekValid = peek()

	checkV(t, peekT{
		actualV:       currV,
		actualValid:   currValid,
		expectedV:     'a',
		expectedValid: true,
	})

	checkV(t, peekT{
		actualV:       peekV,
		actualValid:   peekValid,
		expectedV:     'b',
		expectedValid: true,
	})

	currV, currValid = next()
	peekV, peekValid = peek()

	checkV(t, peekT{
		actualV:       currV,
		actualValid:   currValid,
		expectedV:     'b',
		expectedValid: true,
	})

	checkV(t, peekT{
		actualV:       peekV,
		actualValid:   peekValid,
		expectedV:     'c',
		expectedValid: true,
	})

	peekV, peekValid = peek()

	checkV(t, peekT{
		actualV:       currV,
		actualValid:   currValid,
		expectedV:     'b',
		expectedValid: true,
	})

	checkV(t, peekT{
		actualV:       peekV,
		actualValid:   peekValid,
		expectedV:     'c',
		expectedValid: true,
	})
}

type peek2T struct {
	actualK       int
	actualV       rune
	actualValid   bool
	expectedK     int
	expectedV     rune
	expectedValid bool
}

func checkKV(t *testing.T, tv peek2T) {
	if tv.actualK != tv.expectedK {
		t.Fatalf("Expected key: %d, got: %d", tv.expectedK, tv.actualK)
	}

	if tv.actualV != tv.expectedV {
		t.Fatalf("Expected value: %c, got: %c", tv.expectedV, tv.actualV)
	}

	if tv.actualValid != tv.expectedValid {
		t.Fatalf("Expected valid: %t, got: %t", tv.expectedValid, tv.actualValid)
	}
}

func checkPeek2(t *testing.T, next, peek iter.GetFn2[int, rune], expectedCurrK, expectedPeekK int, expectedCurrV, expectedPeekV rune, expectedCurrValid, expectedPeekValid bool) {
	currK, currV, currValid := next()
	peekK, peekV, peekValid := peek()

	checkKV(t, peek2T{
		actualK:       currK,
		actualV:       currV,
		actualValid:   currValid,
		expectedK:     expectedCurrK,
		expectedV:     expectedCurrV,
		expectedValid: expectedCurrValid,
	})

	checkKV(t, peek2T{
		actualK:       peekK,
		actualV:       peekV,
		actualValid:   peekValid,
		expectedK:     expectedPeekK,
		expectedV:     expectedPeekV,
		expectedValid: expectedPeekValid,
	})
}

func TestPeekable2(t *testing.T) {
	next, peek, _ := iter.Peek2(strIter2("abc"))

	checkNext := func(expectedCurrK, expectedPeekK int, expectedCurrV, expectedPeekV rune, expectedCurrValid, expectedPeekValid bool) {
		checkPeek2(t, next, peek, expectedCurrK, expectedPeekK, expectedCurrV, expectedPeekV, expectedCurrValid, expectedPeekValid)
	}

	checkNext(0, 1, 'a', 'b', true, true)
	checkNext(1, 2, 'b', 'c', true, true)
	checkNext(2, 0, 'c', 0, true, false)
	checkNext(0, 0, 0, 0, false, false)
}

func TestPeekable2Stop(t *testing.T) {
	next, peek, stop := iter.Peek2(strIter2("abc"))

	checkNext := func(expectedCurrK, expectedPeekK int, expectedCurrV, expectedPeekV rune, expectedCurrValid, expectedPeekValid bool) {
		checkPeek2(t, next, peek, expectedCurrK, expectedPeekK, expectedCurrV, expectedPeekV, expectedCurrValid, expectedPeekValid)
	}

	checkNext(0, 1, 'a', 'b', true, true)
	checkNext(1, 2, 'b', 'c', true, true)
	stop()
	checkNext(0, 0, 0, 0, false, false)
}
