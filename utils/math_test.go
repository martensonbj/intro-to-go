package utils

import "testing"

func TestAdd(t *testing.T) {
	total := Add(2, 3, 4)

	if total != 9 {
		t.Errorf("Sum was incorrect! Received %d, Got: %d", total, 9)
	}
}
