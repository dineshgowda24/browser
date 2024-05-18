package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewTV(t *testing.T) {
	Convey("Subject: #NewTv", t, func() {
		Convey("It should return a new Tv instance", func() {
			So(NewTV(NewUAParser("")), ShouldHaveSameTypeAs, &TV{})
		})
	})
}

func TestTVName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return TV", func() {
			s := NewTV(NewUAParser(""))
			So(s.Name(), ShouldEqual, "TV")
		})
	})
}

func TestTVMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewTV(NewUAParser(testDevices["tv-1"])).Match(), ShouldBeTrue)
				So(NewTV(NewUAParser(testDevices["tv-2"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewSwitch(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
