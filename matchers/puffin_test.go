package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPuffin(t *testing.T) {
	Convey("Subject: #NewPuffin", t, func() {
		Convey("It should return a new Puffin instance", func() {
			So(NewPuffin(NewUAParser("")), ShouldHaveSameTypeAs, &Puffin{})
		})
	})
}

func TestPuffinName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Puffin", func() {
			p := NewPuffin(NewUAParser(""))
			So(p.Name(), ShouldEqual, "Puffin")
		})
	})
}

func TestPuffinVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["puffin"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewPuffin(NewUAParser(ua.Android)).Version(), ShouldEqual, "4.3.0.1852")
				So(NewPuffin(NewUAParser(ua.IOS)).Version(), ShouldEqual, "4.5.0")
				So(NewPuffin(NewUAParser(ua.Linux)).Version(), ShouldEqual, "610")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewPuffin(NewUAParser(ua.Mac)).Version(), ShouldEqual, "0.0")
				So(NewPuffin(NewUAParser(ua.Windows)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestPuffinMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Puffin", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["puffin"]

				So(NewPuffin(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewPuffin(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
				So(NewPuffin(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewPuffin(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Puffin", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]

				So(NewPuffin(NewUAParser(ua.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
