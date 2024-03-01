package reflect

import "fmt"

var (
	// ErrNil is returned when the value is nil
	ErrNil = fmt.Errorf("should not be nil")

	// ErrNotPointer is returned when the value is not a pointer
	ErrNotPointer = fmt.Errorf("must be a pointer to a struct")

	// ErrNotStruct is returned when the value is not a struct
	ErrNotStruct = fmt.Errorf("must be a pointer to a struct")
)
