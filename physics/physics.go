package physics

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func Add(a int, b int) int {
	zero := mat.NewDense(3, 5, nil)
	fmt.Println(zero)
	return a + b
}
