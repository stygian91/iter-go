package iter

import stditer "iter"

type GetFn[V any] func() (V, bool)
type GetFn2[K, V any] func() (K, V, bool)
type stopFn func()

func Peek[V any](it stditer.Seq[V]) (GetFn[V], GetFn[V], stopFn) {
	_next, _stop := stditer.Pull(it)
	shouldUpdatePeek := true
	stopped := false
	var peekV, zeroV V
	var peekValid bool

	next := func() (V, bool) {
		if stopped {
			return zeroV, false
		}

		if !shouldUpdatePeek {
			shouldUpdatePeek = true
			return peekV, peekValid
		}

		return _next()
	}

	peek := func() (V, bool) {
		if stopped {
			return zeroV, false
		}

		if shouldUpdatePeek {
			peekV, peekValid = next()
			shouldUpdatePeek = false
		}

		return peekV, peekValid
	}

	stop := func() {
		stopped = true
		_stop()
	}

	return next, peek, stop
}

func Peek2[K, V any](it stditer.Seq2[K, V]) (GetFn2[K, V], GetFn2[K, V], stopFn) {
	_next, _stop := stditer.Pull2(it)
	shouldUpdatePeek := true
	stopped := false
	var peekK, zeroK K
	var peekV, zeroV V
	var peekValid bool

	next := func() (K, V, bool) {
		if stopped {
			return zeroK, zeroV, false
		}

		if !shouldUpdatePeek {
			shouldUpdatePeek = true
			return peekK, peekV, peekValid
		}

		return _next()
	}

	peek := func() (K, V, bool) {
		if stopped {
			return zeroK, zeroV, false
		}

		if shouldUpdatePeek {
			peekK, peekV, peekValid = next()
			shouldUpdatePeek = false
		}

		return peekK, peekV, peekValid
	}

	stop := func() {
		stopped = true
		_stop()
	}

	return next, peek, stop
}
