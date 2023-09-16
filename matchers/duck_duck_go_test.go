package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewDuckDuckGo(t *testing.T) {
	Convey("Subject: #NewDuckDuckGo", t, func() {
		Convey("When the given user agent is empty", func() {
			Convey("It should return a new DuckDuckGo instance", func() {
				So(NewDuckDuckGo(""), ShouldHaveSameTypeAs, &DuckDuckGo{})
			})
		})
	})
}

func TestDuckDuckGoName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			ddg := NewDuckDuckGo("")
			So(ddg.Name(), ShouldEqual, duckDuckGoName)
		})
	})
}

func TestDuckDuckGoVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				ua := testUserAgents["duck_duck_go"]

				So(NewDuckDuckGo(ua.IOS).Version(), ShouldEqual, "2")
				So(NewDuckDuckGo(ua.Android).Version(), ShouldEqual, "5")
				So(NewDuckDuckGo(ua.Windows).Version(), ShouldEqual, "1.0")
				So(NewDuckDuckGo(ua.Mac).Version(), ShouldEqual, "1.0.1")
				So(NewDuckDuckGo(ua.Linux).Version(), ShouldEqual, "5")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				ddg := NewDuckDuckGo("Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/605.1.15")
				So(ddg.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestDuckDuckGoMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["duck_duck_go"]

				So(NewDuckDuckGo(ua.IOS).Match(), ShouldBeTrue)
				So(NewDuckDuckGo(ua.Android).Match(), ShouldBeTrue)
				So(NewDuckDuckGo(ua.Windows).Match(), ShouldBeTrue)
				So(NewDuckDuckGo(ua.Mac).Match(), ShouldBeTrue)
				So(NewDuckDuckGo(ua.Linux).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				ddg := NewDuckDuckGo("Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/605.1.15")
				So(ddg.Match(), ShouldBeFalse)
			})
		})
	})
}
