package types

import "fmt"

// AstTest this is a test type
type AstTest struct {
	A int // this is a test field
}

// Hello this is a test func
func (a AstTest) Hello() {
	fmt.Println("hello world")
}
