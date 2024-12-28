package iter_test

import (
	"maps"
	"reflect"
	"slices"
	"testing"

	"github.com/stygian91/iter-go"
)

func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	exp := []int{1, 4, 9, 16, 25}
	aSeq := iter.Map(slices.Values(a), func(a int) int { return a * a })
	res := slices.Collect(aSeq)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expected: %+v, got %+v\n", exp, res)
	}
}

func TestMap2(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	exp := map[string]int{"a": 2, "b": 4, "c": 6}
	seq := iter.Map2(maps.All(m), func(a int) int { return a * 2 })
	res := maps.Collect(seq)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expected: %+v, got %+v\n", exp, res)
	}
}
