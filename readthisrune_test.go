package eol

import (
	"testing"

	"io"
	"strings"
	"sourcecode.social/reiver/go-utf8"
)

func TestReadThisRune(t *testing.T) {

	tests := []struct{
		Value string
		Rune rune
		ExpectedSize int
	}{
		{
			Value: "\n",
			Rune:  '\n',
			ExpectedSize: 1,
		},
		{
			Value: "\r",
			Rune:  '\r',
			ExpectedSize: 1,
		},
		{
			Value: "\u0085",
			Rune:  '\u0085',
			ExpectedSize: 2,
		},
		{
			Value: "\u2028",
			Rune:  '\u2028',
			ExpectedSize: 3,
		},



		{
			Value: "😈",
			Rune:  '😈',
			ExpectedSize: 4,
		},



		{
			Value: "\napple banana cherry",
			Rune:  '\n',
			ExpectedSize: 1,
		},
		{
			Value: "\rapple banana cherry",
			Rune:  '\r',
			ExpectedSize: 1,
		},
		{
			Value: "\u0085apple banana cherry",
			Rune:  '\u0085',
			ExpectedSize: 2,
		},
		{
			Value: "\u2028apple banana cherry",
			Rune:  '\u2028',
			ExpectedSize: 3,
		},



		{
			Value: "😈apple banana cherry",
			Rune:  '😈',
			ExpectedSize: 4,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		const runeNumber = 999
		actualSize, err := readthisrune(runescanner, test.Rune, runeNumber)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("RUNE: %q", test.Rune)
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
				t.Logf("RUNE: %q", test.Rune)
				continue
			}
		}

	}
}

func TestReadThisRune_fail(t *testing.T) {

	tests := []struct{
		Value string
		Rune rune
		RuneNumber uint64
		ExpectedError string
	}{
		{
			Value: "",
			Rune:  '\n',
			RuneNumber: 7,
			ExpectedError: `eol: problem reading character №7 of end-of-line sequence: EOF`,
		},
		{
			Value: "",
			Rune:  '\r',
			RuneNumber: 8,
			ExpectedError: `eol: problem reading character №8 of end-of-line sequence: EOF`,
		},
		{
			Value: "",
			Rune:  '\u0085',
			RuneNumber: 9,
			ExpectedError: `eol: problem reading character №9 of end-of-line sequence: EOF`,
		},
		{
			Value: "",
			Rune:  '\u2028',
			RuneNumber: 10,
			ExpectedError: `eol: problem reading character №10 of end-of-line sequence: EOF`,
		},



		{
			Value: "",
			Rune:  '😈',
			RuneNumber: 11,
			ExpectedError: `eol: problem reading character №11 of end-of-line sequence: EOF`,
		},



		{
			Value: " \n",
			Rune:  '\n',
			RuneNumber: 12,
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line character №12 — instead found ' ' (U+0020)`,
		},
		{
			Value: " \r",
			Rune:  '\r',
			RuneNumber: 13,
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line character №13 — instead found ' ' (U+0020)`,
		},
		{
			Value: " \u0085",
			Rune:  '\u0085',
			RuneNumber: 14,
			ExpectedError: `eol: next-line (NEL) character (U+0085) not found for end-of-line character №14 — instead found ' ' (U+0020)`,
		},
		{
			Value: " \u2028",
			Rune:  '\u2028',
			RuneNumber: 15,
			ExpectedError: `eol: line-separator (LS) character (U+2028) not found for end-of-line character №15 — instead found ' ' (U+0020)`,
		},



		{
			Value: " 😈",
			Rune:  '😈',
			RuneNumber: 16,
			ExpectedError: `eol: '😈' character (U+1F608) not found for character №16 — instead found ' ' (U+0020)`,
		},



		{
			Value: ".\n",
			Rune:  '\n',
			RuneNumber: 17,
			ExpectedError: `eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line character №17 — instead found '.' (U+002E)`,
		},
		{
			Value: ".\r",
			Rune:  '\r',
			RuneNumber: 18,
			ExpectedError: `eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line character №18 — instead found '.' (U+002E)`,
		},
		{
			Value: ".\u0085",
			Rune:  '\u0085',
			RuneNumber: 19,
			ExpectedError: `eol: next-line (NEL) character (U+0085) not found for end-of-line character №19 — instead found '.' (U+002E)`,
		},
		{
			Value: ".\u2028",
			Rune:  '\u2028',
			RuneNumber: 20,
			ExpectedError: `eol: line-separator (LS) character (U+2028) not found for end-of-line character №20 — instead found '.' (U+002E)`,
		},



		{
			Value: ".😈",
			Rune:  '😈',
			RuneNumber: 21,
			ExpectedError: `eol: '😈' character (U+1F608) not found for character №21 — instead found '.' (U+002E)`,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualSize, err := readthisrune(runescanner, test.Rune, test.RuneNumber)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("EXPECTED-ERROR: %q", test.ExpectedError)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("RUNE: %q", test.Rune)
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
				t.Logf("RUNE: %q", test.Rune)
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
				t.Logf("RUNE: %q", test.Rune)
				continue
			}
		}
	}
}
