package eol_test

import (
	"testing"

	"io"
	"strings"

	"sourcecode.social/reiver/go-utf8"

	"sourcecode.social/reiver/go-eol"
)

func TestReadCRLF(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedSize int
	}{
		{
			Value: "\r\n",
			ExpectedSize: 2,
		},



		{
			Value: "\r\napple banana cherry",
			ExpectedSize: 2,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualSize, err := eol.ReadCRLF(runescanner)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			expected := test.ExpectedSize
			actual   := actualSize

			if expected != actual {
				t.Errorf("For test #%d, the actual size is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

	}
}

func TestReadCRLF_fail(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedError string
		ExpectedSize int
	}{
		{
			Value: "",
			ExpectedError: `eol: problem reading character №1 of end-of-line sequence "\r\n": EOF`,
		},



		{
			Value: "\n",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\n' (U+000A)`,
		},
		{
			Value: "\r",
			ExpectedError: `eol: problem reading character №2 of end-of-line sequence "\r\n": EOF`,
			ExpectedSize: 1,
		},
		{
			Value: "\u0085",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\u0085' (U+0085)`,
		},
		{
			Value: "\u2028",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\u2028' (U+2028)`,
		},



		{
			Value: "😈",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '😈' (U+1F608)`,
		},



		{
			Value: "\napple banana cherry",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\n' (U+000A)`,
		},
		{
			Value: "\u0085apple banana cherry",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\u0085' (U+0085)`,
		},
		{
			Value: "\u2028apple banana cherry",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\u2028' (U+2028)`,
		},



		{
			Value: "😈apple banana cherry",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '😈' (U+1F608)`,
		},



		{
			Value: " \n",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found ' ' (U+0020)`,
		},
		{
			Value: " \r",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found ' ' (U+0020)`,
		},
		{
			Value: " \u0085",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found ' ' (U+0020)`,
		},
		{
			Value: " \u2028",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found ' ' (U+0020)`,
		},



		{
			Value: " 😈",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found ' ' (U+0020)`,
		},



		{
			Value: ".\n",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '.' (U+002E)`,
		},
		{
			Value: ".\r",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '.' (U+002E)`,
		},
		{
			Value: ".\u0085",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '.' (U+002E)`,
		},
		{
			Value: ".\u2028",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '.' (U+002E)`,
		},



		{
			Value: ".😈",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '.' (U+002E)`,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualSize, err := eol.ReadCRLF(runescanner)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("EXPECTED-ERROR: %q", test.ExpectedError)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			expected := test.ExpectedError
			actual   := err.Error()

			if expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
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
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
