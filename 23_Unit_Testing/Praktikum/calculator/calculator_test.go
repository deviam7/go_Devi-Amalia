package kalkulator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5 but got %v", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(3, 2)
    if result != 1 {
        t.Errorf("Expected 1 but got %v", result)
    }
}

func TestMultiply(t *testing.T) {
    result := Multiply(2, 3)
    if result != 6 {
        t.Errorf("Expected 6 but got %v", result)
    }
}

func TestDivide(t *testing.T) {
    t.Run("valid division", func(t *testing.T) {
        result, err := Divide(6, 3)
        if err != nil {
            t.Errorf("Expected no error but got %v", err)
        }
        if result != 2 {
            t.Errorf("Expected 2 but got %v", result)
        }
    })

    t.Run("division by zero", func(t *testing.T) {
        _, err := Divide(6, 0)
        if err == nil {
            t.Errorf("Expected an error but got none")
        }
        if err.Error() != "division by zero" {
            t.Errorf("Expected 'division by zero' error but got '%v'", err)
        }
    })
}
