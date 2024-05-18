package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewOpera(t *testing.T) {
	Convey("Subject: #NewOpera", t, func() {
		Convey("It should return a new Opera instance", func() {
			So(NewOpera(NewUAParser("")), ShouldHaveSameTypeAs, &Opera{})
		})
	})
}

func TestOperaName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Opera", func() {
			o := NewOpera(NewUAParser(""))
			So(o.Name(), ShouldEqual, "Opera")
		})
	})
}

func TestOperaVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				// Opera browsers
				op := testUserAgents["opera"]
				So(NewOpera(NewUAParser(op.Linux)).Version(), ShouldEqual, "31.0.1889.174")
				So(NewOpera(NewUAParser(op.Mac)).Version(), ShouldEqual, "31.0.1889.174")
				So(NewOpera(NewUAParser(op.Windows)).Version(), ShouldEqual, "31.0.1889.174")
				So(NewOpera(NewUAParser(op.IOS)).Version(), ShouldEqual, "12.0")
				So(NewOpera(NewUAParser(op.Android)).Version(), ShouldEqual, "28.0.1764.90386")

				// Opera Mini browsers
				opm := testUserAgents["opera-mini"]
				So(NewOpera(NewUAParser(opm.Linux)).Version(), ShouldEqual, "7.6.11")
				So(NewOpera(NewUAParser(opm.Mac)).Version(), ShouldEqual, "10.2.0.93022")
				So(NewOpera(NewUAParser(opm.Windows)).Version(), ShouldEqual, "8.1.0")
				So(NewOpera(NewUAParser(opm.IOS)).Version(), ShouldEqual, "10.2.0")
				So(NewOpera(NewUAParser(opm.Android)).Version(), ShouldEqual, "10.0.1884")

				// Opera Touch browsers
				opt := testUserAgents["opera-touch"]
				So(NewOpera(NewUAParser(opt.Linux)).Version(), ShouldEqual, "1.6.18")
				So(NewOpera(NewUAParser(opt.Mac)).Version(), ShouldEqual, "69")
				So(NewOpera(NewUAParser(opt.Windows)).Version(), ShouldEqual, "1.6.18")
				So(NewOpera(NewUAParser(opt.IOS)).Version(), ShouldEqual, "1.0.1.58")
				So(NewOpera(NewUAParser(opt.Android)).Version(), ShouldEqual, "1.0.9")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				o := NewOpera(NewUAParser(testUserAgents["chrome"].Linux))
				So(o.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestOperaMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Opera", func() {
			Convey("It should return true", func() {
				// Opera browsers
				op := testUserAgents["opera"]
				So(NewOpera(NewUAParser(op.Linux)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(op.Mac)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(op.Windows)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(op.IOS)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(op.Android)).Match(), ShouldBeTrue)

				// Opera Mini browsers
				opm := testUserAgents["opera-mini"]
				So(NewOpera(NewUAParser(opm.Linux)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(opm.Mac)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(opm.Windows)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(opm.IOS)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(opm.Android)).Match(), ShouldBeTrue)

				// Opera Touch browsers
				opt := testUserAgents["opera-touch"]
				So(NewOpera(NewUAParser(opt.Linux)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(opt.Mac)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(opt.Windows)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(opt.IOS)).Match(), ShouldBeTrue)
				So(NewOpera(NewUAParser(opt.Android)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Opera", func() {
			Convey("It should return false", func() {
				o := NewOpera(NewUAParser(testUserAgents["chrome"].Linux))
				So(o.Match(), ShouldBeFalse)
			})
		})
	})
}
