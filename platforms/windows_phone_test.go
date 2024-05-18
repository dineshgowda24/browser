package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWindowsPhone(t *testing.T) {
	Convey("Subject: #NewWindowsPhone", t, func() {
		Convey("It should return a new WindowsPhone instance", func() {
			So(NewWindowsPhone(NewUAParser("")), ShouldHaveSameTypeAs, &WindowsPhone{})
		})
	})
}

func TestWindowsPhoneName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Windows Phone", func() {
			So(NewWindowsPhone(NewUAParser("")).Name(), ShouldEqual, "Windows Phone")
		})
	})
}

func TestWindowsPhoneVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				wp := NewWindowsPhone(NewUAParser(testPlatforms["windows-phone"]))
				So(wp.Version(), ShouldEqual, "10.0")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				wp := NewWindowsPhone(NewUAParser(testPlatforms["windows-ce"]))
				So(wp.Version(), ShouldEqual, "")
			})
		})
	})
}

func TestWindowsPhoneMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Windows Phone", func() {
			Convey("It should return true", func() {
				So(NewWindowsPhone(NewUAParser(testPlatforms["windows-phone"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Windows Phone", func() {
			Convey("It should return false", func() {
				So(NewWindowsPhone(NewUAParser(testPlatforms["windows-ce"])).Match(), ShouldBeFalse)
			})
		})
	})
}
