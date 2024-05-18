package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPlayStation4(t *testing.T) {
	Convey("Subject: #NewPlayStation4", t, func() {
		Convey("It should return a new PlayStation4 instance", func() {
			So(NewPlayStation4(NewUAParser("")), ShouldHaveSameTypeAs, &PlayStation4{})
		})
	})
}

func TestPlayStation4Name(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return PlayStation 4", func() {
			s := NewPlayStation4(NewUAParser(""))
			So(s.Name(), ShouldEqual, "PlayStation 4")
		})
	})
}

func TestPlayStation4Match(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewPlayStation4(NewUAParser(testDevices["ps4"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewPlayStation4(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
