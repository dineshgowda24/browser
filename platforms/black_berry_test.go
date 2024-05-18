package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBlackBerry(t *testing.T) {
	Convey("Subject: #NewBlackBerry", t, func() {
		Convey("It should return a new BlackBerry instance", func() {
			So(NewBlackBerry(NewUAParser("")), ShouldHaveSameTypeAs, &BlackBerry{})
		})
	})
}

func TestBlackBerryName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			So(NewBlackBerry(NewUAParser("")).Name(), ShouldEqual, "BlackBerry")
		})
	})
}

func TestBlackBerryVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				p := testPlatforms
				So(NewBlackBerry(NewUAParser(p["blackberry"])).Version(), ShouldEqual, "4.1.0")
				So(NewBlackBerry(NewUAParser(p["blackberry-4"])).Version(), ShouldEqual, "4.2.1")
				So(NewBlackBerry(NewUAParser(p["blackberry-5"])).Version(), ShouldEqual, "5.0.0.187")
				So(NewBlackBerry(NewUAParser(p["blackberry-6"])).Version(), ShouldEqual, "6.0.0.480")
				So(NewBlackBerry(NewUAParser(p["blackberry-7"])).Version(), ShouldEqual, "7.1.0.336")
				So(NewBlackBerry(NewUAParser(p["blackberry-10"])).Version(), ShouldEqual, "10.3.2.2339")
			})
		})
	})
}

func TestBlackBerryMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches", func() {
			Convey("It should return true", func() {
				p := testPlatforms
				So(NewBlackBerry(NewUAParser(p["blackberry"])).Match(), ShouldBeTrue)
				So(NewBlackBerry(NewUAParser(p["blackberry-4"])).Match(), ShouldBeTrue)
				So(NewBlackBerry(NewUAParser(p["blackberry-5"])).Match(), ShouldBeTrue)
				So(NewBlackBerry(NewUAParser(p["blackberry-6"])).Match(), ShouldBeTrue)
				So(NewBlackBerry(NewUAParser(p["blackberry-7"])).Match(), ShouldBeTrue)
				So(NewBlackBerry(NewUAParser(p["blackberry-10"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewBlackBerry(NewUAParser("")).Match(), ShouldBeFalse)
			})
		})
	})
}
