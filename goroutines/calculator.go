package goroutine

import "fmt"

type Calculator interface {
	Add(a, b interface{}) int
	Sub(a, b interface{}) int
	Mult(a, b interface{}) int
	Div(a, b interface{}) (int, error)
}

type calculator struct {
}

func typeOf(x interface{}) string { return fmt.Sprintf("%T", x) }

func (c calculator) Add(a, b int) int {

	return a + b
}

func (c calculator) Mult(a, b int) int {

	return a * b
}
func (c calculator) Sub(a, b int) int {

	return a - b
}
func (c calculator) Div(a, b int) (int, error) {
	var err error
	if b == 0 {
		err = fmt.Errorf("Division by zero")
	} else {
		return a / b, err
	}
	return 0, err
}

func CalculatorExample() {
	// calc := calculator{}
	// addRes := calc.Add(5, 6)
}
