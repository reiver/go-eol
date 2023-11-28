package eol

import (
	"fmt"
)

var _ error = internalNotFoundError{}

type internalNotFoundError struct{
	expected rune
	actual rune
	characterNumber uint64
}

func (receiver internalNotFoundError) Error() string {


	var buffer [256]byte
	var p []byte = buffer[0:0]

	var expected rune = receiver.expected
	var actual   rune = receiver.actual

	var characterNumber uint64 = receiver.characterNumber

	switch expected {
	case lf:
		var s string =  fmt.Sprintf(`eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line character №%d — instead found %q (%U)`, characterNumber, actual, actual)
		p = append(p, s...)
	case cr:
		var s string = fmt.Sprintf(`eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line character №%d — instead found %q (%U)`, characterNumber, actual, actual)
		p = append(p, s...)
	case nel:
		var s string = fmt.Sprintf(`eol: next-line (NEL) character (U+0085) not found for end-of-line character №%d — instead found %q (%U)`, characterNumber, actual, actual)
		p = append(p, s...)
	case ls:
		var s string = fmt.Sprintf(`eol: line-separator (LS) character (U+2028) not found for end-of-line character №%d — instead found %q (%U)`, characterNumber, actual, actual)
		p = append(p, s...)
	default:
		var s string = fmt.Sprintf(`eol: %q character (%U) not found for character №%d — instead found %q (%U)`, expected, expected, characterNumber, actual, actual)
		p = append(p, s...)
	}

	return string(p)
}

func (receiver internalNotFoundError) Actual() rune {
	return receiver.actual
}

func (receiver internalNotFoundError) Expected() rune {
	return receiver.expected
}
