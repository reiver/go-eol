package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"

	"sourcecode.social/reiver/go-eol/nel"
)

// ReadNEL tries to read the "\u0085" (i.e., next-line) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\u0085").
//
// If the character read is not a '\u0085', then ReadNEL will try to unread the character.
//
// Example usage:
//
//	size, err := eol.ReadNEL(runescanner)
func ReadNEL(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(nel.String), characterNumber)
	return readthisrune(circumstance, runescanner, nel.Rune)
}
