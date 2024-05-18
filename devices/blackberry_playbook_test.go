package devices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBlackberryPlaybook(t *testing.T) {
	Convey("Subject: #NewBlackberryPlaybook", t, func() {
		Convey("It should return a new BlackberryPlaybook instance", func() {
			So(NewBlackberryPlaybook(NewUAParser("")), ShouldHaveSameTypeAs, &BlackberryPlaybook{})
		})
	})
}

func TestBlackberryPlaybookName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return BlackBerry PlayBook", func() {
			s := NewBlackberryPlaybook(NewUAParser(""))
			So(s.Name(), ShouldEqual, "BlackBerry Playbook")
		})
	})
}

func TestBlackberryPlaybookMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When the user agent matches", func() {
			Convey("It should return true", func() {
				So(NewBlackberryPlaybook(NewUAParser(testDevices["bb-playbook-1"])).Match(), ShouldBeTrue)
				So(NewBlackberryPlaybook(NewUAParser(testDevices["bb-playbook-2"])).Match(), ShouldBeTrue)
			})
		})
	})
}
