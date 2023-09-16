package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewQQ(t *testing.T) {
	Convey("Subject: #NewQQ", t, func() {
		Convey("It should return a new QQ instance", func() {
			So(NewQQ(""), ShouldHaveSameTypeAs, &QQ{})
		})
	})
}

func TestQQName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return QQ", func() {
			qq := NewQQ("")
			So(qq.Name(), ShouldEqual, qqName)
		})
	})
}

func TestQQVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["qq"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewQQ(ua.Windows).Version(), ShouldEqual, "9.1.3441.400")
				So(NewQQ(ua.Mac).Version(), ShouldEqual, "3.8.3858.400")
				So(NewQQ(ua.Android).Version(), ShouldEqual, "6.1")
				So(NewQQ(ua.IOS).Version(), ShouldEqual, "451")
				So(NewQQ(ua.Linux).Version(), ShouldEqual, "3.3")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				ua := testUserAgents["chrome"]
				So(NewQQ(ua.Linux).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestQQMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches QQ", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["qq"]

				So(NewQQ(ua.Windows).Match(), ShouldBeTrue)
				So(NewQQ(ua.Mac).Match(), ShouldBeTrue)
				So(NewQQ(ua.Android).Match(), ShouldBeTrue)
				So(NewQQ(ua.IOS).Match(), ShouldBeTrue)
				So(NewQQ(ua.Linux).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match QQ", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]

				So(NewQQ(ua.Linux).Match(), ShouldBeFalse)
			})
		})
	})
}
