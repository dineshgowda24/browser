package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewXboxOne(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		So(NewXboxOne(NewUAParser(testDevices["xbox-one-1"])), ShouldHaveSameTypeAs, &XboxOne{})
	})
}

func TestXboxOneName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Xbox One", func() {
			s := NewXboxOne(NewUAParser(""))
			So(s.Name(), ShouldEqual, "Xbox One")
		})
	})
}

func TestXboxOneMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewXboxOne(NewUAParser(testDevices["xbox-one-1"])).Match(), ShouldBeTrue)
				So(NewXboxOne(NewUAParser(testDevices["xbox-one-2"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When the user agent does not match", func() {
			Convey("It should return false", func() {
				So(NewXboxOne(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeFalse)
			})
		})
	})
}
