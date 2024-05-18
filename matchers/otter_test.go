package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewOtter(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewOtter(NewUAParser("")), ShouldHaveSameTypeAs, &Otter{})
	})
}

func TestOtterName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Otter", func() {
			s := NewOtter(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Otter")
		})
	})
}

func TestOtterVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("It should return the version", func() {
			ua := testUserAgents["otter"]

			So(NewOtter(NewUAParser(ua.Windows)).Version(), ShouldEqual, "1.0.02")
			So(NewOtter(NewUAParser(ua.Mac)).Version(), ShouldEqual, "0.9.11")
			So(NewOtter(NewUAParser(ua.Linux)).Version(), ShouldEqual, "0.9.08")
		})

		Convey("It should return default version", func() {
			o := NewOtter(NewUAParser(testUserAgents["alipay"].Linux))
			So(o.Version(), ShouldEqual, "0.0")
		})
	})
}

func TestOtterMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["otter"]
				So(NewOtter(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewOtter(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
				So(NewOtter(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent doesn't match", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["safari"]
				So(NewOtter(NewUAParser(ua.Windows)).Match(), ShouldBeFalse)
			})
		})
	})
}
