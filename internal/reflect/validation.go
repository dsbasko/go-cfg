package reflect

import "reflect"

// Validation is a function that checks the validity of the struct pointer passed to it.
// It takes a struct pointer as an argument and returns an error if the struct pointer
// is invalid.
//
// Errors:
//
//	ErrNil: Returned when the struct pointer is invalid.
//	ErrNotPointer: Returned when the struct pointer is not a pointer.
//	ErrNotStruct: Returned when the struct pointer does not point to a struct.
func Validation(structPtr any) error {
	valueOf := reflect.ValueOf(structPtr)

	if valueOf.Kind() == reflect.Invalid {
		return ErrNil
	}

	if valueOf.Kind() != reflect.Ptr {
		return ErrNotPointer
	}

	valueOf = valueOf.Elem()
	if valueOf.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	return nil
}
