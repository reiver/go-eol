package eol

import (
	"fmt"
)

var _ error = internalProblemReadingRuneError{}

func errProblemReadingRune(err error, runeNumber uint64) error {
	return internalProblemReadingRuneError{
		err:err,
		runeNumber:runeNumber,
	}
}

type internalProblemReadingRuneError struct {
	err error
	runeNumber uint64
}

func (receiver internalProblemReadingRuneError) Error() string {
	err := receiver.err
	runeNumber := receiver.runeNumber

	return fmt.Sprintf("eol: problem reading character №%d of end-of-line sequence: %s", runeNumber, err)
}

func (receiver internalProblemReadingRuneError) Unwrap() error {
	return receiver.err
}
