package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFirefox(t *testing.T) {
	Convey("Subject: #NewFirefox", t, func() {
		Convey("It should return a new Firefox instance", func() {
			So(NewFirefox(NewUAParser("")), ShouldHaveSameTypeAs, &Firefox{})
		})
	})
}

func TestFirefoxName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the name", func() {
			So(NewFirefox(NewUAParser("")).Name(), ShouldEqual, "Firefox")
		})
	})
}

func TestFirefoxVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is captured", func() {
			ua := testUserAgents["firefox"]
			Convey("It should return the version", func() {
				So(NewFirefox(NewUAParser(ua.Android)).Version(), ShouldEqual, "59.0.2")
				So(NewFirefox(NewUAParser(ua.IOS)).Version(), ShouldEqual, "7.0.4")
				So(NewFirefox(NewUAParser(ua.Linux)).Version(), ShouldEqual, "3.6.8")
				So(NewFirefox(NewUAParser(ua.Mac)).Version(), ShouldEqual, "64.0")
				So(NewFirefox(NewUAParser(ua.Windows)).Version(), ShouldEqual, "3.6.8")
			})
		})

		Convey("When the version is not captured", func() {
			ua := testUserAgents["chrome"]
			Convey("It should return default version", func() {
				So(NewFirefox(NewUAParser(ua.Android)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestFirefoxMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			ua := testUserAgents["firefox"]
			Convey("It should return true", func() {
				So(NewFirefox(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewFirefox(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
				So(NewFirefox(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewFirefox(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
				So(NewFirefox(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			ua := testUserAgents["chrome"]
			Convey("It should return false", func() {
				So(NewFirefox(NewUAParser(ua.Android)).Match(), ShouldBeFalse)
			})
		})
	})
}
