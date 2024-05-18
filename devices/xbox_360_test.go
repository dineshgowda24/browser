package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewXbox360(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewXbox360(NewUAParser(testDevices["xbox360-1"])), ShouldHaveSameTypeAs, &Xbox360{})
	})
}

func TestXbox360Name(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Xbox 360", func() {
			s := NewXbox360(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Xbox 360")
		})
	})
}

func TestXbox360Match(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewXbox360(NewUAParser(testDevices["xbox360-1"])).Match(), ShouldBeTrue)
				So(NewXbox360(NewUAParser(testDevices["xbox360-2"])).Match(), ShouldBeTrue)
				So(NewXbox360(NewUAParser(testDevices["xbox360-3"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewXbox360(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
