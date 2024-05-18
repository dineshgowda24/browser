package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAdobeAir(t *testing.T) {
	Convey("Subject: #NewAdobeAir", t, func() {
		Convey("It should return a new AdobeAir instance", func() {
			So(NewAdobeAir(NewUAParser("")), ShouldHaveSameTypeAs, &AdobeAir{})
		})
	})
}

func TestAdobeAirName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return AdobeAir", func() {
			aa := NewAdobeAir(NewUAParser(""))
			So(aa.Name(), ShouldEqual, "Adobe AIR")
		})
	})
}

func TestAdobeAirVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				aa := testPlatforms["adobe-air"]

				So(NewAdobeAir(NewUAParser(aa)).Version(), ShouldEqual, "18.0")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				a := testPlatforms["android-10"]
				So(NewAdobeAir(NewUAParser(a)).Version(), ShouldEqual, "")
			})
		})
	})
}

func TestAdobeAirMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches AdobeAir", func() {
			Convey("It should return true", func() {
				aa := testPlatforms["adobe-air"]
				So(NewAdobeAir(NewUAParser(aa)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				a := testPlatforms["android-10"]
				So(NewAdobeAir(NewUAParser(a)).Match(), ShouldBeFalse)
			})
		})
	})
}
