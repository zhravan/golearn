package exercises

// WithTestCatalogLoader replaces the catalog loader temporarily for a test.
// This allows you to inject mock data and reset the singleton state.
func WithTestCatalogLoader(mockLoader func() (Catalog, error), testFunc func()) {
	catalogMu.Lock()

	// Save old state
	oldLoader := catalogLoader
	oldData := catalogData

	// Swap in mock and force reload
	catalogLoader = mockLoader
	catalogData = nil

	catalogMu.Unlock()

	defer func() {
		catalogMu.Lock()
		// Restore old state
		catalogLoader = oldLoader
		catalogData = oldData
		catalogMu.Unlock()
	}()

	testFunc()
}
