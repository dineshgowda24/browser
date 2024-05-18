package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPlayStation3(t *testing.T) {
	Convey("Subject: #NewPlayStation3", t, func() {
		Convey("It should return a new PlayStation3 instance", func() {
			So(NewPlayStation3(NewUAParser("")), ShouldHaveSameTypeAs, &PlayStation3{})
		})
	})
}

func TestPlayStation3Name(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return PlayStation 3", func() {
			s := NewPlayStation3(NewUAParser(""))
			So(s.Name(), ShouldEqual, "PlayStation 3")
		})
	})
}

func TestPlayStation3Match(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewPlayStation3(NewUAParser(testDevices["ps3"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewPlayStation3(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
