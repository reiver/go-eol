package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"
)

// ReadLF tries to read the "\n" (i.e., line-feed) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\n").
//
// If the character read is not a '\n', then ReadLF will try to unread the character.
//
// Example usage:
//
//	size, err := eol.ReadLF(runescanner)
func ReadLF(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(LF), characterNumber)
	return readthisrune(circumstance, runescanner, lf)
}
