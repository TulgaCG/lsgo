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
			description: "Single Existing",
			input:       []string{"testdata"},
		},
		{
			id:          1,
			description: "Single Existing with Flag",
			input:       []string{"testdata"},
			flags: file.ListOpts{
				List:       false,
				ShowHidden: true,
			},
		},
		{
			id:          2,
			description: "Single Non-existing",
			input:       []string{"testfile"},
			wantError:   true,
		},
		{
			id:          3,
			description: "Multiple Existing",
			input:       []string{"testdata", "testdata/golden"},
		},
		{
			id:          4,
			description: "Multiple Existing with Flag",
			input:       []string{"testdata", "testdata/golden"},
			flags: file.ListOpts{
				List:       false,
				ShowHidden: true,
			},
		},
		{
			id:          5,
			description: "Multiple Non-existing",
			input:       []string{"testfile", "testdata/goldenfile"},
			wantError:   true,
		},
		{
			id:          6,
			description: "Existing and Non-existing",
			input:       []string{"testdata", "testfile"},
			wantError:   true,
		},
		{
			id:          7,
			description: "Existing and Non-existing with Flag",
			input:       []string{"testdata", "testfile"},
			wantError:   true,
			flags: file.ListOpts{
				List:       false,
				ShowHidden: true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual := bytes.NewBufferString("")
			file.List(actual, test.flags, test.input...)

			expected, err := os.ReadFile(fmt.Sprintf("./testdata/golden/goldendata%d", test.id))
			require.NoError(t, err)

			require.Equal(t, string(expected), actual.String())
		})
	}
}
