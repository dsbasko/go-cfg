package structs

import (
	"fmt"
	"reflect"
	"strconv"
)

// FromMap maps the values from the data map to the fields of the struct.
// The inMap map is a map[string]string where the key is the name of the field
// and the value is the value to be set in the field.
// The inStruct must be a pointer to a struct.
func FromMap(inMap map[string]string, inStruct any) error { // nolint: funlen,gocyclo,nolintlint
	valueOf := reflect.ValueOf(inStruct)

	if valueOf.Kind() != reflect.Ptr || valueOf.IsNil() {
		return fmt.Errorf("must be a non-nil pointer to a struct")
	}

	valueOf = valueOf.Elem()
	if valueOf.Kind() != reflect.Struct {
		return fmt.Errorf("must be a pointer to a struct")
	}

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Type().Field(i)

		// If the field is a struct, call FromMap recursively
		// to map the values from the data map to the fields of the struct.
		if field.Type.Kind() == reflect.Struct {
			if err := FromMap(inMap, valueOf.Field(i).Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		tag := field.Tag.Get("env")
		if tag != "" {
			switch field.Type.Kind() {
			case reflect.String:
				valueOf.Field(i).SetString(inMap[tag])
			case reflect.Int:
				valInt, _ := strconv.ParseInt(inMap[tag], 10, 0)
				valueOf.Field(i).SetInt(valInt)
			case reflect.Int8:
				valInt, _ := strconv.ParseInt(inMap[tag], 10, 8)
				valueOf.Field(i).SetInt(valInt)
			case reflect.Int16:
				valInt, _ := strconv.ParseInt(inMap[tag], 10, 16)
				valueOf.Field(i).SetInt(valInt)
			case reflect.Int32:
				valInt, _ := strconv.ParseInt(inMap[tag], 10, 32)
				valueOf.Field(i).SetInt(valInt)
			case reflect.Int64:
				valInt, _ := strconv.ParseInt(inMap[tag], 10, 64)
				valueOf.Field(i).SetInt(valInt)
			case reflect.Uint:
				valUint, _ := strconv.ParseUint(inMap[tag], 10, 0)
				valueOf.Field(i).SetUint(valUint)
			case reflect.Uint8:
				valUint, _ := strconv.ParseUint(inMap[tag], 10, 8)
				valueOf.Field(i).SetUint(valUint)
			case reflect.Uint16:
				valUint, _ := strconv.ParseUint(inMap[tag], 10, 16)
				valueOf.Field(i).SetUint(valUint)
			case reflect.Uint32:
				valUint, _ := strconv.ParseUint(inMap[tag], 10, 32)
				valueOf.Field(i).SetUint(valUint)
			case reflect.Uint64:
				valUint, _ := strconv.ParseUint(inMap[tag], 10, 64)
				valueOf.Field(i).SetUint(valUint)
			case reflect.Float64:
				valFloat, _ := strconv.ParseFloat(inMap[tag], 64)
				valueOf.Field(i).SetFloat(valFloat)
			case reflect.Float32:
				valFloat, _ := strconv.ParseFloat(inMap[tag], 32)
				valueOf.Field(i).SetFloat(valFloat)
			case reflect.Bool:
				valBool, _ := strconv.ParseBool(inMap[tag])
				valueOf.Field(i).SetBool(valBool)
			}
		}
	}

	return nil
}
