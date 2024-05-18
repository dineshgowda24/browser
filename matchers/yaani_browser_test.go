package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewYaaniBrowser(t *testing.T) {
	Convey("Subject: #NewYaaniBrowser", t, func() {
		Convey("It should return a new YaaniBrowser instance", func() {
			So(NewYaaniBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &YaaniBrowser{})
		})
	})
}

func TestYaaniBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			yb := NewYaaniBrowser(NewUAParser(""))
			So(yb.Name(), ShouldEqual, "Yaani Browser")
		})
	})
}

func TestYaaniBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				yb := testUserAgents["yaani-browser"]

				So(NewYaaniBrowser(NewUAParser(yb.Android)).Version(), ShouldEqual, "3.3.0.301")
				So(NewYaaniBrowser(NewUAParser(yb.IOS)).Version(), ShouldEqual, "3.6.355")
				So(NewYaaniBrowser(NewUAParser(yb.Linux)).Version(), ShouldEqual, "3.5.0.441")
				So(NewYaaniBrowser(NewUAParser(yb.Mac)).Version(), ShouldEqual, "5.1.2")
				So(NewYaaniBrowser(NewUAParser(yb.Windows)).Version(), ShouldEqual, "0.1")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return 0.0", func() {
				yb := NewYaaniBrowser(NewUAParser(testUserAgents["blackberry"].Windows))
				So(yb.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestYaaniBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches", func() {
			Convey("It should return true", func() {
				yb := testUserAgents["yaani-browser"]

				So(NewYaaniBrowser(NewUAParser(yb.Android)).Match(), ShouldBeTrue)
				So(NewYaaniBrowser(NewUAParser(yb.IOS)).Match(), ShouldBeTrue)
				So(NewYaaniBrowser(NewUAParser(yb.Linux)).Match(), ShouldBeTrue)
				So(NewYaaniBrowser(NewUAParser(yb.Mac)).Match(), ShouldBeTrue)
				So(NewYaaniBrowser(NewUAParser(yb.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"].Linux

				So(NewYaaniBrowser(NewUAParser(ua)).Match(), ShouldBeFalse)
			})
		})
	})
}
