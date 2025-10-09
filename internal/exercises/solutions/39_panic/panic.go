package panicex

// CausePanic panics if input is negative, otherwise returns input
func CausePanic(n int) int {
    if n < 0 {
        panic("negative number")
    }
    return n
}

// SafeDivision divides a by b, recovers from panic if denominator is zero
func SafeDivision(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            result = 0
        }
    }()
    if b == 0 {
        panic("divide by zero")
    }
    return a / b
}

// TriggerMultiplePanics demonstrates defer + recover in a loop
func TriggerMultiplePanics(nums []int) []string {
    results := make([]string, len(nums))
    for i, v := range nums {
        func(idx, val int) {
            defer func() {
                if r := recover(); r != nil {
                    results[idx] = "recovered panic"
                }
            }()
            if val < 0 {
                panic("panic for negative")
            }
            results[idx] = "ok"
        }(i, v)
    }
    return results
}

// PanicWithMessage panics with a custom message
func PanicWithMessage(msg string) {
    panic(msg)
}
