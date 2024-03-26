# go-eol

Package **eol** implements tools for working with end-of-line, for the Go programming language.

The end-of-line sequences it supports is:

* `"\n"     // line-feed           (LF)`
* `"\n\r"   // line-feed           (LF), carriage-return (CR)`
* `"\v"     // vertical-tab        (VT)`
* `"\f"     // form-feed           (FF)`
* `"\r"     // carriage-return     (CR)`
* `"\r\n"   // carriage-return     (CR), line-feed (LF)`
* `"\u0085" // next-line           (NEL)`
* `"\u2028" // line-separator      (LS)`
* `"\u2029" // paragraph-separator (PS)`

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-eol

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-eol?status.svg)](https://godoc.org/sourcecode.social/reiver/go-eol)

## Example

Here is an example:
```golang

import "sourcecode.social/reiver/go-eol"

// ...

eodofline, size, err := eol.ReadEOL(runereader)
if nil != err {
	return err
}
```

## Import

To import package **eol** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-eol"
```

## Installation

To install package **eol** do the following:
```
GOPROXY=direct go get https://sourcecode.social/reiver/go-eol
```

## Author

Package **eol** was written by [Charles Iliya Krempeaux](http://changelog.ca)
