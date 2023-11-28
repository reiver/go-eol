package eol

import (
	"io"
)

func readthisrune(circumstance internalCircumstance, runescanner io.RuneScanner, expected rune) (size int, err error) {
	if nil == runescanner {
		return 0, errNilRuneScanner
	}

	var r rune
	{
		var err error

		r, size, err = runescanner.ReadRune()
		if nil != err {
			return size, errProblemReadingRune(circumstance, err)
		}
	}

	{
		actual := r

		if expected != actual {
			err := runescanner.UnreadRune()
			if nil != err {
				return size, errProblemUnreadingRune(circumstance, err, r)
			}

			return 0, errNotFound(circumstance, expected, actual)
		}
	}

	return size, nil
}
