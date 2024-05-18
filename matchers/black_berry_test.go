package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBlackberry(t *testing.T) {
	Convey("Subject: #NewBlackberry", t, func() {
		Convey("It should return a new Blackberry instance", func() {
			So(NewBlackBerry(NewUAParser("")), ShouldHaveSameTypeAs, &BlackBerry{})
		})
	})
}

func TestBlackberryName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the name", func() {
			So(NewBlackBerry(NewUAParser("")).Name(), ShouldEqual, "BlackBerry")
		})
	})
}

func TestBlackberryVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["blackberry"]
		Convey("When the version is captured", func() {
			Convey("It should return the version", func() {
				So(NewBlackBerry(NewUAParser(ua.Linux)).Version(), ShouldEqual, "10.1.0.2342")
				So(NewBlackBerry(NewUAParser(ua.Windows)).Version(), ShouldEqual, "7.1.0.346")
			})
		})

		Convey("When the version is not captured", func() {
			Convey("It should return default version", func() {
				So(NewBlackBerry(NewUAParser(ua.Android)).Version(), ShouldEqual, "0.0")
				So(NewBlackBerry(NewUAParser(ua.IOS)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestBlackberryMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["blackberry"]

				So(NewBlackBerry(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewBlackBerry(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
				So(NewBlackBerry(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewBlackBerry(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				bb := NewBlackBerry(
					NewUAParser("Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Mobile Safari/537.36"),
				)
				So(bb.Match(), ShouldBeFalse)
			})
		})
	})
}
