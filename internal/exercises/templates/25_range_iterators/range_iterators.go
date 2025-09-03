package range_iterators

// TODO:
// - Build a simple integer iterator type with Next() (val, ok).
// - Provide a constructor to set start/end.
// - Implement IterateInts to apply a callback across the iterator range.

type IntIterator struct {
	current int
	end     int
}

func NewIntIterator(start, end int) *IntIterator {
	// TODO: initialize iterator state
	return &IntIterator{}
}

func (it *IntIterator) Next() (int, bool) {
	// TODO: return next value while current < end
	return 0, false
}

func IterateInts(start, end int, fn func(int)) {
	// TODO: iterate using the iterator and call fn on each value
}
