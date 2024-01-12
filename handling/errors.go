// +build ignore
package main

import (
    "fmt"
)

// CustomError is a type that implements the error interface.
type CustomError struct {
    Message string
}

// Error method makes CustomError implement the error interface.
func (e *CustomError) Error() string {
    return e.Message
}

// SomeFunction is a function that returns an error.
func SomeFunction() error {
    return &CustomError{Message: "Something went wrong"}
}

func main() {
    err := SomeFunction()
    if err != nil {
        fmt.Println(err)
    }
}