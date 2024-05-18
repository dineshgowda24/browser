package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSamsungBrowser(t *testing.T) {
	Convey("Subject: #NewSamsungBrowser", t, func() {
		Convey("It should return a new SamsungBrowser instance", func() {
			So(NewSamsungBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &SamsungBrowser{})
		})
	})
}

func TestSamsungBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Samsung Browser", func() {
			sb := NewSamsungBrowser(NewUAParser(""))
			So(sb.Name(), ShouldEqual, "Samsung Browser")
		})
	})
}

func TestSamsungBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				sb := testUserAgents["samsung-browser"]

				So(NewSamsungBrowser(NewUAParser(sb.Android)).Version(), ShouldEqual, "3.2")
				So(NewSamsungBrowser(NewUAParser(sb.IOS)).Version(), ShouldEqual, "5.4")
				So(NewSamsungBrowser(NewUAParser(sb.Linux)).Version(), ShouldEqual, "1.0")
				So(NewSamsungBrowser(NewUAParser(sb.Windows)).Version(), ShouldEqual, "5.2")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				sb := NewSamsungBrowser(NewUAParser("Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H)"))
				So(sb.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestSamsungBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Samsung Browser", func() {
			Convey("It should return true", func() {
				sb := testUserAgents["samsung-browser"]

				So(NewSamsungBrowser(NewUAParser(sb.Android)).Match(), ShouldBeTrue)
				So(NewSamsungBrowser(NewUAParser(sb.IOS)).Match(), ShouldBeTrue)
				So(NewSamsungBrowser(NewUAParser(sb.Linux)).Match(), ShouldBeTrue)
				So(NewSamsungBrowser(NewUAParser(sb.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Samsung Browser", func() {
			Convey("It should return false", func() {
				sb := NewSamsungBrowser(NewUAParser(testUserAgents["chrome"].Linux))
				So(sb.Match(), ShouldBeFalse)
			})
		})
	})
}
