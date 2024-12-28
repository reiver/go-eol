package eol

import (
	"io"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-eol/cr"
)

// ReadCR tries to read the "\r" (i.e., carriage-return) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\r").
//
// If the character read is not a '\r', then ReadCR will try to unread the character.
//
// Example usage:
//
//	size, err := eol.ReadCR(runescanner)
func ReadCR(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(cr.String), characterNumber)
	return readthisrune(circumstance, runescanner, cr.Rune)
}
