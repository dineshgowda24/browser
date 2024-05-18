package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUCBrowser(t *testing.T) {
	Convey("Subject: #NewUCBrowser", t, func() {
		Convey("It should return a new UCBrowser instance", func() {
			So(NewUCBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &UCBrowser{})
		})
	})
}

func TestUCBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return UCBrowser", func() {
			uc := NewUCBrowser(NewUAParser(""))
			So(uc.Name(), ShouldEqual, "UCBrowser")
		})
	})
}

func TestUCBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				uc := testUserAgents["uc-browser"]

				So(NewUCBrowser(NewUAParser(uc.Android)).Version(), ShouldEqual, "10.6.2.599")
				So(NewUCBrowser(NewUAParser(uc.IOS)).Version(), ShouldEqual, "10.6.5.627")
				So(NewUCBrowser(NewUAParser(uc.Linux)).Version(), ShouldEqual, "7.9.0.94")
				So(NewUCBrowser(NewUAParser(uc.Mac)).Version(), ShouldEqual, "3.2.0.417")
				So(NewUCBrowser(NewUAParser(uc.Windows)).Version(), ShouldEqual, "4.2.1.541")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				uc := NewUCBrowser(NewUAParser(testUserAgents["chrome"].Linux))
				So(uc.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestUCBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches UC Browser", func() {
			Convey("It should return true", func() {
				uc := testUserAgents["uc-browser"]

				So(NewUCBrowser(NewUAParser(uc.Android)).Match(), ShouldBeTrue)
				So(NewUCBrowser(NewUAParser(uc.IOS)).Match(), ShouldBeTrue)
				So(NewUCBrowser(NewUAParser(uc.Linux)).Match(), ShouldBeTrue)
				So(NewUCBrowser(NewUAParser(uc.Mac)).Match(), ShouldBeTrue)
				So(NewUCBrowser(NewUAParser(uc.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match UC Browser", func() {
			Convey("It should return false", func() {
				uc := NewUCBrowser(NewUAParser(testUserAgents["chrome"].Linux))
				So(uc.Match(), ShouldBeFalse)
			})
		})
	})
}
