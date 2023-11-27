package eol

import (
	"io"
)

// ReadEOL tries to read and end-of-line character.
//
// The end-of-line sequences it supports are:
//
//	line-feed       (LF)  (U+000A) ('\n')
//	carriage-return (CR)  (U+000D) ('\r')
//	carriage-return, line-feed    ("\r\n")
//	next-line       (NEL) (U+0085)
//	line-separator  (LS)  (U+2028)
//
// If successful, ReadEOL return the end-of-line sequence and the number-of-bytes read.
func ReadEOL(runescanner io.RuneScanner) (endofline string, size int, err error) {
	if nil == runescanner {
		return "", 0, errNilRuneScanner
	}

	var r0 rune
	var size0 int
	{
		var err error

		r0, size0, err = runescanner.ReadRune()
		if nil != err {
			const runeNumber = 1
			return "", size0, errProblemReadingRune(err, runeNumber)
		}
	}

	switch r0 {
	case lf:
		return LF, size0, nil
	case cr:
		// Nothing here.
	case nel:
		return NEL, size0, nil
	case ls:
		return LS, size0, nil
	default:
		err := runescanner.UnreadRune()
		if nil != err {
			const runeNumber = 1
			return "", size0, errProblemUnreadingRune(err, runeNumber, r0)
		}

		return "", 0, errNotEOL(r0)
	}

	// if we got here, then we had a CR

	var r1 rune
	var size1 int
	{
		var err error

		r1, size1, err = runescanner.ReadRune()
		if io.EOF == err {
			return CR, size0, nil
		}
		if nil != err {
			const runeNumber = 2
			return "", size1+size0, errProblemReadingRune(err, runeNumber)
		}
	}

	switch r1 {
	case lf:
		return CRLF, size1+size0, nil
	default:
		err := runescanner.UnreadRune()
		if nil != err {
			const runeNumber = 2
			return "", size1+size0, errProblemUnreadingRune(err, runeNumber, r1)
		}

		return CR, size0, nil
	}
}
