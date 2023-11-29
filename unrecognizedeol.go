package eol

import (
	"fmt"
)

var _ error = internalUnrecognizedEOLError{}

type internalUnrecognizedEOLError struct {
	value string
}

func errUnrecognizedEOL(value string) error {
	return internalUnrecognizedEOLError{
		value:value,
	}
}

func (receiver internalUnrecognizedEOLError) Error() string {
	return fmt.Sprintf("eol: %q is an unrecognized end-of-line (EOL) sequence", receiver.value)
}

func (receiver internalUnrecognizedEOLError) UnrecognizedEOL() string {
	return receiver.value
}
