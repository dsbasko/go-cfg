package dflt

import (
	"fmt"
	"sync"

	"github.com/dsbasko/go-cfg/internal/reflect"
)

var once sync.Once

// Read is a function that parses default values into the provided cfg structure.
// The cfg parameter should be a pointer to a struct where each field represents a default value.
// The function returns an error if the parsing process fails, wrapping the original error with a message.
func Read(structPtr any) error {
	if err := reflect.Validation(structPtr); err != nil {
		return fmt.Errorf("error validating struct: %w", err)
	}

	var result error
	once.Do(func() {
		parsedStruct, err := reflect.ParseTag(structPtr, "default")
		if err != nil {
			result = fmt.Errorf("error parsing struct: %w", err)
		}

		if err = reflect.WriteToStruct(structPtr, func(fieldName string) string {
			return parsedStruct[fieldName].TagValue
		}); err != nil {
			result = fmt.Errorf("error writing to struct: %w", err)
		}
	})

	return result
}
