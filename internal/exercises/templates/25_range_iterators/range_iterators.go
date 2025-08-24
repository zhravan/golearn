package range_iterators

type IntIterator struct {
	current int
	end     int
}

func NewIntIterator(start, end int) *IntIterator {
	return &IntIterator{current: start, end: end}
}

func (it *IntIterator) Next() (int, bool) {
	if it.current < it.end {
		val := it.current
		it.current++
		return val, true
	}
	return 0, false
}

func IterateInts(start, end int, fn func(int)) {
	it := NewIntIterator(start, end)
	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		fn(val)
	}
}
