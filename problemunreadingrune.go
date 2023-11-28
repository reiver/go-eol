package eol

import (
	"fmt"
)

var _ error = internalProblemUnreadingRuneError{}

func errProblemUnreadingRune(err error, runeNumber uint64, r rune) error {
	return internalProblemUnreadingRuneError{
		err:err,
		runeNumber:runeNumber,
		r:r,
	}
}

type internalProblemUnreadingRuneError struct {
	err error
	runeNumber uint64
	r rune
}

func (receiver internalProblemUnreadingRuneError) Error() string {
	err := receiver.err
	runeNumber := receiver.runeNumber
	r := receiver.r

	return fmt.Sprintf("eol: problem unreading character №%d (%q (%U)) of end-of-line sequence: %s", runeNumber, r, r, err)
}

func (receiver internalProblemUnreadingRuneError) Unwrap() error {
	return receiver.err
}
