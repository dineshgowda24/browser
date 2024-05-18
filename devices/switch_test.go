package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSwitch(t *testing.T) {
	Convey("Subject: #NewSwitch", t, func() {
		Convey("It should return a new Switch instance", func() {
			So(NewSwitch(NewUAParser("")), ShouldHaveSameTypeAs, &Switch{})
		})
	})
}

func TestSwitchName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Nintendo Switch", func() {
			s := NewSwitch(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Nintendo Switch")
		})
	})
}

func TestSwitchMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewSwitch(NewUAParser(testDevices["switch-1"])).Match(), ShouldBeTrue)
				So(NewSwitch(NewUAParser(testDevices["switch-2"])).Match(), ShouldBeTrue)
				So(NewSwitch(NewUAParser(testDevices["switch-3"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewSwitch(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
