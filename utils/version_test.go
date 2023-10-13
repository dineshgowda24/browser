package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersionGTE(t *testing.T) {
	Convey("Subject: #VersionGTE", t, func() {
		Convey("When the actual version is greater than or equal to the expected version", func() {
			Convey("It should return true", func() {
				So(VersionGTE("1.0.0", "1.0.0"), ShouldBeTrue)
				So(VersionGTE("1.0.0", "0.9.0"), ShouldBeTrue)
				So(VersionGTE("1.0.0", "0.9.9"), ShouldBeTrue)
				So(VersionGTE("1.0.0", "0.0.0"), ShouldBeTrue)
				So(VersionGTE("1.0.0", ""), ShouldBeTrue)
				So(VersionGTE("1.0.0", "0"), ShouldBeTrue)
				So(VersionGTE("1.0.0", "0.0"), ShouldBeTrue)
				So(VersionGTE("1.0.0", "0.0.0"), ShouldBeTrue)
			})
		})

		Convey("When the actual version is ''", func() {
			Convey("It should return false", func() {
				So(VersionGTE("", "1.0.0"), ShouldBeFalse)
			})
		})

		Convey("When the actual version is unparsable", func() {
			Convey("It should return false", func() {
				So(VersionGTE("@#", "1.0.0"), ShouldBeFalse)
			})
		})

		Convey("When the expected version is unparsable", func() {
			Convey("It should return false", func() {
				So(VersionGTE("1.0.0", "@#"), ShouldBeFalse)
			})
		})

		Convey("When the actual version is less than the expected version", func() {
			Convey("It should return false", func() {
				So(VersionGTE("1.0.0", "1.0.1"), ShouldBeFalse)
				So(VersionGTE("1.0.0", "1.1.0"), ShouldBeFalse)
				So(VersionGTE("1.0.0", "2.0.0"), ShouldBeFalse)
				So(VersionGTE("1.0.0", "4"), ShouldBeFalse)
				So(VersionGTE("1.0.0", "4.0"), ShouldBeFalse)
			})
		})
	})
}

func TestVersionGT(t *testing.T) {
	Convey("Subject: #VersionGT", t, func() {
		Convey("When the actual version is greater than the expected version", func() {
			Convey("It should return true", func() {
				So(VersionGT("1.0.0", "0.9.0"), ShouldBeTrue)
				So(VersionGT("1.0.0", "0.9.9"), ShouldBeTrue)
				So(VersionGT("1.0.0", "0.0.0"), ShouldBeTrue)
				So(VersionGT("1.0.0", ""), ShouldBeTrue)
				So(VersionGT("1.0.0", "0"), ShouldBeTrue)
				So(VersionGT("1.0.0", "0.0"), ShouldBeTrue)
				So(VersionGT("1.0.0", "0.0.0"), ShouldBeTrue)
			})
		})

		Convey("When the actual version is ''", func() {
			Convey("It should return false", func() {
				So(VersionGT("", "1.0.0"), ShouldBeFalse)
			})
		})

		Convey("When the actual version is unparsable", func() {
			Convey("It should return false", func() {
				So(VersionGT("@#", "1.0.0"), ShouldBeFalse)
			})
		})

		Convey("When the expected version is unparsable", func() {
			Convey("It should return false", func() {
				So(VersionGT("1.0.0", "@#"), ShouldBeFalse)
			})
		})

		Convey("When the actual version is less than or equal to the expected version", func() {
			Convey("It should return false", func() {
				So(VersionGT("1.0.0", "1.0.0"), ShouldBeFalse)
				So(VersionGT("1.0.0", "1.0.1"), ShouldBeFalse)
				So(VersionGT("1.0.0", "1.1.0"), ShouldBeFalse)
				So(VersionGT("1.0.0", "2.0.0"), ShouldBeFalse)
				So(VersionGT("1.0.0", "4"), ShouldBeFalse)
				So(VersionGT("1.0.0", "4.0"), ShouldBeFalse)
			})
		})
	})
}

