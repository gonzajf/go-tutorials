package arrays_slices

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("got %d but expected was %d given %v", got, expected, numbers)
	}

}