package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewYandex(t *testing.T) {
	Convey("Subject: #NewYandex", t, func() {
		Convey("It should return a new Yandex instance", func() {
			So(NewYandex(NewUAParser("")), ShouldHaveSameTypeAs, &Yandex{})
		})
	})
}

func TestYandexName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Yandex", func() {
			yandex := NewYandex(NewUAParser(""))
			So(yandex.Name(), ShouldEqual, "Yandex")
		})
	})
}

func TestYandexVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				yandex := testUserAgents["yandex"]

				So(NewYandex(NewUAParser(yandex.Android)).Version(), ShouldEqual, "15.4.2272.3842.01")
				So(NewYandex(NewUAParser(yandex.IOS)).Version(), ShouldEqual, "15.9.2403.2772.10")
				So(NewYandex(NewUAParser(yandex.Linux)).Version(), ShouldEqual, "15.9.2403.2150")
				So(NewYandex(NewUAParser(yandex.Mac)).Version(), ShouldEqual, "15.4.2272.3909")
				So(NewYandex(NewUAParser(yandex.Windows)).Version(), ShouldEqual, "15.6.2311.5029")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return 0.0", func() {
				yandex := NewYandex(NewUAParser(testUserAgents["chrome"].Linux))
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

				So(NewYandex(NewUAParser(yandex.Android)).Match(), ShouldBeTrue)
				So(NewYandex(NewUAParser(yandex.IOS)).Match(), ShouldBeTrue)
				So(NewYandex(NewUAParser(yandex.Linux)).Match(), ShouldBeTrue)
				So(NewYandex(NewUAParser(yandex.Mac)).Match(), ShouldBeTrue)
				So(NewYandex(NewUAParser(yandex.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Yandex", func() {
			Convey("It should return false", func() {
				yandex := NewYandex(NewUAParser(testUserAgents["chrome"].Linux))
				So(yandex.Match(), ShouldBeFalse)
			})
		})
	})
}
