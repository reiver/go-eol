package eol_test

import (
	"testing"

	"io"
	"strings"

	"sourcecode.social/reiver/go-utf8"

	"sourcecode.social/reiver/go-eol"
)

func TestReadLF(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedSize int
	}{
		{
			Value: "\n",
			ExpectedSize: 1,
		},



		{
			Value: "\napple banana cherry",
			ExpectedSize: 1,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualSize, err := eol.ReadLF(runescanner)
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

func TestReadLF_fail(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedError string
	}{
		{
			Value: "",
			ExpectedError: `eol: problem reading rune №1 of end-of-line sequence: EOF`,
		},



		{
			Value: "\r",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '\r' (U+000D)`,
		},
		{
			Value: "\u0085",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '\u0085' (U+0085)`,
		},
		{
			Value: "\u2028",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '\u2028' (U+2028)`,
		},



		{
			Value: "😈",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '😈' (U+1F608)`,
		},



		{
			Value: "\rapple banana cherry",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '\r' (U+000D)`,
		},
		{
			Value: "\u0085apple banana cherry",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '\u0085' (U+0085)`,
		},
		{
			Value: "\u2028apple banana cherry",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '\u2028' (U+2028)`,
		},



		{
			Value: "😈apple banana cherry",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '😈' (U+1F608)`,
		},



		{
			Value: " \n",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found ' ' (U+0020)`,
		},
		{
			Value: " \r",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found ' ' (U+0020)`,
		},
		{
			Value: " \u0085",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found ' ' (U+0020)`,
		},
		{
			Value: " \u2028",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found ' ' (U+0020)`,
		},



		{
			Value: " 😈",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found ' ' (U+0020)`,
		},



		{
			Value: ".\n",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '.' (U+002E)`,
		},
		{
			Value: ".\r",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '.' (U+002E)`,
		},
		{
			Value: ".\u0085",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '.' (U+002E)`,
		},
		{
			Value: ".\u2028",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '.' (U+002E)`,
		},



		{
			Value: ".😈",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found — instead found '.' (U+002E)`,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualSize, err := eol.ReadLF(runescanner)
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
			expected := 0
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
