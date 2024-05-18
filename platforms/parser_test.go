package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParserMatching(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		userAgent := "Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Mobile Safari/537.36"
		Convey("When platform is matched", func() {
			Convey("It should return true", func() {
				parser := NewUAParser(userAgent)
				So(parser.Match([]string{"Android"}), ShouldBeTrue)
			})
		})

		Convey("When platform is not matched", func() {
			Convey("It should return false", func() {
				parser := NewUAParser(userAgent)
				So(parser.Match([]string{"Windows"}), ShouldBeFalse)
			})
		})
	})
}

func TestParserVersion(t *testing.T) {
	Convey("Subject: #version", t, func() {
		androidUserAgent := "Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Mobile Safari/537.36"
		Convey("When version is matched", func() {
			Convey("When order is 1", func() {
				Convey("It should return version", func() {
					parser := NewUAParser(androidUserAgent)
					So(parser.Version([]string{`Android ([\d.]+)`}, 1, ""), ShouldEqual, "4.4.2")
				})
			})

			Convey("When order is 2", func() {
				snapUserAgent := "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Mobile/15E148 Snapchat/12.33.0.36 (like Safari/8615.1.26.100.1, panda)"
				Convey("It should return version", func() {
					parser := NewUAParser(snapUserAgent)
					So(parser.Version([]string{`Snapchat( ?|/)([\d.]+)`}, 2, ""), ShouldEqual, "12.33.0.36")
				})
			})
		})

		Convey("When version is not matched", func() {
			Convey("It should return default version", func() {
				parser := NewUAParser(androidUserAgent)
				So(parser.Version([]string{`Android ([\d.]+)`}, 2, "0"), ShouldEqual, "0")
			})
		})
	})
}
