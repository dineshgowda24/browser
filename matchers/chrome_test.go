package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewChrome(t *testing.T) {
	Convey("Subject: #NewChrome", t, func() {
		Convey("It should return a new Chrome instance", func() {
			So(NewChrome(NewUAParser("")), ShouldHaveSameTypeAs, &Chrome{})
		})
	})
}

func TestChromeName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the name", func() {
			So(NewChrome(NewUAParser("")).Name(), ShouldEqual, "Chrome")
		})
	})
}

func TestChromeVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is captured", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["chrome"]
				So(NewChrome(NewUAParser(ua.Windows)).Version(), ShouldEqual, "44.0.2403.155")
				So(NewChrome(NewUAParser(ua.Linux)).Version(), ShouldEqual, "42.0.2311.135")
				So(NewChrome(NewUAParser(ua.Mac)).Version(), ShouldEqual, "44.0.2403.155")
				So(NewChrome(NewUAParser(ua.IOS)).Version(), ShouldEqual, "44.0.2403.67")
				So(NewChrome(NewUAParser(ua.Android)).Version(), ShouldEqual, "30.0.0.0")
			})
		})

		Convey("When the version is not captured", func() {
			Convey("It should return default version", func() {
				So(NewChrome(NewUAParser("")).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestChromeMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["chrome"]
				So(NewChrome(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewChrome(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewChrome(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
				So(NewChrome(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
				So(NewChrome(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["opera-mini"]
				So(NewChrome(NewUAParser(ua.Mac)).Match(), ShouldBeFalse)
			})
		})
	})
}
