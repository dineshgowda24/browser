package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSurface(t *testing.T) {
	Convey("Subject: #NewSurface", t, func() {
		Convey("It should return a new Surface instance", func() {
			So(NewSurface(NewUAParser("")), ShouldHaveSameTypeAs, &Surface{})
		})
	})
}

func TestSurfaceName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Surface", func() {
			s := NewSurface(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Surface")
		})
	})
}

func TestSurfaceMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewSurface(NewUAParser(testDevices["surface-1"])).Match(), ShouldBeTrue)
				So(NewSurface(NewUAParser(testDevices["surface-2"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewSurface(NewUAParser(testDevices["ps5"])).Match(), ShouldBeFalse)
			})
		})
	})
}
