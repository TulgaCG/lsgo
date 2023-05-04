package file_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TulgaCG/lsgo/pkg/file"
)

func TestList(t *testing.T) {
	tests := []struct {
		id          int
		description string
		input       []string
		flags       file.ListOpts
		wantError   bool
	}{
		{
			id:          0,
			description: "Without flag",
			input:       []string{"./testdata"},
		},
		{
			id:          1,
			description: "With flag",
			input:       []string{"./testdata"},
			flags: file.ListOpts{
				List:       false,
				ShowHidden: true,
			},
		},
		{
			id:          2,
			description: "non-existing path",
			input:       []string{"./testfile"},
			wantError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual := bytes.NewBufferString("")
			file.List(actual, test.input, test.flags)

			expected, err := os.ReadFile("./testdata/golden/goldendata" + fmt.Sprint(test.id))
			require.NoError(t, err)

			require.Equal(t, string(expected), actual.String())
		})
	}
}
