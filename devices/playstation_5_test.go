package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPlayStation5(t *testing.T) {
	Convey("Subject: #NewPlayStation5", t, func() {
		Convey("It should return a new PlayStation5 instance", func() {
			So(NewPlayStation5(NewUAParser("")), ShouldHaveSameTypeAs, &PlayStation5{})
		})
	})
}

func TestPlayStation5Name(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return PlayStation 5", func() {
			s := NewPlayStation5(NewUAParser(""))
			So(s.Name(), ShouldEqual, "PlayStation 5")
		})
	})
}

func TestPlayStation5Match(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewPlayStation5(NewUAParser(testDevices["ps5"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewPlayStation5(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
