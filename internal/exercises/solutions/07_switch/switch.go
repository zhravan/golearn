package sw

// IsWeekend returns true if day is weekend ("sat" or "sun"), false otherwise
func IsWeekend(day string) bool {
	switch day {
	case "sat", "sun":
		return true
	}
	return false
}
