package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewKonqueror(t *testing.T) {
	Convey("Subject: #NewKonqueror", t, func() {
		Convey("It should return a new Konqueror instance", func() {
			So(NewKonqueror(""), ShouldHaveSameTypeAs, &Konqueror{})
		})
	})
}

func TestKonquerorName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Konqueror", func() {
			k := NewKonqueror("")
			So(k.Name(), ShouldEqual, konquerorName)
		})
	})
}

func TestKonquerorVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["konqueror"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewKonqueror(ua.Linux).Version(), ShouldEqual, "4.14.1")
				So(NewKonqueror(ua.Mac).Version(), ShouldEqual, "5.0")
				So(NewKonqueror(ua.Windows).Version(), ShouldEqual, "4.10")
				So(NewKonqueror(ua.Android).Version(), ShouldEqual, "10.1")
				So(NewKonqueror(ua.IOS).Version(), ShouldEqual, "14.1")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				k := NewKonqueror(testUserAgents["chrome"].Linux)
				So(k.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestKonquerorMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Konqueror", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["konqueror"]
				So(NewKonqueror(ua.Linux).Match(), ShouldBeTrue)
				So(NewKonqueror(ua.Mac).Match(), ShouldBeTrue)
				So(NewKonqueror(ua.Windows).Match(), ShouldBeTrue)
				So(NewKonqueror(ua.Android).Match(), ShouldBeTrue)
				So(NewKonqueror(ua.IOS).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Konqueror", func() {
			Convey("It should return false", func() {
				So(NewKonqueror(testUserAgents["chrome"].Linux).Match(), ShouldBeFalse)
			})
		})
	})
}
