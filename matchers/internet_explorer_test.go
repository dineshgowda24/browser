package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInternetExplorer(t *testing.T) {
	Convey("Subject: #NewInternetExplorer", t, func() {
		Convey("It should return a new InternetExplorer instance", func() {
			So(NewInternetExplorer(""), ShouldHaveSameTypeAs, &InternetExplorer{})
		})
	})
}

func TestInternetExplorerName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			ie := NewInternetExplorer("")
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

				So(NewInternetExplorer(ie.Windows).Version(), ShouldEqual, "11")
				So(NewInternetExplorer(oldIE.Mac).Version(), ShouldEqual, "10")
				So(NewInternetExplorer(oldIE.Windows).Version(), ShouldEqual, "11")
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
				So(NewInternetExplorer(ie.Windows).Match(), ShouldBeTrue)
				So(NewInternetExplorer(oldIE.Mac).Match(), ShouldBeTrue)
				So(NewInternetExplorer(oldIE.Windows).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewInternetExplorer(chrome.Mac).Match(), ShouldBeFalse)
				So(NewInternetExplorer(chrome.Linux).Match(), ShouldBeFalse)
			})
		})
	})
}
