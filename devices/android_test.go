package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAndroid(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewAndroid(NewUAParser("")), ShouldHaveSameTypeAs, &Android{})
	})
}

func TestAndroidName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Android", func() {
			So(NewAndroid(NewUAParser(testDevices["android-eclair-2-0"])).Name(), ShouldEqual, "HUAWEI XXGW")
			So(NewAndroid(NewUAParser(testDevices["android-eclair-2-1"])).Name(), ShouldEqual, "Xperia Ray")
			So(NewAndroid(NewUAParser(testDevices["android-froyo-2-2"])).Name(), ShouldEqual, "Nexus One")
			So(NewAndroid(NewUAParser(testDevices["android-froyo-2-2-0"])).Name(), ShouldEqual, "XELECTRON WS777")
			So(NewAndroid(NewUAParser(testDevices["android-gingerbread-2-3-5"])).Name(), ShouldEqual, "HTC Desire HD A9191")
			So(NewAndroid(NewUAParser(testDevices["android-gingerbread-2-3-6"])).Name(), ShouldEqual, "IM-A760S")
			So(NewAndroid(NewUAParser(testDevices["android-gingerbread-2-3-8"])).Name(), ShouldEqual, "Huawei X3")
			So(NewAndroid(NewUAParser(testDevices["android-honeycomb-3-1-0"])).Name(), ShouldEqual, "SGH-M919")
			So(NewAndroid(NewUAParser(testDevices["android-honeycomb-3-1-4"])).Name(), ShouldEqual, "Nexus 12k")
			So(NewAndroid(NewUAParser(testDevices["android-honeycomb-3-2"])).Name(), ShouldEqual, "SAMSUNG-SGH-I957")
			So(NewAndroid(NewUAParser(testDevices["android-ice-cream-sandwich-4-0-3"])).Name(), ShouldEqual, "HTC EVO 3D X515m")
			So(NewAndroid(NewUAParser(testDevices["android-ice-cream-sandwich-4-0-4"])).Name(), ShouldEqual, "LT18i")
			So(NewAndroid(NewUAParser(testDevices["android-jelly-bean-4.1"])).Name(), ShouldEqual, "MI 2S")
			So(NewAndroid(NewUAParser(testDevices["android-jelly-bean-4.2"])).Name(), ShouldEqual, "GT-I9505")
			So(NewAndroid(NewUAParser(testDevices["android-jelly-bean-4.3"])).Name(), ShouldEqual, "HM 1SW")
			So(NewAndroid(NewUAParser(testDevices["android-kitkat-4.4"])).Name(), ShouldEqual, "sm-n9005")
			So(NewAndroid(NewUAParser(testDevices["android-lollipop-5.0"])).Name(), ShouldEqual, "Redmi Note 2")
			So(NewAndroid(NewUAParser(testDevices["android-loolipop-5.1"])).Name(), ShouldEqual, "Unknown")
			So(NewAndroid(NewUAParser(testDevices["android-marshmallow-6-0"])).Name(), ShouldEqual, "MI 4LTE")
			So(NewAndroid(NewUAParser(testDevices["android-nougat-7-0"])).Name(), ShouldEqual, "Nexus 5x")
			So(NewAndroid(NewUAParser(testDevices["android-nougat-7-1"])).Name(), ShouldEqual, "SM-G935V")
			So(NewAndroid(NewUAParser(testDevices["android-oreo-8-0"])).Name(), ShouldEqual, "Pixel XL")
			So(NewAndroid(NewUAParser(testDevices["android-pie-9-0"])).Name(), ShouldEqual, "Pixel XL")
			So(NewAndroid(NewUAParser(testDevices["android-10"])).Name(), ShouldEqual, "Nexus 7")
			So(NewAndroid(NewUAParser(testDevices["android-11"])).Name(), ShouldEqual, "LG-D335")
			So(NewAndroid(NewUAParser(testDevices["android-12"])).Name(), ShouldEqual, "Pixel 3a")
			So(NewAndroid(NewUAParser(testDevices["android-20-0"])).Name(), ShouldEqual, "Unknown")
			So(NewAndroid(NewUAParser(testDevices["android-20-1"])).Name(), ShouldEqual, "SM-J100H")
			So(NewAndroid(NewUAParser(testDevices["android-21"])).Name(), ShouldEqual, "SM-G360M")
			So(NewAndroid(NewUAParser(testDevices["android-22"])).Name(), ShouldEqual, "Z818L")
		})
	})
}

func TestAndroidMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewAndroid(NewUAParser(testDevices["android-eclair-2-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-eclair-2-1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-froyo-2-2"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-froyo-2-2-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-gingerbread-2-3-5"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-gingerbread-2-3-6"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-gingerbread-2-3-8"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-honeycomb-3-1-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-honeycomb-3-1-4"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-honeycomb-3-2"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-ice-cream-sandwich-4-0-3"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-ice-cream-sandwich-4-0-4"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-jelly-bean-4.1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-jelly-bean-4.2"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-jelly-bean-4.3"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-kitkat-4.4"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-lollipop-5.0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-loolipop-5.1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-marshmallow-6-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-nougat-7-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-nougat-7-1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-oreo-8-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-pie-9-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-10"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-11"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-12"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-20-0"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-20-1"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-21"])).Match(), ShouldBeTrue)
				So(NewAndroid(NewUAParser(testDevices["android-22"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent doesn't match", func() {
			Convey("It should return false", func() {
				So(NewAndroid(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
