package ircbench

import (
	"io/ioutil"
	"testing"

	"github.com/goshuirc/irc-go/ircmsg"
	"github.com/jakebailey/irc"
	jbirc "github.com/jakebailey/ircold"
	sorcix "github.com/sorcix/irc"
)

func BenchmarkEncodeSimple(b *testing.B) {
	var raw = `:jake!jake@jake PRIVMSG #jake :Hello, world!`
	runEncoderBenchmarks(b, raw)
}

func BenchmarkEncodeTwitch(b *testing.B) {
	var raw = `@badges=subscriber/36,premium/1;color=#FF4C05;display-name=Zikaeroh;emote-only=1;emotes=488738:0-6,8-14,16-22,24-30;flags=;id=fb949c9d-bf74-429f-bf10-70fd1e767c1f;mod=0;room-id=16678946;subscriber=1;tmi-sent-ts=1538528941376;turbo=0;user-id=44527297;user-type= :zikaeroh!zikaeroh@zikaeroh.tmi.twitch.tv PRIVMSG #coestar :coeHype coeHype coeHype coeHype`
	runEncoderBenchmarks(b, raw)
}

func BenchmarkEncodeEscaping(b *testing.B) {
	var raw = `@badges=subscriber/48,bits/100;color=#9AB01C;display-name=Epyo;emotes=;flags=;id=bc3ced4f-0569-4b2f-bf9c-32d496829a74;login=epyo;mod=0;msg-id=resub;msg-param-months=53;msg-param-sub-plan-name=Channel\sSubscription\s(coestar);msg-param-sub-plan=1000;room-id=16678946;subscriber=1;system-msg=Epyo\sjust\ssubscribed\swith\sa\sTier\s1\ssub.\sEpyo\ssubscribed\sfor\s53\smonths\sin\sa\srow!;tmi-sent-ts=1538528813049;turbo=0;user-id=20874297;user-type= :tmi.twitch.tv USERNOTICE #coestar`
	runEncoderBenchmarks(b, raw)
}

func runEncoderBenchmarks(b *testing.B, raw string) {
	b.Run("jakebailey/irc", func(b *testing.B) {
		var m irc.Message
		m.Parse(raw) //nolint:errcheck

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Bytes()
		}
	})

	b.Run("jakebailey/irc WriteTo", func(b *testing.B) {
		var m irc.Message
		m.Parse(raw) //nolint:errcheck

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.WriteToWithNewline(ioutil.Discard) //nolint:errcheck
		}
	})

	b.Run("jakebailey/ircold", func(b *testing.B) {
		m := jbirc.ParseMessage(raw)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Bytes()
		}
	})

	b.Run("fluffle/goirc/client", func(b *testing.B) {
		b.Skip("does not support message reencoding")
	})

	b.Run("sorcix/irc", func(b *testing.B) {
		m := sorcix.ParseMessage(raw)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Bytes()
		}
	})

	b.Run("thoj/go-ircevent", func(b *testing.B) {
		b.Skip("does not support message reencoding")
	})

	b.Run("goshuirc/irc-go/ircmsg", func(b *testing.B) {
		m, _ := ircmsg.ParseLine(raw)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.LineBytes() //nolint:errcheck
		}
	})

	b.Run("gemir/go-twitch-irc", func(b *testing.B) {
		b.Skip("does not support message reencoding")
	})
}
