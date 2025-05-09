package utils

import (
	"golang.org/x/exp/constraints"
	"slices"
	"strings"
)

func Min[T constraints.Ordered](args ...T) T {
	_min := args[0]
	for _, x := range args {
		if x < _min {
			_min = x
		}
	}
	return _min
}

func Max[T constraints.Ordered](args ...T) T {
	_max := args[0]
	for _, x := range args {
		if x > _max {
			_max = x
		}
	}
	return _max
}

func MakeNewList[T constraints.Ordered](items []any) []T {
	newList := make([]T, len(items))
	for idx, v := range items {
		newList[idx] = v.(T)
	}
	return newList
}

func IsGreaterThan[T constraints.Ordered](a, b T) bool {
	return a > b
}

func IsLessThan[T constraints.Ordered](a, b T) bool {
	return a < b
}

func IsGreaterThanEqual[T constraints.Ordered](a, b T) bool {
	return a >= b
}

func IsLessThanEqual[T constraints.Ordered](a, b T) bool {
	return a <= b
}

func SliceContains[T constraints.Ordered](items []T, item T) bool {
	return slices.Contains(items, item)
}

func StringsContains(items []string, item string) bool {
	for _, i := range items {
		if strings.Contains(item, i) {
			return true
		}
	}
	return false
}
