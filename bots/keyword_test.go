package bots

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewKeyword(t *testing.T) {
	Convey("Subject: #NewKeyword", t, func() {
		Convey("It returns a Keyword bot", func() {
			k := NewKeyword("Googlebot")
			So(k, ShouldHaveSameTypeAs, &Keyword{})
		})
	})
}

func TestKeywordName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It returns the keyword that matched", func() {
			bots := []string{"crawl", "spider", "fetch", "datapack", "monitoring", "googlebot"}

			for _, v := range bots {
				k := NewKeyword(testBots[v])
				So(k.Name(), ShouldEqual, "Generic Bot")
			}
		})
	})
}

func TestKeywordMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent contains a keyword", func() {
			Convey("It returns true", func() {
				bots := []string{"crawl", "spider", "fetch", "datapack", "monitoring", "googlebot"}

				for _, b := range bots {
					So(NewKeyword(testBots[b]).Match(), ShouldBeTrue)
				}
			})
		})
	})
}
