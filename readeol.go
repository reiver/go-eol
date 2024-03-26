package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/crlf"
	"sourcecode.social/reiver/go-eol/ff"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/lfcr"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
	"sourcecode.social/reiver/go-eol/vt"
)

// ReadEOL tries to read an end-of-line sequence.
//
// The end-of-line sequences it supports are:
//
//	line-feed       (LF)  (U+000A) ('\n')
//	line-feed, carriage-return     ("\n\r")
//	vertical-tab    (VT)  (U+000B) ('\v')
//	vertical-tab    (VT)  (U+000B) ('\f')
//	carriage-return (CR)  (U+000D) ('\r')
//	carriage-return, line-feed     ("\r\n")
//	next-line       (NEL) (U+0085)
//	line-separator  (LS)  (U+2028)
//
// If successful, ReadEOL return the end-of-line sequence it found and the number-of-bytes read (to read in end-of-line sequence it found).
//
// Example usage:
///
//	eolSequence, size, err := eol.ReadEOL(runescanner)
func ReadEOL(runescanner io.RuneScanner) (endofline string, size int, err error) {
	if nil == runescanner {
		return "", 0, errNilRuneScanner
	}

	var r0 rune
	var size0 int
	{
		var err error

		r0, size0, err = runescanner.ReadRune()
		if nil != err && size0 <= 0 {
			const characterNumber uint64 = 1
			var eolSequence opt.Optional[string] // = opt.Nothing[string]() // i.e., unknown
			var circumstance internalCircumstance = specifyCircumstance(eolSequence, characterNumber)

			return "", size0, errProblemReadingRune(circumstance, err)
		}
	}

	switch r0 {
	case lf.Rune:
		// Nothing here.
	case vt.Rune:
		return vt.String, size0, nil
	case ff.Rune:
		return ff.String, size0, nil
	case cr.Rune:
		// Nothing here.
	case nel.Rune:
		return nel.String, size0, nil
	case ls.Rune:
		return ls.String, size0, nil
	default:
		err := runescanner.UnreadRune()
		if nil != err {
			const characterNumber uint64 = 1
			var eolSequence opt.Optional[string] // = opt.Nothing[string]() // i.e., unknown
			var circumstance internalCircumstance = specifyCircumstance(eolSequence, characterNumber)

			return "", size0, errProblemUnreadingRune(circumstance, err, r0)
		}

		return "", 0, errNotEOL(r0)
	}

	// if we got here, then we had a LR or CR

	var r1 rune
	var size1 int
	{
		var err error

		r1, size1, err = runescanner.ReadRune()
		if io.EOF == err {
			return string(r0), size0, nil
		}
		if nil != err && size1 <= 0 {
			const characterNumber uint64 = 2
			var eolSequence opt.Optional[string] // = opt.Nothing[string]() // i.e., unknown
			var circumstance internalCircumstance = specifyCircumstance(eolSequence, characterNumber)

			return "", size1+size0, errProblemReadingRune(circumstance, err)
		}
	}

	switch {
	case cr.Rune == r0 && lf.Rune == r1:
		return crlf.String, size1+size0, nil
	case lf.Rune == r0 && cr.Rune == r1:
		return lfcr.String, size1+size0, nil
	default:
		err := runescanner.UnreadRune()
		if nil != err {
			const characterNumber uint64 = 2
			var eolSequence opt.Optional[string] // = opt.Nothing[string]() // i.e., unknown
			var circumstance internalCircumstance = specifyCircumstance(eolSequence, characterNumber)

			return "", size1+size0, errProblemUnreadingRune(circumstance, err, r1)
		}

		return string(r0), size0, nil
	}
}
