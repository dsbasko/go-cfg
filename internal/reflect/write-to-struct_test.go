package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WriteToStruct(t *testing.T) {
	type InStructFld struct {
		FldBool bool `testTag:"field-bool"`
	}
	type InStruct struct {
		FldString  string  `testTag:"field-string"`
		FldInt     int     `testTag:"field-int"`
		FldInt8    int8    `testTag:"field-int8"`
		FldInt16   int16   `testTag:"field-int16"`
		FldInt32   int32   `testTag:"field-int32"`
		FldInt64   int64   `testTag:"field-int64"`
		FldUint    uint    `testTag:"field-uint"`
		FldUint8   uint8   `testTag:"field-uint8"`
		FldUint16  uint16  `testTag:"field-uint16"`
		FldUint32  uint32  `testTag:"field-uint32"`
		FldUint64  uint64  `testTag:"field-uint64"`
		FldFloat32 float32 `testTag:"field-float32"`
		FldFloat64 float64 `testTag:"field-float64"`
		FldStruct  InStructFld
	}
	wantStruct := InStruct{
		FldString:  "allons-y",
		FldInt:     -42,
		FldInt8:    -8,
		FldInt16:   -16,
		FldInt32:   -32,
		FldInt64:   -64,
		FldUint:    42,
		FldUint8:   8,
		FldUint16:  16,
		FldUint32:  32,
		FldUint64:  64,
		FldFloat32: 32.32,
		FldFloat64: 64.64,
		FldStruct: InStructFld{
			FldBool: true,
		},
	}

	tableTests := []struct {
		name      string
		structPtr any
		fn        func(fieldName string) string
		wantErr   error
	}{
		{
			name:      "Happy Path",
			structPtr: &InStruct{},
			fn: func(fieldName string) string {
				mockData := map[string]string{
					"FldString":         "allons-y",
					"FldInt":            "-42",
					"FldInt8":           "-8",
					"FldInt16":          "-16",
					"FldInt32":          "-32",
					"FldInt64":          "-64",
					"FldUint":           "42",
					"FldUint8":          "8",
					"FldUint16":         "16",
					"FldUint32":         "32",
					"FldUint64":         "64",
					"FldFloat32":        "32.32",
					"FldFloat64":        "64.64",
					"FldStruct.FldBool": "true",
				}
				return mockData[fieldName]
			},
			wantErr: nil,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteToStruct(tt.structPtr, tt.fn)

			if err != nil && tt.wantErr != nil {
				assert.Equal(t, err, tt.wantErr)
			} else {
				assert.Equal(t, &wantStruct, tt.structPtr)
			}
		})
	}
}
