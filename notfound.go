package eol

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
)

var _ error = internalNotFoundError{}

type internalNotFoundError struct{
	expected rune
	actual rune
	circumstance internalCircumstance
}

func errNotFound(circumstance internalCircumstance, expected rune, actual rune) error {
	return internalNotFoundError {
		expected:expected,
		actual:actual,
		circumstance:circumstance,
	}
}

func (receiver internalNotFoundError) Error() string {


	var buffer [256]byte
	var p []byte = buffer[0:0]

	var expected rune = receiver.expected
	var actual   rune = receiver.actual

	var characterNumber uint64           = receiver.circumstance.CharacterNumber()
	var eolSequence opt.Optional[string] = receiver.circumstance.EOLSequence()

	switch expected {
	case lf.Rune:
		var s string = fmt.Sprintf(`eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line sequence character №%d — instead found %q (%U)`, characterNumber, actual, actual)
		eolSequence.WhenSomething(func(sequence string){
			s    = fmt.Sprintf(`eol: line-feed (LF) character ('\n') (U+000A) not found for end-of-line sequence %q character №%d — instead found %q (%U)`, sequence, characterNumber, actual, actual)
		})
		p = append(p, s...)
	case cr.Rune:
			var s string = fmt.Sprintf(`eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence character №%d — instead found %q (%U)`, characterNumber, actual, actual)
		eolSequence.WhenSomething(func(sequence string){
			s            = fmt.Sprintf(`eol: carriage-return (CR) character ('\r') (U+000D) not found for end-of-line sequence %q character №%d — instead found %q (%U)`, sequence, characterNumber, actual, actual)
		})
		p = append(p, s...)
	case nel.Rune:
		var s string = fmt.Sprintf(`eol: next-line (NEL) character (U+0085) not found for end-of-line sequence character №%d — instead found %q (%U)`, characterNumber, actual, actual)
		eolSequence.WhenSomething(func(sequence string){
			s    = fmt.Sprintf(`eol: next-line (NEL) character (U+0085) not found for end-of-line sequence %q character №%d — instead found %q (%U)`, sequence, characterNumber, actual, actual)
		})
		p = append(p, s...)
	case ls.Rune:
		var s string = fmt.Sprintf(`eol: line-separator (LS) character (U+2028) not found for end-of-line sequence character №%d — instead found %q (%U)`, characterNumber, actual, actual)
		eolSequence.WhenSomething(func(sequence string){
			s    = fmt.Sprintf(`eol: line-separator (LS) character (U+2028) not found for end-of-line sequence %q character №%d — instead found %q (%U)`, sequence, characterNumber, actual, actual)
		})
		p = append(p, s...)
	default:
		var s string = fmt.Sprintf(`eol: %q character (%U) not found for sequence character №%d — instead found %q (%U)`, expected, expected, characterNumber, actual, actual)
		eolSequence.WhenSomething(func(sequence string){
			s    = fmt.Sprintf(`eol: %q character (%U) not found for sequence %q character №%d — instead found %q (%U)`, expected, expected, sequence, characterNumber, actual, actual)
		})
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
