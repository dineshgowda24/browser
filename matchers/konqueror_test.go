package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewKonqueror(t *testing.T) {
	Convey("Subject: #NewKonqueror", t, func() {
		Convey("It should return a new Konqueror instance", func() {
			So(NewKonqueror(NewUAParser("")), ShouldHaveSameTypeAs, &Konqueror{})
		})
	})
}

func TestKonquerorName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Konqueror", func() {
			k := NewKonqueror(NewUAParser(""))
			So(k.Name(), ShouldEqual, konquerorName)
		})
	})
}

func TestKonquerorVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["konqueror"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewKonqueror(NewUAParser(ua.Linux)).Version(), ShouldEqual, "4.14.1")
				So(NewKonqueror(NewUAParser(ua.Mac)).Version(), ShouldEqual, "5.0")
				So(NewKonqueror(NewUAParser(ua.Windows)).Version(), ShouldEqual, "4.10")
				So(NewKonqueror(NewUAParser(ua.Android)).Version(), ShouldEqual, "10.1")
				So(NewKonqueror(NewUAParser(ua.IOS)).Version(), ShouldEqual, "14.1")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				k := NewKonqueror(NewUAParser(testUserAgents["chrome"].Linux))
				So(k.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestKonquerorMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Konqueror", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["konqueror"]
				So(NewKonqueror(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewKonqueror(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
				So(NewKonqueror(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewKonqueror(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewKonqueror(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Konqueror", func() {
			Convey("It should return false", func() {
				So(NewKonqueror(NewUAParser(testUserAgents["chrome"].Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
