package eol

import (
	"github.com/reiver/go-opt"
)

// internalCircumstance is primarily used to help to create informative error messages.
//
// internalCircumstance contains the end-of-line sequence and the character-number within that end-of-line sequence.
//
// Example end-of-line sequences are:
//
//	• "\n"
//	• "\n\r"
//	• "\v"
//	• "\f"
//	• "\r"
//	• "\r\n"
//	• "\u0085"
//	• "\u2028"
//	• "\u2029"
//
// No end-of-line sequence can also be specified if it is unknown.
// For example eol.ReadEOL() does NOT know the end-of-line sequence ahead of time.
//
// The character-number is 1-indexed.
// Meaning that the 1st character is 1 (and not 0).
//
// Creating an internalCircumstance usually looks something like these:
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Nothing[string](), 1)
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Something]("\n"), 1)
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Something]("\v"), 1)
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Something]("\f"), 1)
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Something]("\r"), 1)
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Something]("\r\n"), 1)
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Something]("\r\n"), 2)
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Something]("\u0085"), 1)
//
//	var circumstance internalCircumstance = specifyCircumstance(opt.Something]("\u2028"), 1)
//
// These would then be passed to some type of error.
// For example:
//
//	var err error = errNotFound(circumstance, expected, actual)
//
//	var err error = errProblemReadingRune(circumstance, err)
//
//	var err error = errProblemUnreadingRune(circumstance, err, r)
type internalCircumstance struct {
	eolSequence opt.Optional[string]
	characterNumber uint64
}

func specifyCircumstance(eolSequence opt.Optional[string], characterNumber uint64) internalCircumstance {
	return internalCircumstance{
		eolSequence:eolSequence,
		characterNumber:characterNumber,
	}
}

func (receiver internalCircumstance) CharacterNumber() uint64 {
	return receiver.characterNumber
}

func (receiver internalCircumstance) EOLSequence() opt.Optional[string] {
	return receiver.eolSequence
}
