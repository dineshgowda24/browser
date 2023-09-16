package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewElectron(t *testing.T) {
	Convey("Subject: #NewElectron", t, func() {
		Convey("It should return a new Electron instance", func() {
			So(NewElectron(""), ShouldHaveSameTypeAs, &Electron{})
		})
	})
}

func TestElectronName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the name", func() {
			e := NewElectron("")
			So(e.Name(), ShouldEqual, electronName)
		})
	})
}

func TestElectronVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["electron"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {

				So(NewElectron(ua.Linux).Version(), ShouldEqual, "0.30.7")
				So(NewElectron(ua.Windows).Version(), ShouldEqual, "0.36.4")
				So(NewElectron(ua.Android).Version(), ShouldEqual, "7.1.2")
				So(NewElectron(ua.Mac).Version(), ShouldEqual, "0.36.9")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {

				So(NewElectron(ua.IOS).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestElectronMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Electron", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["electron"]

				So(NewElectron(ua.Linux).Match(), ShouldBeTrue)
				So(NewElectron(ua.Windows).Match(), ShouldBeTrue)
				So(NewElectron(ua.Mac).Match(), ShouldBeTrue)
				So(NewElectron(ua.Android).Match(), ShouldBeTrue)
				So(NewElectron(ua.IOS).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Electron", func() {
			Convey("It should return false", func() {
				e := NewElectron("Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko)")
				So(e.Match(), ShouldBeFalse)
			})
		})
	})
}
