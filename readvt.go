package eol

import (
	"io"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-eol/vt"
)

// ReadVT tries to read the "\v" (i.e., carriage-return) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\v").
//
// If the character read is not a '\v', then ReadVT will try to unread the character.
//
// Example usage:
//
//	size, err := eol.ReadVT(runescanner)
func ReadVT(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(vt.String), characterNumber)
	return readthisrune(circumstance, runescanner, vt.Rune)
}
