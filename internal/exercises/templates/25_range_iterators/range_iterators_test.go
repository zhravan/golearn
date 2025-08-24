package range_iterators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntIterator(t *testing.T) {
	it := NewIntIterator(1, 4)

	val, ok := it.Next()
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = it.Next()
	assert.True(t, ok)
	assert.Equal(t, 2, val)

	val, ok = it.Next()
	assert.True(t, ok)
	assert.Equal(t, 3, val)

	val, ok = it.Next()
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestIterateInts(t *testing.T) {
	sum := 0
	IterateInts(1, 4, func(val int) {
		sum += val
	})
	assert.Equal(t, 6, sum)
}
