package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSamsung(t *testing.T) {
	Convey("Subject: #NewSamsung", t, func() {
		Convey("It returns a Samsung platform", func() {
			s := NewSamsung(NewUAParser(testDevices["samsung-1"]))
			So(s, ShouldHaveSameTypeAs, &Samsung{})
		})
	})
}

func TestSamsungName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It returns the correct name", func() {
			So(NewSamsung(NewUAParser(testDevices["samsung-1"])).Name(), ShouldEqual, "Samsung SM-N900")
			So(NewSamsung(NewUAParser(testDevices["samsung-2"])).Name(), ShouldEqual, "Samsung SM-J700F")
			So(NewSamsung(NewUAParser(testDevices["samsung-3"])).Name(), ShouldEqual, "Samsung SM-G928K")
		})
	})
}

func TestSamsungMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It returns true", func() {
				So(NewSamsung(NewUAParser(testDevices["samsung-1"])).Match(), ShouldBeTrue)
				So(NewSamsung(NewUAParser(testDevices["samsung-2"])).Match(), ShouldBeTrue)
				So(NewSamsung(NewUAParser(testDevices["samsung-3"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It returns false", func() {
				s := NewSamsung(NewUAParser(testDevices["iphone-7"]))
				So(s.Match(), ShouldBeFalse)
			})
		})
	})
}
