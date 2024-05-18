package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWindowsMobile(t *testing.T) {
	Convey("Subject: #NewWindowsMobile", t, func() {
		Convey("It should return a new WindowsMobile instance", func() {
			So(NewWindowsMobile(NewUAParser("")), ShouldHaveSameTypeAs, &WindowsMobile{})
		})
	})
}

func TestWindowsMobileName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Windows Mobile", func() {
			wm := NewWindowsMobile(NewUAParser(""))
			So(wm.Name(), ShouldEqual, "Windows Mobile")
		})
	})
}

func TestWindowsMobileVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("It should return 0", func() {
			So(NewWindowsMobile(NewUAParser(testPlatforms["windows-ce"])).Version(), ShouldEqual, "0")
		})
	})
}

func TestWindowsMobileMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Windows Mobile", func() {
			Convey("It should return true", func() {
				So(NewWindowsMobile(NewUAParser(testPlatforms["windows-ce"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Windows Mobile", func() {
			Convey("It should return false", func() {
				So(NewWindowsMobile(NewUAParser(testPlatforms["windows"])).Match(), ShouldBeFalse)
			})
		})
	})
}
