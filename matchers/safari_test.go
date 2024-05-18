package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSafari(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		ua := testUserAgents["safari"]
		So(NewSafari(NewUAParser(ua.Mac)), ShouldHaveSameTypeAs, &Safari{})
	})
}

func TestSafariName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Safari", func() {
			ua := testUserAgents["safari"]
			So(NewSafari(NewUAParser(ua.Mac)).Name(), ShouldEqual, safariName)
		})
	})
}

func TestSafariVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["safari"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewSafari(NewUAParser(ua.Mac)).Version(), ShouldEqual, "8.0.7")
				So(NewSafari(NewUAParser(ua.IOS)).Version(), ShouldEqual, "8.0")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewSafari(NewUAParser("")).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestSafariMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Safari", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["safari"]

				So(NewSafari(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
				So(NewSafari(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
			})
		})
	})
}
