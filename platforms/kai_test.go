package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewKaiOS(t *testing.T) {
	Convey("Subject: #NewKaiOS", t, func() {
		Convey("It should return a new KaiOS instance", func() {
			So(NewKaiOS(NewUAParser("")), ShouldHaveSameTypeAs, &KaiOS{})
		})
	})
}

func TestKaiOSName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return KaiOS", func() {
			So(NewKaiOS(NewUAParser("")).Name(), ShouldEqual, "Kai OS")
		})
	})
}

func TestKaiOSVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewKaiOS(NewUAParser(testPlatforms["kai-os-1"])).Version(), ShouldEqual, "1.0")
				So(NewKaiOS(NewUAParser(testPlatforms["kai-os-2"])).Version(), ShouldEqual, "2.0")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewKaiOS(NewUAParser(testPlatforms["firefox"])).Version(), ShouldEqual, "")
			})
		})
	})
}

func TestKaiOSMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches KaiOS", func() {
			Convey("It should return true", func() {
				So(NewKaiOS(NewUAParser(testPlatforms["kai-os-1"])).Match(), ShouldBeTrue)
				So(NewKaiOS(NewUAParser(testPlatforms["kai-os-2"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match KaiOS", func() {
			Convey("It should return false", func() {
				So(NewKaiOS(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
