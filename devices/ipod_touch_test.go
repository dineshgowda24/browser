package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewIpodTouch(t *testing.T) {
	Convey("Subject: #NewIpodTouch", t, func() {
		Convey("It should return a new IpodTouch instance", func() {
			So(NewIpodTouch(NewUAParser("")), ShouldHaveSameTypeAs, &IpodTouch{})
		})
	})
}

func TestIpodTouchName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return iPod Touch", func() {
			s := NewIpodTouch(NewUAParser(""))
			So(s.Name(), ShouldEqual, "iPod Touch")
		})
	})
}

func TestIpodTouchMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewIpodTouch(NewUAParser(testDevices["ipod-touch-1"])).Match(), ShouldBeTrue)
				So(NewIpodTouch(NewUAParser(testDevices["ipod-touch-2"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewIpodTouch(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
