package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSafari(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		ua := testUserAgents["safari"]
		So(NewSafari(ua.Mac), ShouldHaveSameTypeAs, &Safari{})
	})
}

func TestSafariName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Safari", func() {
			ua := testUserAgents["safari"]
			So(NewSafari(ua.Mac).Name(), ShouldEqual, safariName)
		})
	})
}

func TestSafariVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["safari"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewSafari(ua.Mac).Version(), ShouldEqual, "8.0.7")
				So(NewSafari(ua.IOS).Version(), ShouldEqual, "8.0")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewSafari("").Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestSafariMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Safari", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["safari"]

				So(NewSafari(ua.Mac).Match(), ShouldBeTrue)
				So(NewSafari(ua.IOS).Match(), ShouldBeTrue)
			})
		})
	})
}
