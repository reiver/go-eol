package eol

import (
	"github.com/reiver/go-eol/cr"
	"github.com/reiver/go-eol/ff"
	"github.com/reiver/go-eol/lf"
	"github.com/reiver/go-eol/ls"
	"github.com/reiver/go-eol/nel"
	"github.com/reiver/go-eol/ps"
	"github.com/reiver/go-eol/vt"
)

func IsEOL(r rune) bool {
	switch r {
	case cr.Rune, ff.Rune, lf.Rune, ls.Rune, nel.Rune, ps.Rune, vt.Rune:
		return true
	default:
		return false
	}
}
