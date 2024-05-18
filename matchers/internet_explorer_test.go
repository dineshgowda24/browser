package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInternetExplorer(t *testing.T) {
	Convey("Subject: #NewInternetExplorer", t, func() {
		Convey("It should return a new InternetExplorer instance", func() {
			So(NewInternetExplorer(NewUAParser("")), ShouldHaveSameTypeAs, &InternetExplorer{})
		})
	})
}

func TestInternetExplorerName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			ie := NewInternetExplorer(NewUAParser(""))
			So(ie.Name(), ShouldEqual, "Internet Explorer")
		})
	})
}

func TestInternetExplorerVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ie := testUserAgents["ie"]
		oldIE := testUserAgents["old-ie"]

		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				So(NewInternetExplorer(NewUAParser(ie.Windows)).Version(), ShouldEqual, "11")
				So(NewInternetExplorer(NewUAParser(oldIE.Mac)).Version(), ShouldEqual, "10")
				So(NewInternetExplorer(NewUAParser(oldIE.Windows)).Version(), ShouldEqual, "11")
			})
		})
	})
}

func TestInternetExplorerMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		ie := testUserAgents["ie"]
		oldIE := testUserAgents["old-ie"]
		chrome := testUserAgents["chrome"]

		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewInternetExplorer(NewUAParser(ie.Windows)).Match(), ShouldBeTrue)
				So(NewInternetExplorer(NewUAParser(oldIE.Mac)).Match(), ShouldBeTrue)
				So(NewInternetExplorer(NewUAParser(oldIE.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewInternetExplorer(NewUAParser(chrome.Mac)).Match(), ShouldBeFalse)
				So(NewInternetExplorer(NewUAParser(chrome.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
