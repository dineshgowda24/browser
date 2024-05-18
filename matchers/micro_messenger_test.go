package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMicroMessenger(t *testing.T) {
	Convey("Subject: #NewMicroMessenger", t, func() {
		Convey("It should return a new MicroMessenger instance", func() {
			So(NewMicroMessenger(NewUAParser("")), ShouldHaveSameTypeAs, &MicroMessenger{})
		})
	})
}

func TestMicroMessengerName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return MicroMessenger", func() {
			s := NewMicroMessenger(NewUAParser(""))
			So(s.Name(), ShouldEqual, "MicroMessenger")
		})
	})
}

func TestMicroMessengerVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				s := testUserAgents["micro-messenger"]

				So(NewMicroMessenger(NewUAParser(s.Android)).Version(), ShouldEqual, "6.2.5.51")
				So(NewMicroMessenger(NewUAParser(s.IOS)).Version(), ShouldEqual, "6.2.3")
				So(NewMicroMessenger(NewUAParser(s.Linux)).Version(), ShouldEqual, "7.0.1")
				So(NewMicroMessenger(NewUAParser(s.Mac)).Version(), ShouldEqual, "2.3.24")
				So(NewMicroMessenger(NewUAParser(s.Windows)).Version(), ShouldEqual, "6.5.2.501")
			})
		})
	})
}

func TestMicroMessengerMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches MicroMessenger", func() {
			Convey("It should return true", func() {
				s := testUserAgents["micro-messenger"]

				So(NewMicroMessenger(NewUAParser(s.Android)).Match(), ShouldBeTrue)
				So(NewMicroMessenger(NewUAParser(s.IOS)).Match(), ShouldBeTrue)
				So(NewMicroMessenger(NewUAParser(s.Linux)).Match(), ShouldBeTrue)
				So(NewMicroMessenger(NewUAParser(s.Mac)).Match(), ShouldBeTrue)
				So(NewMicroMessenger(NewUAParser(s.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match MicroMessenger", func() {
			Convey("It should return false", func() {
				s := testUserAgents["chrome"]

				So(NewMicroMessenger(NewUAParser(s.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
