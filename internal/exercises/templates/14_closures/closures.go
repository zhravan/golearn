package closures

// TODO:
// - Implement a closure function that captures state and returns a func() int
//   which increments and returns an internal counter on each call.
// - Each call to closure() must produce an independent counter.
// - Do not change function names or signatures; tests call closure().

// NOTE: Keep imports minimal in skeletons to avoid unused import errors.

func closure() func() int {
	// TODO: implement a counting closure
	return func() int { return 0 }
}
