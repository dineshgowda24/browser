package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSogouBrowser(t *testing.T) {
	Convey("Subject: #NewSogouBrowser", t, func() {
		Convey("It should return a new SogouBrowser instance", func() {
			So(NewSogouBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &SogouBrowser{})
		})
	})
}

func TestSogouBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Sogou Browser", func() {
			s := NewSogouBrowser(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Sogou Browser")
		})
	})
}

func TestSogouBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				s := testUserAgents["sogou-browser"]

				So(NewSogouBrowser(NewUAParser(s.Android)).Version(), ShouldEqual, "3.8.2")
				So(NewSogouBrowser(NewUAParser(s.IOS)).Version(), ShouldEqual, "5.3.0")
			})
		})
	})
}

func TestSogouBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Sogou Browser", func() {
			Convey("It should return true", func() {
				s := testUserAgents["sogou-browser"]

				So(NewSogouBrowser(NewUAParser(s.Android)).Match(), ShouldBeTrue)
				So(NewSogouBrowser(NewUAParser(s.IOS)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Sogou Browser", func() {
			Convey("It should return false", func() {
				s := testUserAgents["chrome"]

				So(NewSogouBrowser(NewUAParser(s.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
