package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewKindleFire(t *testing.T) {
	Convey("Subject: #NewKindleFire", t, func() {
		Convey("It should return a new KindleFire instance", func() {
			So(NewKindleFire(NewUAParser("")), ShouldHaveSameTypeAs, &KindleFire{})
		})
	})
}

func TestKindleFireName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Kindle Fire", func() {
			s := NewKindleFire(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Kindle Fire")
		})
	})
}

func TestKindleFireMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewKindleFire(NewUAParser(testDevices["kindle-fire-1"])).Match(), ShouldBeTrue)
				So(NewKindleFire(NewUAParser(testDevices["kindle-fire-2"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewKindleFire(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
