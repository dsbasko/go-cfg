package dflt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dsbasko/go-cfg/internal/reflect"
)

func Test_Read(t *testing.T) {
	type NestedStruct struct {
		Field string `default:"nestedFieldDefValue"`
	}
	type InStruct struct {
		Field  string `default:"fieldDefValue"`
		Nested NestedStruct
	}
	mockString := ""

	tableTests := []struct {
		name       string
		structPtr  any
		wantStruct *InStruct
		wantErr    error
	}{
		{
			name:      "Happy Path",
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "fieldDefValue",
				Nested: NestedStruct{
					Field: "nestedFieldDefValue",
				},
			},
			wantErr: nil,
		},
		{
			name:       "Validate error",
			structPtr:  &mockString,
			wantStruct: &InStruct{},
			wantErr:    reflect.ErrNotStruct,
		},
		{
			name:       "Parse error",
			structPtr:  InStruct{},
			wantStruct: &InStruct{},
			wantErr:    reflect.ErrNotPointer,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			err := Read(tt.structPtr)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				assert.EqualValues(t, tt.wantStruct, tt.structPtr)
			}
		})
	}
}
