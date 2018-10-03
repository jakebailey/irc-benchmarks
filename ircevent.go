package ircbench

import (
	forceexport "github.com/linux4life798/go-forceexport"
	ircevent "github.com/thoj/go-ircevent"
)

// ircevent's parsing function is unexported, so do a nasty hack to pull it out.
var ircEventParseToEvent func(msg string) (*ircevent.Event, error)

func init() {
	// prevent dead code elimination of parseToEvent
	ircevent.IRC("foo", "bar").Connect("") //nolint:errcheck

	if err := forceexport.GetFunc(&ircEventParseToEvent, "github.com/jakebailey/irc-benchmark/vendor/github.com/thoj/go-ircevent.parseToEvent"); err != nil {
		panic(err)
	}
}
