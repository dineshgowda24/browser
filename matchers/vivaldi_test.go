package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewVivaldi(t *testing.T) {
	Convey("Subject: #NewVivaldi", t, func() {
		Convey("It should return a new Vivaldi instance", func() {
			So(NewVivaldi(NewUAParser("")), ShouldHaveSameTypeAs, &Vivaldi{})
		})
	})
}

func TestVivaldiName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the name", func() {
			So(NewVivaldi(NewUAParser("")).Name(), ShouldEqual, "Vivaldi")
		})
	})
}

func TestVivaldiVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["vivaldi"]
		Convey("When the version is captured", func() {
			Convey("It should return the version", func() {
				So(NewVivaldi(NewUAParser(ua.Linux)).Version(), ShouldEqual, "1.0.264.3")
				So(NewVivaldi(NewUAParser(ua.Mac)).Version(), ShouldEqual, "1.0.252.3")
				So(NewVivaldi(NewUAParser(ua.Windows)).Version(), ShouldEqual, "1.0.219.50")
			})
		})

		Convey("When the version is not captured", func() {
			ua := testUserAgents["chrome"]
			Convey("It should return default version", func() {
				So(NewVivaldi(NewUAParser(ua.Android)).Version(), ShouldEqual, "0.0")
				So(NewVivaldi(NewUAParser(ua.IOS)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestVivaldiMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			ua := testUserAgents["vivaldi"]
			Convey("It should return true", func() {
				So(NewVivaldi(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewVivaldi(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
				So(NewVivaldi(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			ua := testUserAgents["chrome"]
			Convey("It should return false", func() {
				So(NewVivaldi(NewUAParser(ua.Android)).Match(), ShouldBeFalse)
				So(NewVivaldi(NewUAParser(ua.IOS)).Match(), ShouldBeFalse)
			})
		})
	})
}
