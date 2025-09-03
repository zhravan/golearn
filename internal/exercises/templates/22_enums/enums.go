package enums

// TODO:
// - Create an enum-like type using iota for days of the week.
// - Implement IsWeekend to return true for Saturday and Sunday.

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func IsWeekend(d Day) bool {
	// TODO: return true for Saturday or Sunday
	return false
}
