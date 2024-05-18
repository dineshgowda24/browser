package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMaxthon(t *testing.T) {
	Convey("Subject: #NewMaxthon", t, func() {
		Convey("It should return a new Maxthon instance", func() {
			So(NewMaxthon(NewUAParser("")), ShouldHaveSameTypeAs, &Maxthon{})
		})
	})
}

func TestMaxthonName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Maxthon", func() {
			m := NewMaxthon(NewUAParser(""))
			So(m.Name(), ShouldEqual, "Maxthon")
		})
	})
}

func TestMaxthonVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["maxthon"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewMaxthon(NewUAParser(ua.Linux)).Version(), ShouldEqual, "1.0.5.3")
				So(NewMaxthon(NewUAParser(ua.Mac)).Version(), ShouldEqual, "4.5.2")
				So(NewMaxthon(NewUAParser(ua.Windows)).Version(), ShouldEqual, "4.4.3.4000")
				So(NewMaxthon(NewUAParser(ua.IOS)).Version(), ShouldEqual, "4.1.8.2000")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewMaxthon(NewUAParser(ua.Android)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestMaxthonMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Maxthon", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["maxthon"]

				So(NewMaxthon(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewMaxthon(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
				So(NewMaxthon(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewMaxthon(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewMaxthon(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Maxthon", func() {
			Convey("It should return false", func() {
				So(NewMaxthon(NewUAParser(testUserAgents["chrome"].Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
