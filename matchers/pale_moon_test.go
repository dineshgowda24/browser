package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPaleMoon(t *testing.T) {
	Convey("Subject: #NewPaleMoon", t, func() {
		Convey("It should return a new PaleMoon instance", func() {
			So(NewPaleMoon(NewUAParser("")), ShouldHaveSameTypeAs, &PaleMoon{})
		})
	})
}

func TestPaleMoonName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return PaleMoon", func() {
			So(NewPaleMoon(NewUAParser("")).Name(), ShouldEqual, "Pale Moon")
		})
	})
}

func TestPaleMoonVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["pale-moon"]

				So(NewPaleMoon(NewUAParser(ua.Linux)).Version(), ShouldEqual, "25.6.0")
				So(NewPaleMoon(NewUAParser(ua.Mac)).Version(), ShouldEqual, "25.3.1")
				So(NewPaleMoon(NewUAParser(ua.Windows)).Version(), ShouldEqual, "25.6.0")
				So(NewPaleMoon(NewUAParser(ua.Android)).Version(), ShouldEqual, "25.4.1")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				pm := NewPaleMoon(NewUAParser(testUserAgents["chrome"].Linux))
				So(pm.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestPaleMoonMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches PaleMoon", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["pale-moon"]

				So(NewPaleMoon(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewPaleMoon(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
				So(NewPaleMoon(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewPaleMoon(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match PaleMoon", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]

				So(NewPaleMoon(NewUAParser(ua.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
