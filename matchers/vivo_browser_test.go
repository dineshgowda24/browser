package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewVivoBrowser(t *testing.T) {
	Convey("Subject: #NewVivoBrowser", t, func() {
		Convey("It should return a new VivoBrowser instance", func() {
			So(NewVivoBrowser(""), ShouldHaveSameTypeAs, &VivoBrowser{})
		})
	})
}

func TestVivoBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			vb := NewVivoBrowser("")
			So(vb.Name(), ShouldEqual, "Vivo Browser")
		})
	})
}

func TestVivoBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				vb := testUserAgents["vivo-browser"]

				So(NewVivoBrowser(vb.Android).Version(), ShouldEqual, "5.0.10")
				So(NewVivoBrowser(vb.IOS).Version(), ShouldEqual, "5.0.10")
				So(NewVivoBrowser(vb.Linux).Version(), ShouldEqual, "5.0.2")
				So(NewVivoBrowser(vb.Windows).Version(), ShouldEqual, "5.4.11")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return 0.0", func() {
				vb := NewVivoBrowser(testUserAgents["chrome"].Linux)
				So(vb.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestVivoBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches", func() {
			Convey("It should return true", func() {
				vb := testUserAgents["vivo-browser"]

				So(NewVivoBrowser(vb.Android).Match(), ShouldBeTrue)
				So(NewVivoBrowser(vb.IOS).Match(), ShouldBeTrue)
				So(NewVivoBrowser(vb.Linux).Match(), ShouldBeTrue)
				So(NewVivoBrowser(vb.Windows).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"].Linux
				So(NewVivoBrowser(ua).Match(), ShouldBeFalse)
			})
		})
	})
}
