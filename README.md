# log (archived)

**This repository is archived and unmaintained. It is kept for history only.**

A levelled logger for [uniqush-push][uniqush-push], written in 2011 and last
changed in 2018. uniqush-push was its only real user, and it now has this code
in-tree at [`log`][adopted]: same `Logger` interface, one bug fixed.

## Use the copy in uniqush-push

```go
import "github.com/uniqush/uniqush-push/log"

logger := log.NewLogger(os.Stderr, "[uniqush] ", log.LevelInfo)
```

Differences from the code here:

- **`Fatal` and `Fatalf` format their arguments.** See below.
- The level constants follow Go's naming convention: `LOGLEVEL_INFO` is
  `LevelInfo`. Their values are written out rather than derived from `iota` after
  a negative constant, which had made `LOGLEVEL_FATAL` 1 and left an unused slot
  at index 0 of the level table.
- `MultiLogger` is gone. It had no users, and `Fatal` through it only ever
  reached the first logger anyway, because the standard library's `Fatal` exits.
- There are tests.

## The bug, if you have vendored this

`nullLoggerWrapper` — the discarding logger used for every level above the
configured one — forwarded its arguments to the standard library as a single
slice instead of expanding them:

```go
func (l *nullLoggerWrapper) Fatalf(format string, v ...interface{}) {
	l.inner.Fatalf(format, v)     // wants v...
}

func (l *nullLoggerWrapper) Fatal(v ...interface{}) {
	l.inner.Fatal(v)              // wants v...
}
```

So this:

```go
log.NewLogger(os.Stderr, "", log.LOGLEVEL_SILENT).
	Fatalf("cannot start: %v on port %d", "bind failed", 8080)
```

prints:

```
[Fatal] 2026/09/08 18:03:37 cannot start: [bind failed 8080] on port %!d(MISSING)
```

Only the fatal methods are affected, because they are the only ones the
discarding logger passes through — the point of that type is that a disabled
`Debugf` never reaches `Sprintf`, while a fatal still prints even when logging is
switched off.

Which is what makes it worth a section of its own. The bug is confined to the
case where the operator has silenced logging, and in that case the fatal line is
the *only* output they get when the process dies. In uniqush that meant `log=off`
in `uniqush.conf`.

`go vet` reports it:

```
log.go:96:2: missing ... in args forwarded to print-like function
log.go:100:2: missing ... in args forwarded to printf-like function
```

It went unreported for nine years because this repository had no CI, no `go.mod`,
and one user. That is why the code moved rather than being fixed here: fixing it
here would have meant first giving the repository a module and a release, to
serve nobody but uniqush.

The fix is `v...` in both methods.

License: Apache-2.0

[uniqush-push]: https://github.com/uniqush/uniqush-push
[adopted]: https://github.com/uniqush/uniqush-push/tree/master/log
