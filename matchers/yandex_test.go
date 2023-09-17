package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewYandex(t *testing.T) {
	Convey("Subject: #NewYandex", t, func() {
		Convey("It should return a new Yandex instance", func() {
			So(NewYandex(""), ShouldHaveSameTypeAs, &Yandex{})
		})
	})
}

func TestYandexName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Yandex", func() {
			yandex := NewYandex("")
			So(yandex.Name(), ShouldEqual, "Yandex")
		})
	})
}

func TestYandexVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				yandex := testUserAgents["yandex"]

				So(NewYandex(yandex.Android).Version(), ShouldEqual, "15.4.2272.3842.01")
				So(NewYandex(yandex.IOS).Version(), ShouldEqual, "15.9.2403.2772.10")
				So(NewYandex(yandex.Linux).Version(), ShouldEqual, "15.9.2403.2150")
				So(NewYandex(yandex.Mac).Version(), ShouldEqual, "15.4.2272.3909")
				So(NewYandex(yandex.Windows).Version(), ShouldEqual, "15.6.2311.5029")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return 0.0", func() {
				yandex := NewYandex(testUserAgents["chrome"].Linux)
				So(yandex.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestYandexMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Yandex", func() {
			Convey("It should return true", func() {
				yandex := testUserAgents["yandex"]

				So(NewYandex(yandex.Android).Match(), ShouldBeTrue)
				So(NewYandex(yandex.IOS).Match(), ShouldBeTrue)
				So(NewYandex(yandex.Linux).Match(), ShouldBeTrue)
				So(NewYandex(yandex.Mac).Match(), ShouldBeTrue)
				So(NewYandex(yandex.Windows).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Yandex", func() {
			Convey("It should return false", func() {
				yandex := NewYandex(testUserAgents["chrome"].Linux)
				So(yandex.Match(), ShouldBeFalse)
			})
		})
	})
}
