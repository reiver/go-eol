package eol

import (
	"io"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-eol/cr"
	"github.com/reiver/go-eol/crlf"
	"github.com/reiver/go-eol/lf"
)

// ReadCRLF tries to read the "\r\n" (i.e., carriage-return line-feed) end-of-line sequence.
//
// If successful, it returns the number-of-bytes read (to read in end-of-line sequence "\r\n").
//
// If the first character read is not a '\r', then ReadCRLF will try to unread the character.
// If the second character read is not a '\n', then ReadCRLF will also try to unread the second character, but will not be able to unread the first character (i.e., '\r') it already read.
//
// Example usage:
//
//	size, err := eol.ReadCRLF(runescanner)
func ReadCRLF(runescanner io.RuneScanner) (size int, err error) {

	var size0 int
	{
		var err error

		const characterNumber uint64 = 1
		var circumstance internalCircumstance = specifyCircumstance(opt.Something(crlf.String), characterNumber)
		size0, err = readthisrune(circumstance, runescanner, cr.Rune)
		if nil != err {
			return size0, err
		}
	}

	var size1 int
	{
		var err error

		const characterNumber uint64 = 2
		var circumstance internalCircumstance = specifyCircumstance(opt.Something(crlf.String), characterNumber)
		size1, err = readthisrune(circumstance, runescanner, lf.Rune)
		if nil != err {
			return size1+size0, err
		}
	}

	return size1+size0, nil
}
