package bots

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewEmpty(t *testing.T) {
	Convey("Subject: #NewEmpty", t, func() {
		Convey("It should return an empty bot", func() {
			So(NewEmpty(testBots["googlebot"]), ShouldHaveSameTypeAs, &Empty{})
		})
	})
}

func TestEmptyName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Generic Bot", func() {
			So(NewEmpty(testBots["googlebot"]).Name(), ShouldEqual, "Generic Bot")
		})
	})
}

func TestEmptyMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent is empty", func() {
			Convey("It should return true", func() {
				So(NewEmpty("").Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent is not empty", func() {
			Convey("It should return false", func() {
				So(NewEmpty(testBots["googlebot"]).Match(), ShouldBeFalse)
			})
		})
	})
}
