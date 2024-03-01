package flag

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dsbasko/go-cfg/internal/reflect"
)

func Test_Read(t *testing.T) {
	type InStructNested struct {
		Field string `flag:"nested-field" s-flag:"n" default:"defValueNestedField"`
	}
	type InStruct struct {
		Field  string `flag:"field" s-flag:"f" default:"defValueField"`
		Nested InStructNested
	}
	osArgs := os.Args

	tableTests := []struct {
		name       string
		osCfg      func()
		structPtr  any
		wantStruct InStruct
		wantErr    error
	}{
		{
			name: "Default",
			osCfg: func() {
				os.Args = osArgs
			},
			structPtr: &InStruct{},
			wantStruct: InStruct{
				Field: "defValueField",
				Nested: InStructNested{
					Field: "defValueNestedField",
				},
			},
			wantErr: nil,
		},
		{
			name: "Happy Path Full",
			osCfg: func() {
				os.Args = append( //nolint:gocritic
					osArgs,
					"--field=FULL_FIELD_VALUE",
					"--nested-field=FULL_NESTED_VALUE",
				)
			},
			structPtr: &InStruct{},
			wantStruct: InStruct{
				Field: "FULL_FIELD_VALUE",
				Nested: InStructNested{
					Field: "FULL_NESTED_VALUE",
				},
			},
			wantErr: nil,
		},
		{
			name: "Happy Path Short",
			osCfg: func() {
				os.Args = append( //nolint:gocritic
					osArgs,
					"-f=SHORT_FIELD_VALUE",
					"-n=SHORT_NESTED_VALUE",
				)
			},
			structPtr: &InStruct{},
			wantStruct: InStruct{
				Field: "SHORT_FIELD_VALUE",
				Nested: InStructNested{
					Field: "SHORT_NESTED_VALUE",
				},
			},
			wantErr: nil,
		},
		{
			name: "Happy Path All (Short Priority)",
			osCfg: func() {
				os.Args = append( //nolint:gocritic
					osArgs,
					"--field=FULL_FIELD_VALUE",
					"--nested-field=FULL_NESTED_VALUE",
					"-f=SHORT_FIELD_VALUE",
					"-n=SHORT_NESTED_VALUE",
				)
			},
			structPtr: &InStruct{},
			wantStruct: InStruct{
				Field: "SHORT_FIELD_VALUE",
				Nested: InStructNested{
					Field: "SHORT_NESTED_VALUE",
				},
			},
			wantErr: nil,
		},
		{
			name: "Validate Not Pointer",
			osCfg: func() {
				os.Args = osArgs
			},
			structPtr:  InStruct{},
			wantStruct: InStruct{},
			wantErr:    reflect.ErrNotPointer,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.osCfg()
			err := Read(tt.structPtr)

			if err != nil || tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.EqualValues(t, tt.wantStruct.Field, tt.structPtr.(*InStruct).Field)
				assert.EqualValues(t, tt.wantStruct.Nested, tt.structPtr.(*InStruct).Nested)
			}
		})
	}
}
