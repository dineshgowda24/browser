package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewGSA(t *testing.T) {
	Convey("Subject: #NewGSA", t, func() {
		Convey("It should return a new GSA instance", func() {
			So(NewGoogleSearchApp(""), ShouldHaveSameTypeAs, &GoogleSearchApp{})
		})
	})
}

func TestGSAName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Google Search App", func() {
			gsa := NewGoogleSearchApp("")
			So(gsa.Name(), ShouldEqual, googleSearchAppName)
		})
	})
}

func TestGSAVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["gsa"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {

				So(NewGoogleSearchApp(ua.Linux).Version(), ShouldEqual, "11.2.9.21")
				So(NewGoogleSearchApp(ua.Mac).Version(), ShouldEqual, "7.0.55539")
				So(NewGoogleSearchApp(ua.Android).Version(), ShouldEqual, "5.10.32.19")
				So(NewGoogleSearchApp(ua.IOS).Version(), ShouldEqual, "7.0.55539")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewGoogleSearchApp(ua.Windows).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestGSAMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches GSA", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["gsa"]

				So(NewGoogleSearchApp(ua.Linux).Match(), ShouldBeTrue)
				So(NewGoogleSearchApp(ua.Mac).Match(), ShouldBeTrue)
				So(NewGoogleSearchApp(ua.Android).Match(), ShouldBeTrue)
				So(NewGoogleSearchApp(ua.IOS).Match(), ShouldBeTrue)
				So(NewGoogleSearchApp(ua.Windows).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match GSA", func() {
			Convey("It should return false", func() {
				userAgent := "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Mobile Safari/537.36"

				So(NewGoogleSearchApp(userAgent).Match(), ShouldBeFalse)
			})
		})
	})
}
