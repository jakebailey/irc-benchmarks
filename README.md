# irc-benchmark

This repo holds benchmarks comparing various IRC libraries, all of which implement IRCv3 tags.

- `jakebailey/irc`, my IRC library. In the encoding tests,
`WriteTo` indicates that the output was directly written to an `io.Writer`
rather than creating a temporary buffer.
- `jakebailey/ircold`, a fork of `MorpheusXAUT/irc`, which is a fork of
`github.com/sorcix/irc` (all to add IRCv3 tags).
- `github.com/sorcix/irc`, the `ircv3.2-tags` branch.
- `github.com/fluffle/goirc/client`, a handler-based IRC library.
- `github.com/thoj/go-ircevent`, an event-based IRC library. This doesn't
actually export the message parsing function, so `go-forceexport` was used
to find a function pointer rather than fork the project.

```
benchmark                                                         iter       time/iter   bytes alloc         allocs
---------                                                         ----       ---------   -----------         ------
BenchmarkParseSimple/github.com/jakebailey/irc-4              10000000    131.00 ns/op       32 B/op    1 allocs/op
BenchmarkParseSimple/github.com/jakebailey/ircold-4           10000000    225.00 ns/op      144 B/op    3 allocs/op
BenchmarkParseSimple/github.com/fluffle/goirc/client-4         5000000    366.00 ns/op      288 B/op    4 allocs/op
BenchmarkParseSimple/github.com/sorcix/irc-4                  10000000    221.00 ns/op      144 B/op    3 allocs/op
BenchmarkParseSimple/github.com/thoj/go-ircevent-4             5000000    352.00 ns/op      256 B/op    4 allocs/op
BenchmarkParseTwitch/github.com/jakebailey/irc-4               1000000   1355.00 ns/op     1234 B/op    3 allocs/op
BenchmarkParseTwitch/github.com/jakebailey/ircold-4             300000   5798.00 ns/op     4015 B/op   64 allocs/op
BenchmarkParseTwitch/github.com/fluffle/goirc/client-4          300000   5964.00 ns/op     4159 B/op   65 allocs/op
BenchmarkParseTwitch/github.com/sorcix/irc-4                    300000   5761.00 ns/op     4015 B/op   64 allocs/op
BenchmarkParseTwitch/github.com/thoj/go-ircevent-4              300000   4732.00 ns/op     3071 B/op   23 allocs/op
BenchmarkParseEscaping/github.com/jakebailey/irc-4              500000   2576.00 ns/op     1553 B/op    9 allocs/op
BenchmarkParseEscaping/github.com/jakebailey/ircold-4           200000   7977.00 ns/op     4877 B/op   84 allocs/op
BenchmarkParseEscaping/github.com/fluffle/goirc/client-4        200000   8082.00 ns/op     4958 B/op   84 allocs/op
BenchmarkParseEscaping/github.com/sorcix/irc-4                  200000   7885.00 ns/op     4877 B/op   84 allocs/op
BenchmarkParseEscaping/github.com/thoj/go-ircevent-4            200000   6523.00 ns/op     3549 B/op   31 allocs/op
BenchmarkEncodeSimple/github.com/jakebailey/irc-4             10000000    119.00 ns/op       48 B/op    1 allocs/op
BenchmarkEncodeSimple/github.com/jakebailey/irc_WriteTo-4     20000000    101.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeSimple/github.com/jakebailey/ircold-4          10000000    137.00 ns/op      112 B/op    1 allocs/op
BenchmarkEncodeSimple/github.com/sorcix/irc-4                 10000000    136.00 ns/op      112 B/op    1 allocs/op
BenchmarkEncodeTwitch/github.com/jakebailey/irc-4              2000000    920.00 ns/op      352 B/op    1 allocs/op
BenchmarkEncodeTwitch/github.com/jakebailey/irc_WriteTo-4      2000000    773.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeTwitch/github.com/jakebailey/ircold-4           1000000   1236.00 ns/op     1221 B/op    4 allocs/op
BenchmarkEncodeTwitch/github.com/sorcix/irc-4                  1000000   1259.00 ns/op     1219 B/op    4 allocs/op
BenchmarkEncodeEscaping/github.com/jakebailey/irc-4            1000000   1394.00 ns/op      480 B/op    1 allocs/op
BenchmarkEncodeEscaping/github.com/jakebailey/irc_WriteTo-4    1000000   1277.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeEscaping/github.com/jakebailey/ircold-4         1000000   1425.00 ns/op     1400 B/op    4 allocs/op
BenchmarkEncodeEscaping/github.com/sorcix/irc-4                1000000   1417.00 ns/op     1463 B/op    4 allocs/op
```

`github.com/fluffle/goirc/client` and `github.com/thoj/go-ircevent` are omitted from the encoding tests,
as they do not support encoding of their message types, only sending strings (with or without helpers
for common commands).


Formatted using `prettybench`: https://github.com/cespare/prettybench