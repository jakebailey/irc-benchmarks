package ircbench

import (
	"testing"

	fluffle "github.com/fluffle/goirc/client"
	twitch "github.com/gempir/go-twitch-irc"
	"github.com/goshuirc/irc-go/ircmsg"
	"github.com/jakebailey/irc"
	jbirc "github.com/jakebailey/ircold"
	sorcix "github.com/sorcix/irc"
)

func BenchmarkParseSimple(b *testing.B) {
	var raw = `:jake!jake@jake PRIVMSG #jake :Hello, world!`
	runParserBenchmarks(b, raw)
}

func BenchmarkParseTwitch(b *testing.B) {
	var raw = `@badges=subscriber/36,premium/1;color=#FF4C05;display-name=Zikaeroh;emote-only=1;emotes=488738:0-6,8-14,16-22,24-30;flags=;id=fb949c9d-bf74-429f-bf10-70fd1e767c1f;mod=0;room-id=16678946;subscriber=1;tmi-sent-ts=1538528941376;turbo=0;user-id=44527297;user-type= :zikaeroh!zikaeroh@zikaeroh.tmi.twitch.tv PRIVMSG #coestar :coeHype coeHype coeHype coeHype`
	runParserBenchmarks(b, raw)
}

func BenchmarkParseEscaping(b *testing.B) {
	var raw = `@badges=subscriber/48,bits/100;color=#9AB01C;display-name=Epyo;emotes=;flags=;id=bc3ced4f-0569-4b2f-bf9c-32d496829a74;login=epyo;mod=0;msg-id=resub;msg-param-months=53;msg-param-sub-plan-name=Channel\sSubscription\s(coestar);msg-param-sub-plan=1000;room-id=16678946;subscriber=1;system-msg=Epyo\sjust\ssubscribed\swith\sa\sTier\s1\ssub.\sEpyo\ssubscribed\sfor\s53\smonths\sin\sa\srow!;tmi-sent-ts=1538528813049;turbo=0;user-id=20874297;user-type= :tmi.twitch.tv USERNOTICE #coestar`
	runParserBenchmarks(b, raw)
}

func runParserBenchmarks(b *testing.B, raw string) {
	b.Run("github.com/jakebailey/irc", func(b *testing.B) {
		var m irc.Message
		for i := 0; i < b.N; i++ {
			m.Parse(raw) //nolint:errcheck
		}
	})

	b.Run("github.com/jakebailey/ircold", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jbirc.ParseMessage(raw)
		}
	})

	b.Run("github.com/fluffle/goirc/client", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fluffle.ParseLine(raw)
		}
	})

	b.Run("github.com/sorcix/irc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sorcix.ParseMessage(raw)
		}
	})

	b.Run("github.com/thoj/go-ircevent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ircEventParseToEvent(raw) //nolint:errcheck
		}
	})

	b.Run("github.com/goshuirc/irc-go/ircmsg", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ircmsg.ParseLine(raw)
		}
	})

	b.Run("github.com/gempir/go-twitch-irc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			twitch.ParseMessage(raw)
		}
	})
}