func TestVersionLTE(t *testing.T) {
	Convey("Subject: #VersionLTE", t, func() {
		Convey("When the actual version is less than or equal to the expected version", func() {
			Convey("It should return true", func() {
				So(VersionLTE("1.0.0", "1.0.0"), ShouldBeTrue)
				So(VersionLTE("1.0.0", "1.0.1"), ShouldBeTrue)
				So(VersionLTE("1.0.0", "1.1.0"), ShouldBeTrue)
				So(VersionLTE("1.0.0", "2.0.0"), ShouldBeTrue)
				So(VersionLTE("1.0.0", "4"), ShouldBeTrue)
				So(VersionLTE("1.0.0", "4.0"), ShouldBeTrue)
			})
		})

		Convey("When the actual version is ''", func() {
			Convey("It should return true", func() {
				So(VersionLTE("", "1.0.0"), ShouldBeTrue)
			})
		})

		Convey("When the actual version is unparsable", func() {
			Convey("It should return false", func() {
				So(VersionLTE("@#", "1.0.0"), ShouldBeFalse)
			})
		})

		Convey("When the expected version is unparsable", func() {
			Convey("It should return false", func() {
				So(VersionLTE("1.0.0", "@#"), ShouldBeFalse)
			})
		})

		Convey("When the actual version is greater than the expected version", func() {
			Convey("It should return false", func() {
				So(VersionLTE("1.0.0", "0.9.0"), ShouldBeFalse)
				So(VersionLTE("1.0.0", "0.9.9"), ShouldBeFalse)
				So(VersionLTE("1.0.0", "0.0.0"), ShouldBeFalse)
				So(VersionLTE("1.0.0", ""), ShouldBeFalse)
				So(VersionLTE("1.0.0", "0"), ShouldBeFalse)
				So(VersionLTE("1.0.0", "0.0"), ShouldBeFalse)
				So(VersionLTE("1.0.0", "0.0.0"), ShouldBeFalse)
			})
		})
	})
}

func TestVersionLT(t *testing.T) {
	Convey("Subject: #VersionLT", t, func() {
		Convey("When the actual version is less than the expected version", func() {
			Convey("It should return true", func() {
				So(VersionLT("1.0.0", "1.0.1"), ShouldBeTrue)
				So(VersionLT("1.0.0", "1.1.0"), ShouldBeTrue)
				So(VersionLT("1.0.0", "2.0.0"), ShouldBeTrue)
				So(VersionLT("1.0.0", "4"), ShouldBeTrue)
				So(VersionLT("1.0.0", "4.0"), ShouldBeTrue)
			})
		})

		Convey("When the actual version is ''", func() {
			Convey("It should return true", func() {
				So(VersionLT("", "1.0.0"), ShouldBeTrue)
			})
		})

		Convey("When the actual version is unparsable", func() {
			Convey("It should return false", func() {
				So(VersionLT("@#", "1.0.0"), ShouldBeFalse)
			})
		})

		Convey("When the expected version is unparsable", func() {
			Convey("It should return false", func() {
				So(VersionLT("1.0.0", "@#"), ShouldBeFalse)
			})
		})

		Convey("When the actual version is greater than or equal to the expected version", func() {
			Convey("It should return false", func() {
				So(VersionLT("1.0.0", "1.0.0"), ShouldBeFalse)
				So(VersionLT("1.0.0", "0.9.0"), ShouldBeFalse)
				So(VersionLT("1.0.0", "0.9.9"), ShouldBeFalse)
				So(VersionLT("1.0.0", "0.0.0"), ShouldBeFalse)
				So(VersionLT("1.0.0", ""), ShouldBeFalse)
				So(VersionLT("1.0.0", "0"), ShouldBeFalse)
				So(VersionLT("1.0.0", "0.0"), ShouldBeFalse)
				So(VersionLT("1.0.0", "0.0.0"), ShouldBeFalse)
			})
		})
	})
}
