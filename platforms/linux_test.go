package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewLinux(t *testing.T) {
	Convey("Subject: #NewLinux", t, func() {
		Convey("It should return a new Linux instance", func() {
			So(NewLinux(NewUAParser("")), ShouldHaveSameTypeAs, &Linux{})
		})
	})
}

func TestLinuxName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Linux", func() {
			So(NewLinux(NewUAParser("")).Name(), ShouldEqual, "Generic Linux")
		})
	})
}

func TestLinuxVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewLinux(NewUAParser(testPlatforms["linux"])).Version(), ShouldEqual, "0")
			})
		})
	})
}

func TestLinuxMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Linux", func() {
			Convey("It should return true", func() {
				So(NewLinux(NewUAParser(testPlatforms["linux"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Linux", func() {
			Convey("It should return false", func() {
				So(NewLinux(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
