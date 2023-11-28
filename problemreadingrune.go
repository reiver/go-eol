package eol

import (
	"fmt"
)

var _ error = internalProblemReadingRuneError{}

func errProblemReadingRune(circumstance internalCircumstance, err error) error {
	return internalProblemReadingRuneError{
		err:err,
		circumstance:circumstance,
	}
}

type internalProblemReadingRuneError struct {
	err error
	circumstance internalCircumstance
}

func (receiver internalProblemReadingRuneError) Error() string {
	err := receiver.err
	characterNumber := receiver.circumstance.CharacterNumber()
	eolSequence := receiver.circumstance.EOLSequence()

	var s string = fmt.Sprintf("eol: problem reading character №%d of end-of-line sequence: %s", characterNumber, err)
	eolSequence.WhenSomething(func(sequence string){
		s    = fmt.Sprintf("eol: problem reading character №%d of end-of-line sequence %q: %s", characterNumber, sequence, err)
	})

	return s
}

func (receiver internalProblemReadingRuneError) Unwrap() error {
	return receiver.err
}
