package basicerrors

import (
	"errors"
	"fmt"
)

var ErrorValue = errors.New("this is a typed error")

type TypedError struct {
	error
}

func BasicErrors() {
	err := errors.New("This is a quick and eqsy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("value err: ", err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error: ", err)
}
