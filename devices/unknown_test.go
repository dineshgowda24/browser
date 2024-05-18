package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUnknown(t *testing.T) {
	Convey("Subject: #NewUnknown", t, func() {
		Convey("It should return a new Unknown instance", func() {
			So(NewUnknown(NewUAParser("")), ShouldHaveSameTypeAs, &Unknown{})
		})
	})
}

func TestUnknownName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Unknown", func() {
			s := NewUnknown(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Unknown")
		})
	})
}

func TestUnknownMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("It should always return true", func() {
			So(NewUnknown(NewUAParser("")).Match(), ShouldBeTrue)
		})
	})
}
