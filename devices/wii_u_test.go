package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWiiU(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewWiiU(testDevices["wiiu-1"]), ShouldHaveSameTypeAs, &WiiU{})
	})
}

func TestWiiUName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return WiiU", func() {
			s := NewWiiU("")
			So(s.Name(), ShouldEqual, "Nintendo WiiU")
		})
	})
}

func TestWiiUMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewWiiU(testDevices["wiiu-1"]).Match(), ShouldBeTrue)
				So(NewWiiU(testDevices["wiiu-2"]).Match(), ShouldBeTrue)
				So(NewWiiU(testDevices["wiiu-3"]).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewWiiU(testDevices["bb-playbook-1"]).Match(), ShouldBeFalse)
			})
		})
	})
}
