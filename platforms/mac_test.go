package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMac(t *testing.T) {
	Convey("Subject: #NewMac", t, func() {
		Convey("It should return a new Mac instance", func() {
			So(NewMac(NewUAParser("")), ShouldHaveSameTypeAs, &Mac{})
		})
	})
}

func TestMacName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return macOS", func() {
			So(NewMac(NewUAParser(testPlatforms["mac-os"])).Name(), ShouldEqual, "Mac OS X")
			So(NewMac(NewUAParser(testPlatforms["mac-os-x"])).Name(), ShouldEqual, "Mac OS X")
			So(NewMac(NewUAParser(testPlatforms["mac-os-1"])).Name(), ShouldEqual, "macOS")
		})
	})
}

func TestMacVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewMac(NewUAParser(testPlatforms["mac-os-x"])).Version(), ShouldEqual, "10.11")
				So(NewMac(NewUAParser(testPlatforms["mac-os-1"])).Version(), ShouldEqual, "10.14.6")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewMac(NewUAParser(testPlatforms["mac-os"])).Version(), ShouldEqual, "0")
			})
		})
	})
}

func TestMacMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Mac", func() {
			Convey("It should return true", func() {
				So(NewMac(NewUAParser(testPlatforms["mac-os"])).Match(), ShouldBeTrue)
				So(NewMac(NewUAParser(testPlatforms["mac-os-x"])).Match(), ShouldBeTrue)
				So(NewMac(NewUAParser(testPlatforms["mac-os-1"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Mac", func() {
			Convey("It should return false", func() {
				So(NewMac(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
