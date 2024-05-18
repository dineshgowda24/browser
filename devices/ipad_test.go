package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewIpad(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewIpad(NewUAParser(testDevices["ipad-1"])), ShouldHaveSameTypeAs, &Ipad{})
	})
}

func TestIpadName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return iPad", func() {
			s := NewIpad(NewUAParser(""))
			So(s.Name(), ShouldEqual, "iPad")
		})
	})
}

func TestIpadMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewIpad(NewUAParser(testDevices["ipad-1"])).Match(), ShouldBeTrue)
				So(NewIpad(NewUAParser(testDevices["ipad-2"])).Match(), ShouldBeTrue)
				So(NewIpad(NewUAParser(testDevices["ipad-3"])).Match(), ShouldBeTrue)
				So(NewIpad(NewUAParser(testDevices["ipad-4"])).Match(), ShouldBeTrue)
				So(NewIpad(NewUAParser(testDevices["ipad-5"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent doesn't match", func() {
			Convey("It should return false", func() {
				So(NewIpad(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
