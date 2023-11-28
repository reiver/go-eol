package eol

import (
	"io"

	"sourcecode.social/reiver/go-opt"
)

func ReadCRLF(runescanner io.RuneScanner) (size int, err error) {

	var size0 int
	{
		var err error

		const characterNumber uint64 = 1
		var circumstance internalCircumstance = specifyCircumstance(opt.Something(CRLF), characterNumber)
		size0, err = readthisrune(circumstance, runescanner, cr)
		if nil != err {
			return size0, err
		}
	}

	var size1 int
	{
		var err error

		const characterNumber uint64 = 2
		var circumstance internalCircumstance = specifyCircumstance(opt.Something(CRLF), characterNumber)
		size1, err = readthisrune(circumstance, runescanner, lf)
		if nil != err {
			return size1+size0, err
		}
	}

	return size1+size0, nil
}
