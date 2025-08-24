package pointers

func modifyValue(ptr *int) *int {
	*ptr = 100
	return ptr
}
