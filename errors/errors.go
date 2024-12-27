package errors

import "fmt"

type AppError struct {
	err           error
	Custom, Field string
}

func (e *AppError) Error() string {

	fmt.Println(e.Custom)
	return e.err.Error()
}

func (e *AppError) Unwrap() error {
	return e.err
}

func MainFunc() {
	err := method1()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(err)
	fmt.Println(err.Unwrap())
}

func method1() *AppError {
	return method2()
}

func method2() *AppError {
	return &AppError{
		err:    fmt.Errorf("error"),
		Custom: "Something goes wrong",
	}
}
