package day14

import (
	"testing"
)

func TestWrapVector(t *testing.T) {

	startingVector := vector{8, 8, 4, 4}
	expected := vector{2, 2, 4, 4}
	actual := startingVector.Move(10, 10, 1)
	if expected != actual {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}

	startingVector = vector{5, 5, -2, 3}
	expected = vector{7, 2, -2, 3}
	actual = startingVector.Move(10, 10, 9)
	if expected != actual {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
