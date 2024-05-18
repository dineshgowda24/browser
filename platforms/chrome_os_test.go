package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewChromeOS(t *testing.T) {
	Convey("Subject: #NewChromeOS", t, func() {
		Convey("It should return a new ChromeOS instance", func() {
			So(NewChromeOS(NewUAParser("")), ShouldHaveSameTypeAs, &ChromeOS{})
		})
	})
}

func TestChromeOSName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Chrome OS", func() {
			So(NewChromeOS(NewUAParser("")).Name(), ShouldEqual, "Chrome OS")
		})
	})
}

func TestChromeOSVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				p := testPlatforms
				So(NewChromeOS(NewUAParser(p["chrome-os-armv7l"])).Version(), ShouldEqual, "7077.134.0")
				So(NewChromeOS(NewUAParser(p["chrome-os-armv7i"])).Version(), ShouldEqual, "6158.49.0")
				So(NewChromeOS(NewUAParser(p["chrome-os-armv8l"])).Version(), ShouldEqual, "8872.54.0")
				So(NewChromeOS(NewUAParser(p["chrome-os-x86-64"])).Version(), ShouldEqual, "7647.84.0")
				So(NewChromeOS(NewUAParser(p["chrome-os-i686"])).Version(), ShouldEqual, "6310.68.0")
				So(NewChromeOS(NewUAParser(p["chrome-os-aarch64"])).Version(), ShouldEqual, "9202.64.0")
			})
		})
	})
}

func TestChromeOSMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Chrome OS", func() {
			Convey("It should return true", func() {
				p := testPlatforms
				So(NewChromeOS(NewUAParser(p["chrome-os-armv7l"])).Match(), ShouldBeTrue)
				So(NewChromeOS(NewUAParser(p["chrome-os-armv7i"])).Match(), ShouldBeTrue)
				So(NewChromeOS(NewUAParser(p["chrome-os-armv8l"])).Match(), ShouldBeTrue)
				So(NewChromeOS(NewUAParser(p["chrome-os-x86-64"])).Match(), ShouldBeTrue)
				So(NewChromeOS(NewUAParser(p["chrome-os-i686"])).Match(), ShouldBeTrue)
				So(NewChromeOS(NewUAParser(p["chrome-os-aarch64"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewChromeOS(NewUAParser(testPlatforms["android-10"])).Match(), ShouldBeFalse)
			})
		})
	})
}
