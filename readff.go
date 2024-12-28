package eol

import (
	"io"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-eol/ff"
)

// ReadFF tries to read the "\f" (i.e., form-feed) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\f").
//
// If the character read is not a '\f', then ReadFF will try to unread the character.
//
// Example usage:
//
//	size, err := eol.ReadFF(runescanner)
func ReadFF(runescanner io.RuneScanner) (size int, err error) {
	const characterNumber uint64 = 1
	var circumstance internalCircumstance = specifyCircumstance(opt.Something(ff.String), characterNumber)
	return readthisrune(circumstance, runescanner, ff.Rune)
}
