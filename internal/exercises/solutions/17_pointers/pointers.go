package pointers

func modifyValue(ptr *int) *int {
	if ptr == nil {
		return nil
	}
	*ptr = 100
	return ptr
}
