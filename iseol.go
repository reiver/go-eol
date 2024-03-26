package eol

import (
	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/ff"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
	"sourcecode.social/reiver/go-eol/ps"
	"sourcecode.social/reiver/go-eol/vt"
)

func IsEOL(r rune) bool {
	switch r {
	case cr.Rune, ff.Rune, lf.Rune, ls.Rune, nel.Rune, ps.Rune, vt.Rune:
		return true
	default:
		return false
	}
}
