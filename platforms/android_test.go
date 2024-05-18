package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAndroid(t *testing.T) {
	Convey("Subject: #NewAndroid", t, func() {
		Convey("It should return a new Android instance", func() {
			So(NewAndroid(NewUAParser("")), ShouldHaveSameTypeAs, &Android{})
		})
	})
}

func TestAndroidName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Android", func() {
			a := NewAndroid(NewUAParser(""))
			So(a.Name(), ShouldEqual, androidName)
		})
	})
}

func TestAndroidVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			p := testPlatforms
			Convey("It should return the version", func() {
				So(NewAndroid(NewUAParser(p["android-eclair-2-0"])).Version(), ShouldEqual, "2.0")
				So(NewAndroid(NewUAParser(p["android-eclair-2-1"])).Version(), ShouldEqual, "2.1.6")
				So(NewAndroid(NewUAParser(p["android-froyo-2-2"])).Version(), ShouldEqual, "2.2")
				So(NewAndroid(NewUAParser(p["android-froyo-2-2-0"])).Version(), ShouldEqual, "2.2.0")
				So(NewAndroid(NewUAParser(p["android-gingerbread-2-3-5"])).Version(), ShouldEqual, "2.3.5")
				So(NewAndroid(NewUAParser(p["android-gingerbread-2-3-6"])).Version(), ShouldEqual, "2.3.6")
				So(NewAndroid(NewUAParser(p["android-gingerbread-2-3-8"])).Version(), ShouldEqual, "2.3.8")
				So(NewAndroid(NewUAParser(p["android-honeycomb-3-1-4"])).Version(), ShouldEqual, "3.1.4")
				So(NewAndroid(NewUAParser(p["android-honeycomb-3-2"])).Version(), ShouldEqual, "3.2")
				So(NewAndroid(NewUAParser(p["android-ice-cream-sandwich-4-0-3"])).Version(), ShouldEqual, "4.0.3")
				So(NewAndroid(NewUAParser(p["android-ice-cream-sandwich-4-0-4"])).Version(), ShouldEqual, "4.0.4")
				So(NewAndroid(NewUAParser(p["android-jelly-bean-4.1"])).Version(), ShouldEqual, "4.1.1")
				So(NewAndroid(NewUAParser(p["android-jelly-bean-4.2"])).Version(), ShouldEqual, "4.2.2")
				So(NewAndroid(NewUAParser(p["android-jelly-bean-4.3"])).Version(), ShouldEqual, "4.3")
				So(NewAndroid(NewUAParser(p["android-kitkat-4.4"])).Version(), ShouldEqual, "4.4.2")
				So(NewAndroid(NewUAParser(p["android-lollipop-5.0"])).Version(), ShouldEqual, "5.0.2")
				So(NewAndroid(NewUAParser(p["android-loolipop-5.1"])).Version(), ShouldEqual, "5.1")
				So(NewAndroid(NewUAParser(p["android-marshmallow-6-0"])).Version(), ShouldEqual, "6.0.1")
				So(NewAndroid(NewUAParser(p["android-nougat-7-0"])).Version(), ShouldEqual, "7.0")
				So(NewAndroid(NewUAParser(p["android-nougat-7-1"])).Version(), ShouldEqual, "7.1")
				So(NewAndroid(NewUAParser(p["android-oreo-8-0"])).Version(), ShouldEqual, "8.0")
				So(NewAndroid(NewUAParser(p["android-pie-9-0"])).Version(), ShouldEqual, "9")
				So(NewAndroid(NewUAParser(p["android-10"])).Version(), ShouldEqual, "10.0.0")
				So(NewAndroid(NewUAParser(p["android-11"])).Version(), ShouldEqual, "11.0")
				So(NewAndroid(NewUAParser(p["android-12"])).Version(), ShouldEqual, "12")
				So(NewAndroid(NewUAParser(p["android-20-0"])).Version(), ShouldEqual, "20.0")
				So(NewAndroid(NewUAParser(p["android-20-1"])).Version(), ShouldEqual, "20.10")
				So(NewAndroid(NewUAParser(p["android-21"])).Version(), ShouldEqual, "21")
				So(NewAndroid(NewUAParser(p["android-22"])).Version(), ShouldEqual, "22")
			})
		})
	})
}

func TestAndroidMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Android", func() {
			Convey("It should return true", func() {
				p := testPlatforms

				So(NewAndroid(NewUAParser(p["android-eclair-2-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-eclair-2-1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-froyo-2-2"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-froyo-2-2-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-gingerbread-2-3-5"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-gingerbread-2-3-6"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-gingerbread-2-3-8"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-honeycomb-3-1-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-honeycomb-3-1-4"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-honeycomb-3-2"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-ice-cream-sandwich-4-0-3"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-ice-cream-sandwich-4-0-4"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-jelly-bean-4.1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-jelly-bean-4.2"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-jelly-bean-4.3"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-kitkat-4.4"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-lollipop-5.0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-loolipop-5.1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-marshmallow-6-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-nougat-7-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-nougat-7-1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-oreo-8-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-pie-9-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-10"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-11"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-12"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-20-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-20-1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-21"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(p["android-22"])).Match(), ShouldBeTrue)
			})
		})
	})
}
