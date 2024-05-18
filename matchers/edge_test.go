package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewEdge(t *testing.T) {
	Convey("Subject: #NewEdge", t, func() {
		Convey("It should return an Edge instance", func() {
			So(NewEdge(NewUAParser("")), ShouldHaveSameTypeAs, &Edge{})
		})
	})
}

func TestEdgeName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct browser name", func() {
			So(NewEdge(NewUAParser("")).Name(), ShouldEqual, "Microsoft Edge")
		})
	})
}

func TestEdgeVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("It should return the correct browser version", func() {
			ua := testUserAgents["edge"]

			So(NewEdge(NewUAParser(ua.Windows)).Version(), ShouldEqual, "12.10240")
			So(NewEdge(NewUAParser(ua.Android)).Version(), ShouldEqual, "14.14393")
		})
	})
}

func TestEdgeMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["edge"]

				So(NewEdge(NewUAParser(ua.Windows)).Match(), ShouldBeTrue)
				So(NewEdge(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent doesn't match", func() {
			Convey("It should return false", func() {
				So(NewEdge(NewUAParser(testUserAgents["chrome"].Android)).Match(), ShouldBeFalse)
			})
		})
	})
}
