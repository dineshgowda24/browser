package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFirefoxOS(t *testing.T) {
	Convey("Subject: #NewFirefoxOS", t, func() {
		Convey("It should return a new FirefoxOS instance", func() {
			So(NewFirefoxOS(""), ShouldHaveSameTypeAs, &FirefoxOS{})
		})
	})
}

func TestFirefoxOSName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Firefox OS", func() {
			So(NewFirefoxOS("").Name(), ShouldEqual, "Firefox OS")
		})
	})
}
func TestFirefoxOSVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			p := testPlatforms
			So(NewFirefoxOS(p["firefox"]).Version(), ShouldEqual, "0")
		})
	})
}

func TestFirefoxOSMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Firefox OS", func() {
			Convey("It should return true", func() {
				p := testPlatforms
				So(NewFirefoxOS(p["firefox"]).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewFirefoxOS(testPlatforms["android-10"]).Match(), ShouldBeFalse)
			})
		})
	})
}
