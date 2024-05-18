package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewNokia(t *testing.T) {
	Convey("Subject: #NewNokia", t, func() {
		Convey("It should return a new Nokia instance", func() {
			So(NewNokia(NewUAParser("")), ShouldHaveSameTypeAs, &Nokia{})
		})
	})
}

func TestNokiaName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Nokia S40 Ovi Browser", func() {
			n := NewNokia(NewUAParser(""))
			So(n.Name(), ShouldEqual, "Nokia S40 Ovi Browser")
		})
	})
}

func TestNokiaVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["nokia"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewNokia(NewUAParser(ua.Android)).Version(), ShouldEqual, "1.2.0.12")
				So(NewNokia(NewUAParser(ua.Linux)).Version(), ShouldEqual, "2.3.0.0.48")
			})
		})
	})

	Convey("When the version is not matched", t, func() {
		Convey("It should return default version", func() {
			n := NewNokia(NewUAParser(testUserAgents["chrome"].Linux))
			So(n.Version(), ShouldEqual, "0.0")
		})
	})
}

func TestNokiaMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Nokia", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["nokia"]

				So(NewNokia(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewNokia(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Nokia", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]

				So(NewNokia(NewUAParser(ua.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
