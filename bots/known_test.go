package bots

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewKnown(t *testing.T) {
	Convey("Subject: #NewKnown", t, func() {
		Convey("It returns a Known bot", func() {
			k := NewKnown("Googlebot", nil)
			So(k, ShouldHaveSameTypeAs, &Known{})
		})
	})
}

func TestKnownName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It returns the name of matched bot", func() {
			bots := map[string]string{
				"apis-google":       "APIs-Google",
				"apache-httpclient": "Java http library",
				"barkrowler":        "Barkrowler",
				"bingbot":           "Microsoft Bing",
			}

			for key, val := range bots {
				k := NewKnown(testBots[key], bots)

				So(k.Name(), ShouldEqual, val)
			}
		})
	})
}

func TestKnownMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		known := map[string]string{
			"apis-google":       "APIs-Google",
			"apache-httpclient": "Java http library",
			"barkrowler":        "Barkrowler",
			"bingbot":           "Microsoft Bing",
		}
		Convey("When user agent contains a keyword", func() {
			Convey("It returns true", func() {
				for key, _ := range known {
					k := NewKnown(testBots[key], known)

					So(k.Match(), ShouldBeTrue)
				}
			})
		})

		Convey("When user agent does not contain a keyword", func() {
			Convey("It returns false", func() {
				u := "Mozilla/5.0 (Linux; Android 10.0.0; Nexus 7 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3359.158 Safari/537.36"
				So(NewKnown(u, known).Match(), ShouldBeFalse)
			})
		})
	})
}
