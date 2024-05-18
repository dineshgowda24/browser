package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSputnik(t *testing.T) {
	Convey("Subject: #NewSputnik", t, func() {
		Convey("It should return a new Sputnik instance", func() {
			So(NewSputnik(NewUAParser("")), ShouldHaveSameTypeAs, &Sputnik{})
		})
	})
}

func TestSputnikName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Sputnik", func() {
			s := NewSputnik(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Sputnik")
		})
	})
}

func TestSputnikVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				s := testUserAgents["sputnik"]

				So(NewSputnik(NewUAParser(s.Android)).Version(), ShouldEqual, "0.6.1")
				So(NewSputnik(NewUAParser(s.IOS)).Version(), ShouldEqual, "2.3.8")
				So(NewSputnik(NewUAParser(s.Linux)).Version(), ShouldEqual, "3.3.2038.2")
				So(NewSputnik(NewUAParser(s.Windows)).Version(), ShouldEqual, "1.9.48.2")
			})
		})
	})
}

func TestSputnikMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Sputnik", func() {
			Convey("It should return true", func() {
				s := testUserAgents["sputnik"]

				So(NewSputnik(NewUAParser(s.Android)).Match(), ShouldBeTrue)
				So(NewSputnik(NewUAParser(s.IOS)).Match(), ShouldBeTrue)
				So(NewSputnik(NewUAParser(s.Linux)).Match(), ShouldBeTrue)
				So(NewSputnik(NewUAParser(s.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Sputnik", func() {
			Convey("It should return false", func() {
				s := testUserAgents["chrome"]

				So(NewSputnik(NewUAParser(s.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
