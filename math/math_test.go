package math_test

import (
	"github.com/AlejandroWaiz/Tests/math"
	"testing"
)



func TestInts (t *testing.T) {

	testTable := []struct{
		name	string
		numbers []int
		result int
	}{
		{"One to five",[]int{1,2,3,4,5}, 15},
		{"No numbers",nil,0},
		{"1 - 1",[]int{1,-1},0},
	}

	for _, currentTest := range testTable {

		t.Run(currentTest.name, func(t *testing.T) {

			sum := math.Ints(currentTest.numbers...)

			if sum != currentTest.result {
				t.Fatalf("Ups, case %v go wrong. Should be %v, got %v. ", currentTest.name, currentTest.result, sum)

			}
		})
	}
}