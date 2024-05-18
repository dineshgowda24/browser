package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInstagram(t *testing.T) {
	Convey("Subject: #NewInstagram", t, func() {
		Convey("It should return a new Instagram instance", func() {
			So(NewInstagram(NewUAParser("")), ShouldHaveSameTypeAs, &Instagram{})
		})
	})
}

func TestInstagramName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			ig := NewInstagram(NewUAParser(""))
			So(ig.Name(), ShouldEqual, instagramName)
		})
	})
}

func TestInstagramVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["instagram"]
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				So(NewInstagram(NewUAParser(ua.IOS)).Version(), ShouldEqual, "10.18.0")
				So(NewInstagram(NewUAParser(ua.Android)).Version(), ShouldEqual, "9.7.0")
				So(NewInstagram(NewUAParser(ua.Windows)).Version(), ShouldEqual, "10.8.1")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewInstagram(NewUAParser(ua.Mac)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestInstagramMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches", func() {
			ua := testUserAgents["instagram"]

			So(NewInstagram(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
			So(NewInstagram(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
			So(NewInstagram(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
			So(NewInstagram(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				userAgent := "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko)"
				So(NewInstagram(NewUAParser(userAgent)).Match(), ShouldBeFalse)
			})
		})
	})
}
