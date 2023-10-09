package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUnknown(t *testing.T) {
	Convey("Subject: #NewUnknown", t, func() {
		Convey("Should return a new Unknown instance", func() {
			So(NewUnknown(""), ShouldHaveSameTypeAs, &Unknown{})
		})
	})
}

func TestUnknownName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("Should return the name of the unknown browser", func() {

			ua := testUserAgents["unknown"]
			So(NewUnknown(ua.IOS).Name(), ShouldEqual, "Apple CoreMedia")
			So(NewUnknown(ua.Mac).Name(), ShouldEqual, "QuickTime")
			So(NewUnknown("").Name(), ShouldEqual, "Unknown Browser")
		})
	})
}

func TestUnknownVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("Should return the version of the unknown browser", func() {

			ua := testUserAgents["unknown"]
			So(NewUnknown(ua.IOS).Version(), ShouldEqual, "1.0.0.11")
			So(NewUnknown(ua.Mac).Version(), ShouldEqual, "0.0")
			So(NewUnknown("").Version(), ShouldEqual, "0.0")
		})
	})
}

func TestUnknownMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("Should return true", func() {
			So(NewUnknown("").Match(), ShouldBeTrue)
		})
	})
}
