package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSputnik(t *testing.T) {
	Convey("Subject: #NewSputnik", t, func() {
		Convey("It should return a new Sputnik instance", func() {
			So(NewSputnik(""), ShouldHaveSameTypeAs, &Sputnik{})
		})
	})
}

func TestSputnikName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Sputnik", func() {
			s := NewSputnik("")
			So(s.Name(), ShouldEqual, "Sputnik")
		})
	})
}

func TestSputnikVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				s := testUserAgents["sputnik"]

				So(NewSputnik(s.Android).Version(), ShouldEqual, "0.6.1")
				So(NewSputnik(s.IOS).Version(), ShouldEqual, "2.3.8")
				So(NewSputnik(s.Linux).Version(), ShouldEqual, "3.3.2038.2")
				So(NewSputnik(s.Windows).Version(), ShouldEqual, "1.9.48.2")
			})
		})
	})
}

func TestSputnikMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Sputnik", func() {
			Convey("It should return true", func() {
				s := testUserAgents["sputnik"]

				So(NewSputnik(s.Android).Match(), ShouldBeTrue)
				So(NewSputnik(s.IOS).Match(), ShouldBeTrue)
				So(NewSputnik(s.Linux).Match(), ShouldBeTrue)
				So(NewSputnik(s.Windows).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Sputnik", func() {
			Convey("It should return false", func() {
				s := testUserAgents["chrome"]

				So(NewSputnik(s.Linux).Match(), ShouldBeFalse)
			})
		})
	})
}
