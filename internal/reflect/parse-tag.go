package reflect

import (
	"fmt"
	"reflect"
)

// StructFields represents a struct field and its associated tag.
type StructFields struct {
	FieldName string // The name of the field.
	TagName   string // The name of the tag.
	TagValue  string // The value of the tag.
}

// ParseTag parses the tags of the struct pointed to by structPtr.
// It returns a map where the keys are the fully qualified names of the fields and the values are the associated tags.
// If a field is a nested struct, its fields are also included in the map, with their names prefixed by the name of the parent struct.
// The function returns an error if structPtr is not a pointer to a struct.
func ParseTag(
	structPtr any,
	tagNames ...string,
) (map[string]StructFields, error) {
	return parseTagRecursive(structPtr, tagNames, "")
}

// parseTagRecursive is a helper function for ParseTag.
// It recursively parses the tags of the struct pointed to by structPtr and any nested structs.
// The prefix parameter is used to build the fully qualified names of the fields.
func parseTagRecursive(structPtr any, tagNames []string, prefix string) (map[string]StructFields, error) {
	valueOf := reflect.ValueOf(structPtr)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	var result = map[string]StructFields{}
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Type().Field(i)

		if field.Type.Kind() == reflect.Struct {
			ptr := valueOf.Field(i).Addr().Interface()
			newPrefix := fmt.Sprintf("%s%s.", prefix, field.Name)
			recursionValue, _ := parseTagRecursive(ptr, tagNames, newPrefix)
			for k, v := range recursionValue {
				result[k] = v
			}
			continue
		}

		for _, tagName := range tagNames {
			if tag := field.Tag.Get(tagName); tag != "" {
				fieldName := fmt.Sprintf("%s%s", prefix, field.Name)
				result[fieldName] = StructFields{
					FieldName: fieldName,
					TagName:   tagName,
					TagValue:  tag,
				}
			}
		}
	}

	return result, nil
}
