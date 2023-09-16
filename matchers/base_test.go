package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBaseMatch(t *testing.T) {
	Convey("Subject: #match", t, func() {
		Convey("When user agent matches", func() {
			Convey("It should return true", func() {
				b := newBase("Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Mobile Safari/537.36")
				So(b.match([]string{`Android`}), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				b := newBase("Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Mobile Safari/537.36")
				So(b.match([]string{`Konqueror`}), ShouldBeFalse)
			})
		})
	})
}

func TestBaseVersion(t *testing.T) {
	Convey("Subject: #version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				b := newBase("Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Mobile Safari/537.36")
				So(b.version([]string{`Chrome/(\d+)\.(\d+)`}, 1), ShouldEqual, "87")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return 0.0", func() {
				b := newBase("Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko)")
				So(b.version([]string{`Chrome/(\d+)\.(\d+)`}, 1), ShouldEqual, "0.0")
			})
		})
	})
}
