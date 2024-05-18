package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSnapchat(t *testing.T) {
	Convey("Subject: #NewSnapchat", t, func() {
		Convey("It should return a new Snapchat instance", func() {
			So(NewSnapchat(NewUAParser("")), ShouldHaveSameTypeAs, &Snapchat{})
		})
	})
}

func TestSnapchatName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Snapchat", func() {
			s := NewSnapchat(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Snapchat")
		})
	})
}

func TestSnapchatVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				s := testUserAgents["snapchat"]

				So(NewSnapchat(NewUAParser(s.Android)).Version(), ShouldEqual, "10.18.2.0")
				So(NewSnapchat(NewUAParser(s.IOS)).Version(), ShouldEqual, "10.22.2.0")
				So(NewSnapchat(NewUAParser(s.Windows)).Version(), ShouldEqual, "10.72.5.0")
			})
		})
	})
}

func TestSnapchatMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Snapchat", func() {
			Convey("It should return true", func() {
				s := testUserAgents["snapchat"]

				So(NewSnapchat(NewUAParser(s.Android)).Match(), ShouldBeTrue)
				So(NewSnapchat(NewUAParser(s.IOS)).Match(), ShouldBeTrue)
				So(NewSnapchat(NewUAParser(s.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Snapchat", func() {
			Convey("It should return false", func() {
				s := testUserAgents["chrome"]

				So(NewSnapchat(NewUAParser(s.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
