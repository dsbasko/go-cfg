package reflect

import (
	"fmt"
	"reflect"
	"strconv"
)

// WriteToStruct is a function that takes a pointer to a struct and a function as
// arguments. The function argument should take a string (field name) and return a string.
// It uses reflection to iterate over the fields of the struct and calls the provided
// function with the field name. The returned value from the function is then used to set
// the value of the field in the struct. If the field is another struct, it recursively
// calls itself to set the values of the nested struct's fields.
func WriteToStruct(structPtr any, fn func(fieldName string) string) error {
	return writeToStructRecursive(structPtr, fn, "")
}

// writeToStructRecursive is a helper function for WriteToStruct. It takes a pointer to a
// struct, a function, and a prefix string as arguments. The function argument should take
// a string (field name) and return a string. The prefix is used to build the field name
// for nested struct fields. It uses reflection to iterate over the fields of the struct
// and calls the provided function with the field name. The returned value from the
// function is then used to set the value of the field in the struct. If the field is
// another struct, it recursively calls itself to set the values of the nested
// struct's fields.
func writeToStructRecursive(structPtr any, fn func(fieldName string) string, prefix string) error { //nolint:funlen,gocyclo
	valueOf := reflect.ValueOf(structPtr)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Type().Field(i)

		if field.Type.Kind() == reflect.Struct {
			_ = writeToStructRecursive(
				valueOf.Field(i).Addr().Interface(),
				fn,
				fmt.Sprintf("%s%s.", prefix, field.Name),
			)
			continue
		}

		fieldName := fmt.Sprintf("%s%s", prefix, field.Name)
		switch field.Type.Kind() {
		case reflect.String:
			val := fn(fieldName)
			if val != "" {
				valueOf.Field(i).SetString(fn(fieldName))
			}
		case reflect.Int:
			valInt, _ := strconv.ParseInt(fn(fieldName), 10, 0)
			if valInt != 0 {
				valueOf.Field(i).SetInt(valInt)
			}
		case reflect.Int8:
			valInt, _ := strconv.ParseInt(fn(fieldName), 10, 8)
			if valInt != 0 {
				valueOf.Field(i).SetInt(valInt)
			}
		case reflect.Int16:
			valInt, _ := strconv.ParseInt(fn(fieldName), 10, 16)
			if valInt != 0 {
				valueOf.Field(i).SetInt(valInt)
			}
		case reflect.Int32:
			valInt, _ := strconv.ParseInt(fn(fieldName), 10, 32)
			if valInt != 0 {
				valueOf.Field(i).SetInt(valInt)
			}
		case reflect.Int64:
			valInt, _ := strconv.ParseInt(fn(fieldName), 10, 64)
			if valInt != 0 {
				valueOf.Field(i).SetInt(valInt)
			}
		case reflect.Uint:
			valUint, _ := strconv.ParseUint(fn(fieldName), 10, 0)
			if valUint != 0 {
				valueOf.Field(i).SetUint(valUint)
			}
		case reflect.Uint8:
			valUint, _ := strconv.ParseUint(fn(fieldName), 10, 8)
			if valUint != 0 {
				valueOf.Field(i).SetUint(valUint)
			}
		case reflect.Uint16:
			valUint, _ := strconv.ParseUint(fn(fieldName), 10, 16)
			if valUint != 0 {
				valueOf.Field(i).SetUint(valUint)
			}
		case reflect.Uint32:
			valUint, _ := strconv.ParseUint(fn(fieldName), 10, 32)
			if valUint != 0 {
				valueOf.Field(i).SetUint(valUint)
			}
		case reflect.Uint64:
			valUint, _ := strconv.ParseUint(fn(fieldName), 10, 64)
			if valUint != 0 {
				valueOf.Field(i).SetUint(valUint)
			}
		case reflect.Float32:
			valFloat, _ := strconv.ParseFloat(fn(fieldName), 32)
			if valFloat != 0 {
				valueOf.Field(i).SetFloat(valFloat)
			}
		case reflect.Float64:
			valFloat, _ := strconv.ParseFloat(fn(fieldName), 64)
			if valFloat != 0 {
				valueOf.Field(i).SetFloat(valFloat)
			}
		case reflect.Bool:
			valBool, _ := strconv.ParseBool(fn(fieldName))
			valueOf.Field(i).SetBool(valBool)
		}
	}

	return nil
}
