# irc-benchmark

This repo holds benchmarks comparing various IRC libraries, all of which implement IRCv3 tags.

- [jakebailey/irc](https://github.com/jakebailey/irc), my IRC library. In the encoding tests,
`WriteTo` indicates that the output was directly written to an `io.Writer`
rather than creating a temporary buffer.
- [jakebailey/ircold](https://github.com/jakebailey/ircold), a fork of `MorpheusXAUT/irc`, which is a fork of
`github.com/sorcix/irc` (all to add IRCv3 tags).
- [sorcix/irc](https://github.com/sorcix/irc), the `ircv3.2-tags` branch.
- [fluffle/goirc/client](https://github.com/fluffle/goirc/client), a handler-based IRC library.
- [thoj/go-ircevent](https://github.com/thoj/go-ircevent), an event-based IRC library. This doesn't
actually export the message parsing function, so `go-forceexport` was used
to find a function pointer rather than fork the project.
- [goshuirc/irc-go/ircmsg](https://github.com/goshuirc/irc-go/ircmsg), an IRC message handler lib.
- [gempir/go-twitch-irc](https://github.com/gempir/go-twitch-irc), an IRC lib specifically for Twitch. This
comparison is likely unfair because it includes extra functionailty around
extracting tags.

Running `go test -bench . -benchmem | prettybench`:

```
benchmark                                              iter       time/iter   bytes alloc         allocs
---------                                              ----       ---------   -----------         ------
BenchmarkParseSimple/jakebailey/irc-4              20000000    102.00 ns/op       16 B/op    1 allocs/op
BenchmarkParseSimple/jakebailey/ircold-4           10000000    213.00 ns/op      144 B/op    3 allocs/op
BenchmarkParseSimple/fluffle/goirc/client-4         5000000    363.00 ns/op      288 B/op    4 allocs/op
BenchmarkParseSimple/sorcix/irc-4                  10000000    211.00 ns/op      144 B/op    3 allocs/op
BenchmarkParseSimple/thoj/go-ircevent-4             5000000    341.00 ns/op      256 B/op    4 allocs/op
BenchmarkParseSimple/goshuirc/irc-go/ircmsg-4      10000000    168.00 ns/op       48 B/op    2 allocs/op
BenchmarkParseSimple/gempir/go-twitch-irc-4        10000000    193.00 ns/op      448 B/op    3 allocs/op

BenchmarkParseTwitch/jakebailey/irc-4               1000000   1268.00 ns/op     1218 B/op    3 allocs/op
BenchmarkParseTwitch/jakebailey/ircold-4             300000   5765.00 ns/op     4015 B/op   64 allocs/op
BenchmarkParseTwitch/fluffle/goirc/client-4          300000   5843.00 ns/op     4159 B/op   65 allocs/op
BenchmarkParseTwitch/sorcix/irc-4                    300000   5633.00 ns/op     4015 B/op   64 allocs/op
BenchmarkParseTwitch/thoj/go-ircevent-4              300000   4551.00 ns/op     3071 B/op   23 allocs/op
BenchmarkParseTwitch/goshuirc/irc-go/ircmsg-4       1000000   2236.00 ns/op     2191 B/op    6 allocs/op
BenchmarkParseTwitch/gempir/go-twitch-irc-4          300000   5691.00 ns/op     3879 B/op   35 allocs/op

BenchmarkParseEscaping/jakebailey/irc-4              500000   2538.00 ns/op     1552 B/op    9 allocs/op
BenchmarkParseEscaping/jakebailey/ircold-4           200000   7812.00 ns/op     4878 B/op   84 allocs/op
BenchmarkParseEscaping/fluffle/goirc/client-4        200000   7966.00 ns/op     4958 B/op   84 allocs/op
BenchmarkParseEscaping/sorcix/irc-4                  200000   7842.00 ns/op     4878 B/op   84 allocs/op
BenchmarkParseEscaping/thoj/go-ircevent-4            200000   6373.00 ns/op     3549 B/op   31 allocs/op
BenchmarkParseEscaping/goshuirc/irc-go/ircmsg-4      500000   3108.00 ns/op     2462 B/op    9 allocs/op
BenchmarkParseEscaping/gempir/go-twitch-irc-4        300000   5388.00 ns/op     3678 B/op   30 allocs/op

BenchmarkEncodeSimple/jakebailey/irc-4             10000000    117.00 ns/op       48 B/op    1 allocs/op
BenchmarkEncodeSimple/jakebailey/irc_WriteTo-4     20000000    101.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeSimple/jakebailey/ircold-4          20000000    101.00 ns/op       64 B/op    1 allocs/op
BenchmarkEncodeSimple/sorcix/irc-4                 20000000    102.00 ns/op       64 B/op    1 allocs/op
BenchmarkEncodeSimple/goshuirc/irc-go/ircmsg-4     20000000    105.00 ns/op       64 B/op    1 allocs/op

BenchmarkEncodeTwitch/jakebailey/irc-4              2000000    863.00 ns/op      352 B/op    1 allocs/op
BenchmarkEncodeTwitch/jakebailey/irc_WriteTo-4      2000000    782.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeTwitch/jakebailey/ircold-4           1000000   1158.00 ns/op     1221 B/op    4 allocs/op
BenchmarkEncodeTwitch/sorcix/irc-4                  1000000   1139.00 ns/op     1057 B/op    3 allocs/op
BenchmarkEncodeTwitch/goshuirc/irc-go/ircmsg-4      1000000   1497.00 ns/op     1161 B/op    4 allocs/op

BenchmarkEncodeEscaping/jakebailey/irc-4            1000000   1421.00 ns/op      480 B/op    1 allocs/op
BenchmarkEncodeEscaping/jakebailey/irc_WriteTo-4    1000000   1304.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeEscaping/jakebailey/ircold-4         1000000   1336.00 ns/op     1381 B/op    4 allocs/op
BenchmarkEncodeEscaping/sorcix/irc-4                1000000   1342.00 ns/op     1290 B/op    4 allocs/op
BenchmarkEncodeEscaping/goshuirc/irc-go/ircmsg-4    1000000   2170.00 ns/op     1677 B/op    8 allocs/op
```

`github.com/fluffle/goirc/client`, `github.com/thoj/go-ircevent` and `github.com/gemir/go-twitch-irc`
are omitted from the encoding tests,
as they do not support encoding of their message types, only sending strings (with or without helpers
for common commands).


Formatted using [prettybench](https://github.com/cespare/prettybench).