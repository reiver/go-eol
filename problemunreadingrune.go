package eol

import (
	"fmt"
)

var _ error = internalProblemUnreadingRuneError{}

func errProblemUnreadingRune(circumstance internalCircumstance, err error, r rune) error {
	return internalProblemUnreadingRuneError{
		err:err,
		r:r,
		circumstance:circumstance,
	}
}

type internalProblemUnreadingRuneError struct {
	err error
	r rune
	circumstance internalCircumstance
}

func (receiver internalProblemUnreadingRuneError) Error() string {
	err := receiver.err
	r := receiver.r
	characterNumber := receiver.circumstance.CharacterNumber()
	eolSequence := receiver.circumstance.EOLSequence()

	var s string = fmt.Sprintf("eol: problem unreading character №%d (%q (%U)) of end-of-line sequence: %s", characterNumber, r, r, err)
	eolSequence.WhenSomething(func(sequence string){
		s    = fmt.Sprintf("eol: problem unreading character №%d (%q (%U)) of end-of-line sequence %q: %s", characterNumber, r, r, sequence, err)
	})

	return s
}

func (receiver internalProblemUnreadingRuneError) Unwrap() error {
	return receiver.err
}
