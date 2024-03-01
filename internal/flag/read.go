package flag

import (
	"fmt"
	"os"
	rf "reflect"
	"sync"

	"github.com/spf13/pflag"

	"github.com/dsbasko/go-cfg/internal/dflt"
	"github.com/dsbasko/go-cfg/internal/reflect"
)

var (
	once    sync.Once
	flagSet *pflag.FlagSet
	dataPtr map[string]*string
)

// Read is a function that reads the input structure, validates it, reads the default values,
// parses the flags from the command line arguments and writes the values to the input structure.
// It returns an error if any of these operations fail.
func Read(structPtr any) error {
	if err := reflect.Validation(structPtr); err != nil {
		return fmt.Errorf("failed to validate in struct: %w", err)
	}

	if errDefault := dflt.Read(structPtr); errDefault != nil {
		return errDefault
	}

	once.Do(func() {
		dataPtr = make(map[string]*string)
		flagSet = pflag.NewFlagSet("cfg", pflag.ContinueOnError)
	})

	if err := parseFlags(structPtr, flagSet, ""); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	dataMap := make(map[string]string)
	for key, value := range dataPtr {
		dataMap[key] = *value
	}

	findFn := func(fieldName string) string { return dataMap[fieldName] }
	if err := reflect.WriteToStruct(structPtr, findFn); err != nil {
		return fmt.Errorf("failed to write to struct: %w", err)
	}

	return nil
}

// parseFlags is a recursive function that parses the flags from the input structure and
// the command line arguments. It adds the flags to the flagSet and the dataPtr map.
// It returns an error if the parsing fails.
func parseFlags(structPtr any, flagSet *pflag.FlagSet, prefix string) error {
	valueOf := rf.ValueOf(structPtr)
	if valueOf.Kind() == rf.Ptr {
		valueOf = valueOf.Elem()
	}

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Type().Field(i)
		flagFullName := field.Tag.Get("flag")
		flagShortName := field.Tag.Get("s-flag")
		flagUsage := field.Tag.Get("description")

		if field.Type.Kind() == rf.Struct {
			if err := parseFlags(
				valueOf.Field(i).Addr().Interface(),
				flagSet,
				fmt.Sprintf("%s%s.", prefix, field.Name),
			); err != nil {
				return fmt.Errorf("failed to parse flags: %w", err)
			}
		}

		fieldName := fmt.Sprintf("%s%s", prefix, field.Name)
		foundFlag := flagSet.Lookup(flagFullName)
		if _, ok := dataPtr[fieldName]; !ok && foundFlag == nil && (flagFullName != "" || flagShortName != "") {
			dataPtr[fieldName] = new(string)
			flagSet.StringVarP(dataPtr[fieldName], flagFullName, flagShortName, "", flagUsage)
		}
	}

	return nil
}
