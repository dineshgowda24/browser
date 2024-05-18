package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewIphone(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewIphone(NewUAParser(testDevices["iphone-1"])), ShouldHaveSameTypeAs, &Iphone{})
	})
}

func TestIphoneName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return iPhone", func() {
			s := NewIphone(NewUAParser(""))
			So(s.Name(), ShouldEqual, "iPhone")
		})
	})
}

func TestIphoneMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewIphone(NewUAParser(testDevices["iphone-1"])).Match(), ShouldBeTrue)
				So(NewIphone(NewUAParser(testDevices["iphone-2"])).Match(), ShouldBeTrue)
				So(NewIphone(NewUAParser(testDevices["iphone-3"])).Match(), ShouldBeTrue)
				So(NewIphone(NewUAParser(testDevices["iphone-4"])).Match(), ShouldBeTrue)
				So(NewIphone(NewUAParser(testDevices["iphone-5"])).Match(), ShouldBeTrue)
				So(NewIphone(NewUAParser(testDevices["iphone-6"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent doesn't match", func() {
			Convey("It should return false", func() {
				So(NewIphone(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
