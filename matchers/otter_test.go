package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewOtter(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewOtter(""), ShouldHaveSameTypeAs, &Otter{})
	})
}

func TestOtterName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Otter", func() {
			s := NewOtter("")
			So(s.Name(), ShouldEqual, "Otter")
		})
	})
}

func TestOtterVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("It should return the version", func() {
			ua := testUserAgents["otter"]

			So(NewOtter(ua.Windows).Version(), ShouldEqual, "1.0.02")
			So(NewOtter(ua.Mac).Version(), ShouldEqual, "0.9.11")
			So(NewOtter(ua.Linux).Version(), ShouldEqual, "0.9.08")
		})

		Convey("It should return default version", func() {
			o := NewOtter(testUserAgents["alipay"].Linux)
			So(o.Version(), ShouldEqual, "0.0")
		})
	})
}

func TestOtterMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["otter"]
				So(NewOtter(ua.Windows).Match(), ShouldBeTrue)
				So(NewOtter(ua.Mac).Match(), ShouldBeTrue)
				So(NewOtter(ua.Linux).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent doesn't match", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["safari"]
				So(NewOtter(ua.Windows).Match(), ShouldBeFalse)
			})
		})
	})
}
