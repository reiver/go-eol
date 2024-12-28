package eol_test

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-utf8"

	"github.com/reiver/go-eol"
	"github.com/reiver/go-eol/cr"
	"github.com/reiver/go-eol/crlf"
	"github.com/reiver/go-eol/ff"
	"github.com/reiver/go-eol/lf"
	"github.com/reiver/go-eol/lfcr"
	"github.com/reiver/go-eol/ls"
	"github.com/reiver/go-eol/nel"
	"github.com/reiver/go-eol/ps"
	"github.com/reiver/go-eol/vt"
)

func TestReadEOL(t *testing.T) {

	tests := []struct{
		Value    string
		ExpectedEOL string
		ExpectedSize int
	}{
		{
			Value: "\n",
			ExpectedEOL: lf.String,
			ExpectedSize: 1,
		},
		{
			Value: "\n\r",
			ExpectedEOL: lfcr.String,
			ExpectedSize: 2,
		},
		{
			Value: "\v",
			ExpectedEOL: vt.String,
			ExpectedSize: 1,
		},
		{
			Value: "\f",
			ExpectedEOL: ff.String,
			ExpectedSize: 1,
		},
		{
			Value: "\r",
			ExpectedEOL: cr.String,
			ExpectedSize: 1,
		},
		{
			Value: "\r\n",
			ExpectedEOL: crlf.String,
			ExpectedSize: 2,
		},
		{
			Value: "\u0085",
			ExpectedEOL: nel.String,
			ExpectedSize: 2,
		},
		{
			Value: "\u2028",
			ExpectedEOL: ls.String,
			ExpectedSize: 3,
		},
		{
			Value: "\u2029",
			ExpectedEOL: ps.String,
			ExpectedSize: 3,
		},



		{
			Value: "\napple banana cherry",
			ExpectedEOL: lf.String,
			ExpectedSize: 1,
		},
		{
			Value: "\n\rapple banana cherry",
			ExpectedEOL: lfcr.String,
			ExpectedSize: 2,
		},
		{
			Value: "\vapple banana cherry",
			ExpectedEOL: vt.String,
			ExpectedSize: 1,
		},
		{
			Value: "\fapple banana cherry",
			ExpectedEOL: ff.String,
			ExpectedSize: 1,
		},
		{
			Value: "\rapple banana cherry",
			ExpectedEOL: cr.String,
			ExpectedSize: 1,
		},
		{
			Value: "\r\napple banana cherry",
			ExpectedEOL: crlf.String,
			ExpectedSize: 2,
		},
		{
			Value: "\u0085apple banana cherry",
			ExpectedEOL: nel.String,
			ExpectedSize: 2,
		},
		{
			Value: "\u2028apple banana cherry",
			ExpectedEOL: ls.String,
			ExpectedSize: 3,
		},
		{
			Value: "\u2029apple banana cherry",
			ExpectedEOL: ps.String,
			ExpectedSize: 3,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualEOL, actualSize, err := eol.ReadEOL(runescanner)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE:        %q", test.Value)
			t.Logf("EXPECTED-EOL: %q", test.ExpectedEOL)
			t.Logf("EXPECTED-SIZE: %d", test.ExpectedSize)
			continue
		}

		{
			expected := test.ExpectedEOL
			actual   := actualEOL

			if expected != actual {
				t.Errorf("For test #%d, the actual end-of-line sequence is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:    %q", test.Value)
				t.Logf("EXPECTED-SIZE: %d", test.ExpectedSize)
				continue
			}
		}

		{
			expected := test.ExpectedSize
			actual   := actualSize

			if expected != actual {
				t.Errorf("For test #%d, the actual size is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("VALUE:        %q", test.Value)
				t.Logf("EXPECTED-EOL: %q", test.ExpectedEOL)
				continue
			}
		}
	}
}

func TestReadEOL_fail(t *testing.T) {

	tests := []struct{
		Value    string
		ExpectedEOL string
		ExpectedSize int
		ExpectedError string
	}{
		{
			Value:               "apple",
			ExpectedError: "eol: 'a' (U+0061) is not an end-of-line character",
		},
		{
			Value:               "banana",
			ExpectedError: "eol: 'b' (U+0062) is not an end-of-line character",
		},
		{
			Value:               "cherry",
			ExpectedError: "eol: 'c' (U+0063) is not an end-of-line character",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualEOL, actualSize, err := eol.ReadEOL(runescanner)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one." , testNumber)
			t.Logf("EXPECTED-ERROR: %q", test.ExpectedError)
			t.Logf("VALUE:        %q", test.Value)
			t.Logf("EXPECTED-EOL: %q", test.ExpectedEOL)
			t.Logf("EXPECTED-SIZE: %d", test.ExpectedSize)
			continue
		}

		{
			expected := test.ExpectedError
			actual   := err.Error()

			if expected != actual {
				t.Errorf("For test %d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:        %q", test.Value)
				t.Logf("EXPECTED-EOL: %q", test.ExpectedEOL)
				t.Logf("EXPECTED-SIZE: %d", test.ExpectedSize)
				continue
			}
		}

		{
			expected := test.ExpectedEOL
			actual   := actualEOL

			if expected != actual {
				t.Errorf("For test %d, the actual end-of-line sequence is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:    %q", test.Value)
				t.Logf("EXPECTED-SIZE: %d", test.ExpectedSize)
				t.Logf("EXPECTED-ERROR: %q", test.ExpectedError)
				continue
			}
		}

		{
			expected := test.ExpectedSize
			actual   := actualSize

			if expected != actual {
				t.Errorf("For test %d, the actual size is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("VALUE:        %q", test.Value)
				t.Logf("EXPECTED-EOL: %q", test.ExpectedEOL)
				t.Logf("EXPECTED-ERROR: %q", test.ExpectedError)
				continue
			}
		}
	}
}
