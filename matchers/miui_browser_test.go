package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMiuiBrowser(t *testing.T) {
	Convey("Subject: #NewMiuiBrowser", t, func() {
		Convey("It should return a new MiuiBrowser instance", func() {
			So(NewMiuiBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &MiuiBrowser{})
		})
	})
}

func TestMiuiBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Miui Browser", func() {
			So(NewMiuiBrowser(NewUAParser("")).Name(), ShouldEqual, "Miui Browser")
		})
	})
}

func TestMiuiBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["miui-browser"]

				So(NewMiuiBrowser(NewUAParser(ua.Android)).Version(), ShouldEqual, "2.1.1")
				So(NewMiuiBrowser(NewUAParser(ua.IOS)).Version(), ShouldEqual, "1.0")
				So(NewMiuiBrowser(NewUAParser(ua.Linux)).Version(), ShouldEqual, "2.0.1")
				So(NewMiuiBrowser(NewUAParser(ua.Windows)).Version(), ShouldEqual, "12.5.2")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				mb := NewMiuiBrowser(NewUAParser(testUserAgents["chrome"].Linux))
				So(mb.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestMiuiBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches MiuiBrowser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["miui-browser"]

				So(NewMiuiBrowser(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewMiuiBrowser(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
				So(NewMiuiBrowser(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewMiuiBrowser(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match MiuiBrowser", func() {
			Convey("It should return false", func() {
				mb := NewMiuiBrowser(NewUAParser(testUserAgents["chrome"].Linux))
				So(mb.Match(), ShouldBeFalse)
			})
		})
	})
}
