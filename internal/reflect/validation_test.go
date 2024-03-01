package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validation(t *testing.T) {
	type InStruct struct{}
	InString := ""

	tableTests := []struct {
		name      string
		tagName   string
		structPtr any
		wantErr   error
	}{
		{
			name:      "Happy Path",
			tagName:   "testTag",
			structPtr: &InStruct{},
			wantErr:   nil,
		},
		{
			name:      "Nil",
			tagName:   "testTag",
			structPtr: nil,
			wantErr:   ErrNil,
		},
		{
			name:      "Not Pointer",
			tagName:   "testTag",
			structPtr: InStruct{},
			wantErr:   ErrNotPointer,
		},
		{
			name:      "Not Struct",
			tagName:   "testTag",
			structPtr: &InString,
			wantErr:   ErrNotStruct,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validation(tt.structPtr)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
