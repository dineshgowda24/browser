package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewIOS(t *testing.T) {
	Convey("Subject: #NewIOS", t, func() {
		Convey("It should return a new IOS instance", func() {
			So(NewIOS(NewUAParser("")), ShouldHaveSameTypeAs, &IOS{})
		})
	})
}

func TestIOSName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return iOS", func() {
			So(NewIOS(NewUAParser(testPlatforms["ios"])).Name(), ShouldEqual, "iOS (iPhone)")
			So(NewIOS(NewUAParser(testPlatforms["ios-3"])).Name(), ShouldEqual, "iOS (iPhone)")
			So(NewIOS(NewUAParser(testPlatforms["ios-4"])).Name(), ShouldEqual, "iOS (iPhone)")
			So(NewIOS(NewUAParser(testPlatforms["ios-6"])).Name(), ShouldEqual, "iOS (iPad)")
			So(NewIOS(NewUAParser(testPlatforms["ios-10"])).Name(), ShouldEqual, "iOS (iPhone)")
			So(NewIOS(NewUAParser(testPlatforms["ios-11"])).Name(), ShouldEqual, "iOS (iPhone)")
		})
	})
}

func TestIOSVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			p := testPlatforms
			Convey("It should return the version", func() {
				So(NewIOS(NewUAParser(p["ios"])).Version(), ShouldEqual, "0.1")
				So(NewIOS(NewUAParser(p["ios-3"])).Version(), ShouldEqual, "3")
				So(NewIOS(NewUAParser(p["ios-4"])).Version(), ShouldEqual, "4.0.2")
				So(NewIOS(NewUAParser(p["ios-5"])).Version(), ShouldEqual, "5.1.1")
				So(NewIOS(NewUAParser(p["ios-6"])).Version(), ShouldEqual, "6.0.1")
				So(NewIOS(NewUAParser(p["ios-7"])).Version(), ShouldEqual, "7.8.1")
				So(NewIOS(NewUAParser(p["ios-8"])).Version(), ShouldEqual, "8.7.2")
				So(NewIOS(NewUAParser(p["ios-12"])).Version(), ShouldEqual, "12.1")
				So(NewIOS(NewUAParser(p["ios-13"])).Version(), ShouldEqual, "13.1.1")
				So(NewIOS(NewUAParser(p["ios-14"])).Version(), ShouldEqual, "14.7")
				So(NewIOS(NewUAParser(p["ios-15"])).Version(), ShouldEqual, "15")
				So(NewIOS(NewUAParser(p["ios-16"])).Version(), ShouldEqual, "16")
				So(NewIOS(NewUAParser(p["ios-17"])).Version(), ShouldEqual, "17.6")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewIOS(NewUAParser(testPlatforms["firefo)x"])).Version(), ShouldEqual, "0")
			})
		})
	})
}

func TestIOSMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches iOS", func() {
			Convey("It should return true", func() {
				p := testPlatforms
				So(NewIOS(NewUAParser(p["ios"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-3"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-4"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-5"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-6"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-7"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-8"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-12"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-13"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-14"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-15"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-16"])).Match(), ShouldBeTrue)
				So(NewIOS(NewUAParser(p["ios-17"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match iOS", func() {
			Convey("It should return false", func() {
				So(NewIOS(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
