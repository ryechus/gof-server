package utils_test

import (
	"github.com/placer14/gof-server/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	ints := []int64{0, 9, 10, 5}

	minInt := utils.Min(ints...)

	assert.Equal(t, minInt, int64(0))
}

func TestMax(t *testing.T) {
	ints := []int64{0, 9, 10, 5}

	maxInt := utils.Max(ints...)

	assert.Equal(t, maxInt, int64(10))
}

func TestMakeNewList(t *testing.T) {
	anyList := []any{1, 2, 3}
	intList := []int{1, 2, 3}
	newList := utils.MakeNewList[int](anyList)

	assert.Equal(t, newList, intList)
}

func TestIsGreaterThan(t *testing.T) {
	assert.Equal(t, utils.IsGreaterThan(6, 5), true)
	assert.Equal(t, utils.IsGreaterThan(5, 6), false)
	assert.Equal(t, utils.IsGreaterThan(5, 5), false)
}

func TestIsGreaterThanEqual(t *testing.T) {
	assert.Equal(t, utils.IsGreaterThanEqual(6, 5), true)
	assert.Equal(t, utils.IsGreaterThanEqual(6, 6), true)
	assert.Equal(t, utils.IsGreaterThanEqual(5, 6), false)
}

func TestIsLessThan(t *testing.T) {
	assert.Equal(t, utils.IsLessThan(5, 6), true)
	assert.Equal(t, utils.IsLessThan(6, 6), false)
	assert.Equal(t, utils.IsLessThan(6, 5), false)
}

func TestIsLessThanEqual(t *testing.T) {
	assert.True(t, utils.IsLessThanEqual(5, 6))
	assert.True(t, utils.IsLessThanEqual(5, 5))
	assert.False(t, utils.IsLessThanEqual(5, 4))
}

func TestSliceContains(t *testing.T) {
	stringList := []string{"a", "b", "c"}

	assert.True(t, utils.SliceContains(stringList, "a"))
	assert.False(t, utils.SliceContains(stringList, "d"))
}

func TestStringsContains(t *testing.T) {
	stringList := []string{"a quick", "brown fox", "jumped over"}

	assert.True(t, utils.StringsContains(stringList, "a quick brown fox jumped over"))
	assert.False(t, utils.StringsContains(stringList, "the lazy dog"))
}
