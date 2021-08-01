package math_test

import (
	"fmt"
	"github.com/AlejandroWaiz/Tests/math"
)

func ExampleInts() {
	result := math.Ints(1,2,3,4,5)

	fmt.Println("Sum from one to five is", result)

	//Output:
	//Sum from one to five is 15
}