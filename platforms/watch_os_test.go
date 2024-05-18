package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWatchOS(t *testing.T) {
	Convey("Subject: #NewWatchOS", t, func() {
		Convey("It should return a new WatchOS instance", func() {
			So(NewWatchOS(NewUAParser("")), ShouldHaveSameTypeAs, &WatchOS{})
		})
	})
}

func TestWatchOSName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			So(NewWatchOS(NewUAParser(testPlatforms["watch-3"])).Name(), ShouldEqual, "Apple Watch OS")
		})
	})
}

func TestWatchOSVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the correct version", func() {
				So(NewWatchOS(NewUAParser(testPlatforms["apple-watch-3"])).Version(), ShouldEqual, "5.2.1")
				So(NewWatchOS(NewUAParser(testPlatforms["apple-watch-4"])).Version(), ShouldEqual, "5.3")
				So(NewWatchOS(NewUAParser(testPlatforms["apple-watch-5"])).Version(), ShouldEqual, "6.0.1")
				So(NewWatchOS(NewUAParser(testPlatforms["apple-watch-6"])).Version(), ShouldEqual, "7.0")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return 0.0", func() {
				So(NewWatchOS(NewUAParser(testPlatforms["mac-os"])).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestWatchOSMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches WatchOS", func() {
			Convey("It should return true", func() {
				So(NewWatchOS(NewUAParser(testPlatforms["apple-watch-3"])).Match(), ShouldBeTrue)
				So(NewWatchOS(NewUAParser(testPlatforms["apple-watch-4"])).Match(), ShouldBeTrue)
				So(NewWatchOS(NewUAParser(testPlatforms["apple-watch-5"])).Match(), ShouldBeTrue)
				So(NewWatchOS(NewUAParser(testPlatforms["apple-watch-6"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match WatchOS", func() {
			Convey("It should return false", func() {
				So(NewWatchOS(NewUAParser(testPlatforms["mac-os"])).Match(), ShouldBeFalse)
			})
		})
	})
}
