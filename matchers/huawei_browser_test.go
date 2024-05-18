package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewHuaweiBrowser(t *testing.T) {
	Convey("Subject: #NewHuaweiBrowser", t, func() {
		Convey("It should return a new HuaweiBrowser instance", func() {
			So(NewHuaweiBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &HuaweiBrowser{})
		})
	})
}

func TestHuaweiBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			hb := NewHuaweiBrowser(NewUAParser(""))
			So(hb.Name(), ShouldEqual, huaweiBrowserName)
		})
	})
}

func TestHuaweiBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				ua := testUserAgents["huawei-browser"]

				So(NewHuaweiBrowser(NewUAParser(ua.Windows)).Version(), ShouldEqual, "12.1.3.303")
				So(NewHuaweiBrowser(NewUAParser(ua.Android)).Version(), ShouldEqual, "14.0.2.300")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				ua := testUserAgents["uc-browser"]
				So(NewHuaweiBrowser(NewUAParser(ua.Android)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestHuaweiBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches HuaweiBrowser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["huawei-browser"]
				So(NewHuaweiBrowser(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewHuaweiBrowser(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match HuaweiBrowser", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["uc-browser"]
				So(NewHuaweiBrowser(NewUAParser(ua.Android)).Match(), ShouldBeFalse)
			})
		})
	})
}
