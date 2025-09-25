package calc

import "errors"

// Add returns a + b
func Add(a, b int) int {
    return a + b
}

// Divide returns a / b as float64. Returns error on division by zero.
func Divide(a, b int) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return float64(a) / float64(b), nil
}
