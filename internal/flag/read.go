package flag

import (
	"fmt"
	"reflect"

	"github.com/integrii/flaggy"
)

// Read is a function that parses command-line flags into the provided cfg structure.
// The cfg parameter should be a pointer to a struct where each field represents a command-line flag.
// The function returns an error if the parsing process fails, wrapping the original error with a message.
func Read(cfg any) error {
	if err := parseFlags(cfg, 1); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	return nil
}

// parseFlags is a helper function used by Read to parse the command-line flags.
// It uses reflection to iterate over the fields of the cfg structure and sets the corresponding command-line flags.
// The function returns an error if the cfg parameter is not a non-nil pointer to a struct or if the parsing process fails.
func parseFlags(cfg any, deepCount int) error { //nolint:gocyclo
	flaggy.DefaultParser.ShowHelpWithHFlag = false
	flaggy.DefaultParser.ShowVersionWithVersionFlag = false

	valueOf := reflect.ValueOf(cfg)

	if valueOf.Kind() != reflect.Ptr || valueOf.IsNil() {
		return fmt.Errorf("must be a non-nil pointer to a struct")
	}

	valueOf = valueOf.Elem()
	if valueOf.Kind() != reflect.Struct {
		return fmt.Errorf("must be a pointer to a struct")
	}

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Type().Field(i)
		flagFullName := field.Tag.Get("flag")
		flagShortName := field.Tag.Get("s-flag")
		flagDescription := field.Tag.Get("description")

		// If the field is a struct, the function calls itself recursively to parse the flags in the struct.
		if field.Type.Kind() == reflect.Struct {
			if err := parseFlags(valueOf.Field(i).Addr().Interface(), deepCount+1); err != nil {
				return err
			}
		}

		if flagFullName != "" || flagShortName != "" {
			// The switch statement is used to handle different types of fields in the cfg structure.
			// For each type, it sets the corresponding command-line flag using the flaggy package.
			switch field.Type.Kind() {
			case reflect.String:
				flaggy.String(valueOf.Field(i).Addr().Interface().(*string), flagShortName, flagFullName, flagDescription)
			case reflect.Int:
				flaggy.Int(valueOf.Field(i).Addr().Interface().(*int), flagShortName, flagFullName, flagDescription)
			case reflect.Int8:
				flaggy.Int8(valueOf.Field(i).Addr().Interface().(*int8), flagShortName, flagFullName, flagDescription)
			case reflect.Int16:
				flaggy.Int16(valueOf.Field(i).Addr().Interface().(*int16), flagShortName, flagFullName, flagDescription)
			case reflect.Int32:
				flaggy.Int32(valueOf.Field(i).Addr().Interface().(*int32), flagShortName, flagFullName, flagDescription)
			case reflect.Int64:
				flaggy.Int64(valueOf.Field(i).Addr().Interface().(*int64), flagShortName, flagFullName, flagDescription)
			case reflect.Uint:
				flaggy.UInt(valueOf.Field(i).Addr().Interface().(*uint), flagShortName, flagFullName, flagDescription)
			case reflect.Uint8:
				flaggy.UInt8(valueOf.Field(i).Addr().Interface().(*uint8), flagShortName, flagFullName, flagDescription)
			case reflect.Uint16:
				flaggy.UInt16(valueOf.Field(i).Addr().Interface().(*uint16), flagShortName, flagFullName, flagDescription)
			case reflect.Uint32:
				flaggy.UInt32(valueOf.Field(i).Addr().Interface().(*uint32), flagShortName, flagFullName, flagDescription)
			case reflect.Uint64:
				flaggy.UInt64(valueOf.Field(i).Addr().Interface().(*uint64), flagShortName, flagFullName, flagDescription)
			case reflect.Float64:
				flaggy.Float64(valueOf.Field(i).Addr().Interface().(*float64), flagShortName, flagFullName, flagDescription)
			case reflect.Float32:
				flaggy.Float32(valueOf.Field(i).Addr().Interface().(*float32), flagShortName, flagFullName, flagDescription)
			case reflect.Bool:
				flaggy.Bool(valueOf.Field(i).Addr().Interface().(*bool), flagShortName, flagFullName, flagDescription)
			}
		}
	}

	// The flaggy.Parse function is called only once to parse the command-line flags.
	if deepCount == 1 {
		flaggy.Parse()
		flaggy.ResetParser()
	}

	return nil
}
