package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWeibo(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewWeibo(NewUAParser("")), ShouldHaveSameTypeAs, &Weibo{})
	})
}

func TestWeiboName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Weibo", func() {
			s := NewWeibo(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Weibo")
		})
	})
}

func TestWeiboVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("It should return default version", func() {
			ua := testUserAgents["weibo"]

			So(NewWeibo(NewUAParser(ua.Android)).Version(), ShouldEqual, "6.5.0")
			So(NewWeibo(NewUAParser(ua.IOS)).Version(), ShouldEqual, "10.9.2")
		})
	})
}

func TestWeiboMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["weibo"]
				So(NewWeibo(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewWeibo(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent doesn't match", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["safari"]
				So(NewWeibo(NewUAParser(ua.Android)).Match(), ShouldBeFalse)
			})
		})
	})
}
