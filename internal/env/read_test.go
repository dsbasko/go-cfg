package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dsbasko/go-cfg/internal/reflect"
)

func TestRead(t *testing.T) {
	type InStructNested struct {
		Field string `env:"NESTED_FIELD" default:"defValueNestedField"`
	}
	type InStruct struct {
		Field  string `env:"FIELD" default:"defValueField"`
		Nested InStructNested
	}

	tests := []struct {
		name       string
		osCfg      func()
		structPtr  any
		wantStruct *InStruct
		wantErr    error
	}{
		{
			name:      "Default",
			osCfg:     func() {},
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "defValueField",
				Nested: InStructNested{
					Field: "defValueNestedField",
				},
			},
			wantErr: nil,
		},
		{
			name: "Happy Path",
			osCfg: func() {
				_ = os.Setenv("FIELD", "fieldValue")
				_ = os.Setenv("NESTED_FIELD", "nestedFieldValue")
			},
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "fieldValue",
				Nested: InStructNested{
					Field: "nestedFieldValue",
				},
			},
			wantErr: nil,
		},
		{
			name:       "Validation Fail",
			osCfg:      func() {},
			structPtr:  InStruct{},
			wantStruct: &InStruct{},
			wantErr:    reflect.ErrNotPointer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.osCfg()
			err := Read(tt.structPtr)

			if err != nil || tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.EqualValues(t, tt.wantStruct, tt.structPtr)
			}

			os.Clearenv()
		})
	}
}
