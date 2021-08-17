package main

import "errors"

type MyError struct{ hint string }

func (m MyError) Error() string {
	return ""
}
func (MyError) Is(target interface{}) bool { return true } // target is interface{} instead of error.
func Foo() bool {
	x, y := MyError{"A"}, MyError{"B"}
	return errors.Is(x, y) // returns false as x != y and MyError does not have an `Is(error) bool` function.
}

func main() {
}
