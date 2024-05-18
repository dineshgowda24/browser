package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWindows(t *testing.T) {
	Convey("Subject: #NewWindows", t, func() {
		Convey("It should return a new Windows instance", func() {
			So(NewWindows(NewUAParser(testPlatforms["windows-xp"])), ShouldHaveSameTypeAs, &Windows{})
		})
	})
}

func TestWindowsName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Windows", func() {
			w := NewWindows(NewUAParser(testPlatforms["windows-xp"]))
			So(w.Name(), ShouldEqual, windowsName)
		})
	})
}

func TestWindowsVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewWindows(NewUAParser(testPlatforms["windows-xp"])).Version(), ShouldEqual, "5.1")
				So(NewWindows(NewUAParser(testPlatforms["windows-vista"])).Version(), ShouldEqual, "6.0")
				So(NewWindows(NewUAParser(testPlatforms["windows-7"])).Version(), ShouldEqual, "6.1")
				So(NewWindows(NewUAParser(testPlatforms["windows-8"])).Version(), ShouldEqual, "6.2")
				So(NewWindows(NewUAParser(testPlatforms["windows-8-1"])).Version(), ShouldEqual, "6.3")
				So(NewWindows(NewUAParser(testPlatforms["windows-10"])).Version(), ShouldEqual, "10.0")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewWindows(NewUAParser(testPlatforms["linux"])).Version(), ShouldEqual, "0")
			})
		})
	})
}

func TestWindowsMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Windows", func() {
			Convey("It should return true", func() {
				So(NewWindows(NewUAParser(testPlatforms["windows-xp"])).Match(), ShouldBeTrue)
				So(NewWindows(NewUAParser(testPlatforms["windows-vista"])).Match(), ShouldBeTrue)
				So(NewWindows(NewUAParser(testPlatforms["windows-7"])).Match(), ShouldBeTrue)
				So(NewWindows(NewUAParser(testPlatforms["windows-8"])).Match(), ShouldBeTrue)
				So(NewWindows(NewUAParser(testPlatforms["windows-8-1"])).Match(), ShouldBeTrue)
				So(NewWindows(NewUAParser(testPlatforms["windows-10"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Windows", func() {
			Convey("It should return false", func() {
				So(NewWindows(NewUAParser(testPlatforms["linux"])).Match(), ShouldBeFalse)
			})
		})
	})
}
