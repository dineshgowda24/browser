package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewKindle(t *testing.T) {
	Convey("Subject: #NewKindle", t, func() {
		Convey("It should return a new Kindle instance", func() {
			So(NewKindle(NewUAParser("")), ShouldHaveSameTypeAs, &Kindle{})
		})
	})
}

func TestKindleName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Kindle", func() {
			s := NewKindle(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Kindle")
		})
	})
}

func TestKindleMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewKindle(NewUAParser(testDevices["kindle-1"])).Match(), ShouldBeTrue)
				So(NewKindle(NewUAParser(testDevices["kindle-2"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewKindle(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
