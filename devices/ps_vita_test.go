package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPSViat(t *testing.T) {
	Convey("Subject: #NewPSVita", t, func() {
		Convey("It should return a new PSVita instance", func() {
			So(NewPSVita(NewUAParser("")), ShouldHaveSameTypeAs, &PSVita{})
		})
	})
}

func TestPSVitaName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return PlayStation Vita", func() {
			s := NewPSVita(NewUAParser(""))
			So(s.Name(), ShouldEqual, "PlayStation Vita")
		})
	})
}

func TestPSVitaMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewPSVita(NewUAParser(testDevices["ps-vita-1"])).Match(), ShouldBeTrue)
				So(NewPSVita(NewUAParser(testDevices["ps-vita-2"])).Match(), ShouldBeTrue)
				So(NewPSVita(NewUAParser(testDevices["ps-vita-3"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewPSVita(NewUAParser(testDevices["ps5"])).Match(), ShouldBeFalse)
			})
		})
	})
}
