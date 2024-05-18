package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBrave(t *testing.T) {
	Convey("Subject: #NewBrave", t, func() {
		Convey("It should return a new Brave instance", func() {
			So(NewBrave(NewUAParser("")), ShouldHaveSameTypeAs, &Brave{})
		})
	})
}

func TestBraveName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the name", func() {
			So(NewBrave(NewUAParser("")).Name(), ShouldEqual, "Brave")
		})
	})
}

func TestBraveVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["brave"]
		Convey("When the version is captured", func() {
			Convey("It should return the version", func() {
				So(NewBrave(NewUAParser(ua.Windows)).Version(), ShouldEqual, "0.7.7")
				So(NewBrave(NewUAParser(ua.Linux)).Version(), ShouldEqual, "0.7.7")
				So(NewBrave(NewUAParser(ua.Mac)).Version(), ShouldEqual, "0.7.7")
			})
		})

		Convey("When the version is not captured", func() {
			ua := testUserAgents["chrome"]
			Convey("It should return default version", func() {
				So(NewBrave(NewUAParser(ua.Android)).Version(), ShouldEqual, "0.0")
				So(NewBrave(NewUAParser(ua.IOS)).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestBraveMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			ua := testUserAgents["brave"]
			Convey("It should return true", func() {
				So(NewBrave(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewBrave(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
				So(NewBrave(NewUAParser(ua.Mac)).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			ua := testUserAgents["chrome"]
			Convey("It should return false", func() {
				So(NewBrave(NewUAParser(ua.Android)).Match(), ShouldBeFalse)
				So(NewBrave(NewUAParser(ua.IOS)).Match(), ShouldBeFalse)
			})
		})
	})
}
