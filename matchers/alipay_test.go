package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAlipay(t *testing.T) {
	Convey("Subject: #NewAlipay", t, func() {
		Convey("It should return a new Alipay instance", func() {
			alipay := NewAlipay(NewUAParser(""))
			So(alipay, ShouldHaveSameTypeAs, &Alipay{})
		})
	})
}

func TestAlipayName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			alipay := NewAlipay(NewUAParser(""))
			So(alipay.Name(), ShouldEqual, "Alipay")
		})
	})
}

func TestAlipayVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is captured", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["alipay"]

				So(NewAlipay(NewUAParser(ua.IOS)).Version(), ShouldEqual, "9.9.0.000001")
				So(NewAlipay(NewUAParser(ua.Android)).Version(), ShouldEqual, "1.7.0.123008")
			})
		})

		Convey("When the version is not captured", func() {
			Convey("It should return default version", func() {
				ua := testUserAgents["alipay"]

				So(NewAlipay(NewUAParser(ua.Windows)).Version(), ShouldEqual, "0.0")
				So(NewAlipay(NewUAParser(ua.Mac)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestAlipayMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Alipay", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["alipay"]
				So(NewAlipay(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
				So(NewAlipay(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
				So(NewAlipay(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewAlipay(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Alipay", func() {
			Convey("It should return false", func() {
				ua := "Mozilla/5.0 (Linux; Android 10; Redmi K30 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Mobile Safari/537.36"
				alipay := NewAlipay(NewUAParser(ua))
				So(alipay.Match(), ShouldBeFalse)
			})
		})
	})
}
