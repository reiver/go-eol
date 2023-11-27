package eol

import (
	"io"
)

func readthisrune(runescanner io.RuneScanner, expected rune) (size int, err error) {
	if nil == runescanner {
		return 0, errNilRuneScanner
	}

	var r rune
	{
		var err error

		r, size, err = runescanner.ReadRune()
		if nil != err {
			const runeNumber = 1
			return size, errProblemReadingRune(err, runeNumber)
		}
	}

	{
		actual := r

		if expected != actual {
			err := runescanner.UnreadRune()
			if nil != err {
				const runeNumber = 1
				return size, errProblemUnreadingRune(err, runeNumber, r)
			}

			return 0, internalNotFoundError{expected: expected, actual: r}
		}
	}

	return size, nil
}
