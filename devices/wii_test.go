package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWii(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewWii(NewUAParser(testDevices["wii-1"])), ShouldHaveSameTypeAs, &Wii{})
	})
}

func TestWiiName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Wii", func() {
			s := NewWii(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Nintendo Wii")
		})
	})
}

func TestWiiMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewWii(NewUAParser(testDevices["wii-1"])).Match(), ShouldBeTrue)
				So(NewWii(NewUAParser(testDevices["wii-2"])).Match(), ShouldBeTrue)
				So(NewWii(NewUAParser(testDevices["wii-3"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewWii(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
