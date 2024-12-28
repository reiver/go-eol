package eol_test

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-utf8"

	"github.com/reiver/go-eol"
)

func TestReadThisEOL(t *testing.T) {

	tests := []struct{
		Value string
		EOL string
	}{
		{
			Value: "\n",
			EOL:   "\n",
		},
		{
			Value: "\r",
			EOL:   "\r",
		},
		{
			Value: "\r\n",
			EOL:   "\r\n",
		},
		{
			Value: "\u0085",
			EOL:   "\u0085",
		},
		{
			Value: "\u2028",
			EOL:   "\u2028",
		},



		{
			Value: "\n"     + "12345",
			EOL:   "\n",
		},
		{
			Value: "\r"     + "12345",
			EOL:   "\r",
		},
		{
			Value: "\r\n"   + "12345",
			EOL:   "\r\n",
		},
		{
			Value: "\u0085" + "12345",
			EOL:   "\u0085",
		},
		{
			Value: "\u2028" + "12345",
			EOL:   "\u2028",
		},



		{
			Value: "\n"     + "12345" + "\n",
			EOL:   "\n",
		},
		{
			Value: "\r"     + "12345" + "\r",
			EOL:   "\r",
		},
		{
			Value: "\r\n"   + "12345" + "\r\n",
			EOL:   "\r\n",
		},
		{
			Value: "\u0085" + "12345" + "\u0085",
			EOL:   "\u0085",
		},
		{
			Value: "\u2028" + "12345" + "\u2028",
			EOL:   "\u2028",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualNumRead, err := eol.ReadThisEOL(runescanner, test.EOL)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("EOL:   %q", test.EOL)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			var expected int = len(test.EOL)
			var actual   int = actualNumRead

			if expected != actual {
				t.Errorf("For tst #%d, the actual number-of-bytes-read is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("EOL:   %q", test.EOL)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}

func TestReadThisEOL_fail(t *testing.T) {

	tests := []struct{
		Value string
		EOL string
		ExpectedError string
		ExpectedNumRead int
	}{
		{
			Value: "",
			EOL:   "\n",
			ExpectedError: `eol: problem reading character №1 of end-of-line sequence "\n": EOF`,
		},
		{
			Value: "",
			EOL:   "\r",
			ExpectedError: `eol: problem reading character №1 of end-of-line sequence "\r": EOF`,
		},
		{
			Value: "",
			EOL:   "\r\n",
			ExpectedError: `eol: problem reading character №1 of end-of-line sequence "\r\n": EOF`,
		},
		{
			Value: "",
			EOL:   "\u0085",
			ExpectedError: `eol: problem reading character №1 of end-of-line sequence "\u0085": EOF`,
		},
		{
			Value: "",
			EOL:   "\u2028",
			ExpectedError: `eol: problem reading character №1 of end-of-line sequence "\u2028": EOF`,
		},



		{
			Value: "\n",
			EOL:   "\r",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r" character №1 — instead found '\n' (U+000A)`,
		},
		{
			Value: "\n",
			EOL:   "\r\n",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\n' (U+000A)`,
		},
		{
			Value: "\n",
			EOL:   "\u0085",
			ExpectedError: `eol: next-line (NEL) character (U+0085) not found for end-of-line sequence "\u0085" character №1 — instead found '\n' (U+000A)`,
		},
		{
			Value: "\n",
			EOL:   "\u2028",
			ExpectedError: `eol: line-separator (LS) character (U+2028) not found for end-of-line sequence "\u2028" character №1 — instead found '\n' (U+000A)`,
		},



		{
			Value: "\r",
			EOL:   "\n",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line sequence "\n" character №1 — instead found '\r' (U+000D)`,
		},
		{
			Value: "\r",
			EOL:   "\r\n",
			ExpectedError: `eol: problem reading character №2 of end-of-line sequence "\r\n": EOF`,
			ExpectedNumRead: 1,
		},
		{
			Value: "\r",
			EOL:   "\u0085",
			ExpectedError: `eol: next-line (NEL) character (U+0085) not found for end-of-line sequence "\u0085" character №1 — instead found '\r' (U+000D)`,
		},
		{
			Value: "\r",
			EOL:   "\u2028",
			ExpectedError: `eol: line-separator (LS) character (U+2028) not found for end-of-line sequence "\u2028" character №1 — instead found '\r' (U+000D)`,
		},



		{
			Value: "\r\n",
			EOL:   "\n",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line sequence "\n" character №1 — instead found '\r' (U+000D)`,
		},
		{
			Value: "\r\n",
			EOL:   "\u0085",
			ExpectedError: `eol: next-line (NEL) character (U+0085) not found for end-of-line sequence "\u0085" character №1 — instead found '\r' (U+000D)`,
		},
		{
			Value: "\r\n",
			EOL:   "\u2028",
			ExpectedError: `eol: line-separator (LS) character (U+2028) not found for end-of-line sequence "\u2028" character №1 — instead found '\r' (U+000D)`,
		},



		{
			Value: "\u0085",
			EOL:   "\n",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line sequence "\n" character №1 — instead found '\u0085' (U+0085)`,
		},
		{
			Value: "\u0085",
			EOL:   "\r",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r" character №1 — instead found '\u0085' (U+0085)`,
		},
		{
			Value: "\u0085",
			EOL:   "\r\n",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\u0085' (U+0085)`,
		},
		{
			Value: "\u0085",
			EOL:   "\u2028",
			ExpectedError: `eol: line-separator (LS) character (U+2028) not found for end-of-line sequence "\u2028" character №1 — instead found '\u0085' (U+0085)`,
		},



		{
			Value: "\u2028",
			EOL:   "\n",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line sequence "\n" character №1 — instead found '\u2028' (U+2028)`,
		},
		{
			Value: "\u2028",
			EOL:   "\r",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r" character №1 — instead found '\u2028' (U+2028)`,
		},
		{
			Value: "\u2028",
			EOL:   "\r\n",
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence "\r\n" character №1 — instead found '\u2028' (U+2028)`,
		},
		{
			Value: "\u2028",
			EOL:   "\u0085",
			ExpectedError: `eol: next-line (NEL) character (U+0085) not found for end-of-line sequence "\u0085" character №1 — instead found '\u2028' (U+2028)`,
		},









		{
			Value: "\rapple banana cherry",
			EOL:   "\r\n",
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line sequence "\r\n" character №2 — instead found 'a' (U+0061)`,
			ExpectedNumRead: 1,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualNumRead, err := eol.ReadThisEOL(runescanner, test.EOL)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("EXPECTED-ERROR: %s", test.ExpectedError)
			t.Logf("EOL:   %q", test.EOL)
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
				t.Logf("EOL:   %q", test.EOL)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedNumRead
			actual   := actualNumRead

			if expected != actual {
				t.Errorf("For test #%d, the actual number-of-bytes-read is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("EOL:   %q", test.EOL)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
