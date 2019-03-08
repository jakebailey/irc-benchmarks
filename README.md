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

```
benchmark                                              iter       time/iter   bytes alloc         allocs
---------                                              ----       ---------   -----------         ------
BenchmarkParseSimple/jakebailey/irc-4              20000000    102.00 ns/op       16 B/op    1 allocs/op
BenchmarkParseSimple/jakebailey/ircold-4            5000000    221.00 ns/op      144 B/op    3 allocs/op
BenchmarkParseSimple/fluffle/goirc/client-4         5000000    348.00 ns/op      288 B/op    4 allocs/op
BenchmarkParseSimple/sorcix/irc-4                  10000000    206.00 ns/op      144 B/op    3 allocs/op
BenchmarkParseSimple/thoj/go-ircevent-4             5000000    339.00 ns/op      256 B/op    4 allocs/op
BenchmarkParseSimple/goshuirc/irc-go/ircmsg-4      10000000    161.00 ns/op       48 B/op    2 allocs/op
BenchmarkParseSimple/gempir/go-twitch-irc-4        10000000    192.00 ns/op      448 B/op    3 allocs/op

BenchmarkParseTwitch/jakebailey/irc-4               1000000   1270.00 ns/op     1218 B/op    3 allocs/op
BenchmarkParseTwitch/jakebailey/ircold-4             300000   5661.00 ns/op     4015 B/op   64 allocs/op
BenchmarkParseTwitch/fluffle/goirc/client-4          300000   5849.00 ns/op     4159 B/op   65 allocs/op
BenchmarkParseTwitch/sorcix/irc-4                    300000   5700.00 ns/op     4015 B/op   64 allocs/op
BenchmarkParseTwitch/thoj/go-ircevent-4              300000   4535.00 ns/op     3071 B/op   23 allocs/op
BenchmarkParseTwitch/goshuirc/irc-go/ircmsg-4       1000000   2222.00 ns/op     2191 B/op    6 allocs/op
BenchmarkParseTwitch/gempir/go-twitch-irc-4          300000   5562.00 ns/op     3879 B/op   35 allocs/op

BenchmarkParseEscaping/jakebailey/irc-4              500000   2544.00 ns/op     1553 B/op    9 allocs/op
BenchmarkParseEscaping/jakebailey/ircold-4           200000   7807.00 ns/op     4878 B/op   84 allocs/op
BenchmarkParseEscaping/fluffle/goirc/client-4        200000   7915.00 ns/op     4957 B/op   84 allocs/op
BenchmarkParseEscaping/sorcix/irc-4                  200000   7817.00 ns/op     4878 B/op   84 allocs/op
BenchmarkParseEscaping/thoj/go-ircevent-4            200000   6294.00 ns/op     3550 B/op   31 allocs/op
BenchmarkParseEscaping/goshuirc/irc-go/ircmsg-4      500000   3083.00 ns/op     2462 B/op    9 allocs/op
BenchmarkParseEscaping/gempir/go-twitch-irc-4        300000   5242.00 ns/op     3678 B/op   30 allocs/op

BenchmarkEncodeSimple/jakebailey/irc-4             10000000    119.00 ns/op       48 B/op    1 allocs/op
BenchmarkEncodeSimple/jakebailey/irc_WriteTo-4     20000000     99.20 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeSimple/jakebailey/ircold-4          20000000    103.00 ns/op       64 B/op    1 allocs/op
BenchmarkEncodeSimple/sorcix/irc-4                 20000000    104.00 ns/op       64 B/op    1 allocs/op
BenchmarkEncodeSimple/goshuirc/irc-go/ircmsg-4     20000000    104.00 ns/op       64 B/op    1 allocs/op

BenchmarkEncodeTwitch/jakebailey/irc-4              2000000    895.00 ns/op      352 B/op    1 allocs/op
BenchmarkEncodeTwitch/jakebailey/irc_WriteTo-4      2000000    800.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeTwitch/jakebailey/ircold-4           1000000   1192.00 ns/op     1184 B/op    3 allocs/op
BenchmarkEncodeTwitch/sorcix/irc-4                  1000000   1219.00 ns/op     1151 B/op    3 allocs/op
BenchmarkEncodeTwitch/goshuirc/irc-go/ircmsg-4      1000000   1489.00 ns/op     1118 B/op    3 allocs/op

BenchmarkEncodeEscaping/jakebailey/irc-4            1000000   1456.00 ns/op      480 B/op    1 allocs/op
BenchmarkEncodeEscaping/jakebailey/irc_WriteTo-4    1000000   1324.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeEscaping/jakebailey/ircold-4         1000000   1347.00 ns/op     1335 B/op    4 allocs/op
BenchmarkEncodeEscaping/sorcix/irc-4                1000000   1387.00 ns/op     1426 B/op    4 allocs/op
BenchmarkEncodeEscaping/goshuirc/irc-go/ircmsg-4    1000000   2202.00 ns/op     1584 B/op    7 allocs/op
```

`github.com/fluffle/goirc/client`, `github.com/thoj/go-ircevent` and `github.com/gemir/go-twitch-irc`
are omitted from the encoding tests,
as they do not support encoding of their message types, only sending strings (with or without helpers
for common commands).


Formatted using [prettybench](https://github.com/cespare/prettybench).