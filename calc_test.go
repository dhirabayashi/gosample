package gosample

import "testing"

func TestAdd(t *testing.T) {
	answer := Add(10, 2)
	if answer != 12 {
		t.Error("answer should be 12.")
	}
}

func TestSubtract(t *testing.T) {
	answer := Subtract(10, 2)
	if answer != 8 {
		t.Error("answer should be 8.")
	}
}

func TestMultiply(t *testing.T) {
	answer := Multiply(10, 2)
	if answer != 20 {
		t.Error("answer should be 20.")
	}
}

func TestDivide(t *testing.T) {
	answer, rest, err := Divide(10, 2)
	if answer != 5 {
		t.Error("answer should be 5.")
	}
	if rest != 0 {
		t.Error("rest should be 0.")
	}
	if err != nil {
		t.Error("err should be nil.")
	}
}

func TestDivideRestIsNotZero(t *testing.T) {
	answer, rest, err := Divide(10, 3)
	if answer != 3 {
		t.Error("answer should be 3.")
	}
	if rest != 1 {
		t.Error("rest should be 1.")
	}
	if err != nil {
		t.Error("err should be nil.")
	}
}

func TestDivideByZero(t *testing.T) {
	_, _, err := Divide(10, 0)

	if err == nil {
		t.Error("err should not be nil.")
	}
}
