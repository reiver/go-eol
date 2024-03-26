package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"

	"sourcecode.social/reiver/go-eol/ps"
)

// ReadPS tries to read the "\r" (i.e., carriage-return) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\r").
//
// If the character read is not a '\r', then ReadPS will try to unread the character.
//
// Example usage:
//
//	size, err := eol.ReadPS(runescanner)
func ReadPS(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(ps.String), characterNumber)
	return readthisrune(circumstance, runescanner, ps.Rune)
}
