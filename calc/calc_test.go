package calc

import "testing"

func TestAdd(t *testing.T) {
    if got := Add(2, 3); got != 5 {
        t.Fatalf("Add(2,3) = %d; want 5", got)
    }
}

func TestDivideNormal(t *testing.T) {
    got, err := Divide(10, 2)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if got != 5.0 {
        t.Fatalf("Divide(10,2) = %v; want 5.0", got)
    }
}

func TestDivideByZero(t *testing.T) {
    if _, err := Divide(1, 0); err == nil {
        t.Fatalf("expected error for division by zero")
    }
}
