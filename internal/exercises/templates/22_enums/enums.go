package enums

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
	return d == Sunday || d == Saturday
}
