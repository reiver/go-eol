package eol

import (
	"io"
)

func readthisrune(runescanner io.RuneScanner, expected rune, characterNumber uint64) (size int, err error) {
	if nil == runescanner {
		return 0, errNilRuneScanner
	}

	var r rune
	{
		var err error

		r, size, err = runescanner.ReadRune()
		if nil != err {
			return size, errProblemReadingRune(err, characterNumber)
		}
	}

	{
		actual := r

		if expected != actual {
			err := runescanner.UnreadRune()
			if nil != err {
				return size, errProblemUnreadingRune(err, characterNumber, r)
			}

			return 0, internalNotFoundError{expected: expected, actual: r, characterNumber:characterNumber}
		}
	}

	return size, nil
}
