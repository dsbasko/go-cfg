package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseTag(t *testing.T) {
	type InStructFld struct {
		FldInt int `testTag:"field-int"`
	}
	type InStruct struct {
		FldString string `testTag:"field-string"`
		FldStruct InStructFld
	}

	tableTests := []struct {
		name       string
		tagName    string
		structPtr  any
		wantStruct map[string]StructFields
		wantErr    error
	}{
		{
			name:      "Happy Path",
			tagName:   "testTag",
			structPtr: &InStruct{},
			wantStruct: map[string]StructFields{
				"FldString": {
					FieldName: "FldString",
					TagName:   "testTag",
					TagValue:  "field-string",
				},
				"FldStruct.FldInt": {
					FieldName: "FldStruct.FldInt",
					TagName:   "testTag",
					TagValue:  "field-int",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			parsedData, err := ParseTag(tt.structPtr, tt.tagName)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, parsedData, tt.wantStruct)
		})
	}
}
