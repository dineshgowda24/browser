package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPSP(t *testing.T) {
	Convey("Subject: #NewPSP", t, func() {
		Convey("It should return a new PSP instance", func() {
			So(NewPSP(""), ShouldHaveSameTypeAs, &PSP{})
		})
	})
}

func TestPSPName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return PlayStation Portable", func() {
			s := NewPSP("")
			So(s.Name(), ShouldEqual, "PlayStation Portable")
		})
	})
}

func TestPSPMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewPSP(testDevices["psp-1"]).Match(), ShouldBeTrue)
				So(NewPSP(testDevices["psp-2"]).Match(), ShouldBeTrue)
				So(NewPSP(testDevices["psp-3"]).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewPSP(testDevices["ps5"]).Match(), ShouldBeFalse)
			})
		})
	})
}
