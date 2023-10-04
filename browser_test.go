package browser

import (
	"log"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"
)

// ua is a struct to hold the user agent strings for each platform from the matchers.yml file
// only used for testing
type ua struct {
	Android string `yaml:"android"`
	IOS     string `yaml:"ios"`
	Mac     string `yaml:"mac"`
	Windows string `yaml:"windows"`
	Linux   string `yaml:"linux"`
}

// testUserAgents is a map of user agent strings for each platform from the matchers.yml file
var testUserAgents map[string]ua

// initTestUserAgents reads the matchers.yml file and
// unmarshals it into the testUserAgents map.
func initTestUserAgents() {
	yamlFile, err := os.ReadFile("assets/_test/matchers.yml")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var data map[string]ua
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	testUserAgents = data
}

// TestMain is the entry point for running tests in this package.
func TestMain(m *testing.M) {
	initTestUserAgents()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("When the browser is Alipay", func() {
			Convey("It should return Alipay", func() {
				ua := testUserAgents["alipay"]
				b, _ := NewBrowser(ua.Android)
				So(b.Name(), ShouldEqual, "Alipay")
			})
		})

		Convey("When the browser is Nokia", func() {
			Convey("It should return Nokia", func() {
				ua := testUserAgents["nokia"]
				b, _ := NewBrowser(ua.Android)
				So(b.Name(), ShouldEqual, "Nokia S40 Ovi Browser")
			})
		})

		Convey("When the browser is UC Browser", func() {
			Convey("It should return UC Browser", func() {
				ua := testUserAgents["uc-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.Name(), ShouldEqual, "UCBrowser")
			})
		})

		Convey("When the browser is BlackBerry", func() {
			Convey("It should return BlackBerry", func() {
				ua := testUserAgents["blackberry"]
				b, _ := NewBrowser(ua.Windows)
				So(b.Name(), ShouldEqual, "BlackBerry")
			})
		})

		Convey("When the browser is Opera", func() {
			Convey("It should return Opera", func() {
				ua := testUserAgents["opera"]
				b, _ := NewBrowser(ua.Windows)
				So(b.Name(), ShouldEqual, "Opera")
			})
		})

		Convey("When the browser is Otter", func() {
			Convey("It should return Otter", func() {
				ua := testUserAgents["otter"]
				b, _ := NewBrowser(ua.Windows)
				So(b.Name(), ShouldEqual, "Otter")
			})
		})

		Convey("When the browser is Instagram", func() {
			Convey("It should return Instagram", func() {
				ua := testUserAgents["instagram"]
				b, _ := NewBrowser(ua.IOS)
				So(b.Name(), ShouldEqual, "Instagram")
			})
		})

		Convey("When the browser is Snapchat", func() {
			Convey("It should return Snapchat", func() {
				ua := testUserAgents["snapchat"]
				b, _ := NewBrowser(ua.IOS)
				So(b.Name(), ShouldEqual, "Snapchat")
			})
		})

		Convey("When the browser is Weibo", func() {
			Convey("It should return Weibo", func() {
				ua := testUserAgents["weibo"]
				b, _ := NewBrowser(ua.IOS)
				So(b.Name(), ShouldEqual, "Weibo")
			})
		})

		Convey("When the browser is Chrome", func() {
			Convey("It should return Chrome", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.Name(), ShouldEqual, "Chrome")
			})
		})

		Convey("When the browser is Safari", func() {
			Convey("It should return Safari", func() {
				ua := testUserAgents["safari"]
				b, _ := NewBrowser(ua.Mac)
				So(b.Name(), ShouldEqual, "Safari")
			})
		})
	})
}

func TestBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the browser is Alipay", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["alipay"]
				b, _ := NewBrowser(ua.Android)
				So(b.Version(), ShouldEqual, "1.7.0.123008")
			})
		})

		Convey("When the browser is Nokia", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["nokia"]
				b, _ := NewBrowser(ua.Android)
				So(b.Version(), ShouldEqual, "1.2.0.12")
			})
		})

		Convey("When the browser is UC Browser", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["uc-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.Version(), ShouldEqual, "10.6.2.599")
			})
		})

		Convey("When the browser is BlackBerry", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["blackberry"]
				b, _ := NewBrowser(ua.Windows)
				So(b.Version(), ShouldEqual, "7.1.0.346")
			})
		})

		Convey("When the browser is Opera", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["opera"]
				b, _ := NewBrowser(ua.Windows)
				So(b.Version(), ShouldEqual, "31.0.1889.174")
			})
		})

		Convey("When the browser is Otter", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["otter"]
				b, _ := NewBrowser(ua.Windows)
				So(b.Version(), ShouldEqual, "1.0.02")
			})
		})

		Convey("When the browser is Instagram", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["instagram"]
				b, _ := NewBrowser(ua.IOS)
				So(b.Version(), ShouldEqual, "10.18.0")
			})
		})
	})
}

func TestBrowserShortVersion(t *testing.T) {
	Convey("Subject: #ShortVersion", t, func() {
		Convey("When the browser is Alipay", func() {
			Convey("It should return the short version", func() {
				ua := testUserAgents["alipay"]
				b, _ := NewBrowser(ua.Android)
				So(b.ShortVersion(), ShouldEqual, "1")
			})
		})

		Convey("When the browser is Nokia", func() {
			Convey("It should return the short version", func() {
				ua := testUserAgents["nokia"]
				b, _ := NewBrowser(ua.Android)
				So(b.ShortVersion(), ShouldEqual, "1")
			})
		})

		Convey("When the browser is UC Browser", func() {
			Convey("It should return the short version", func() {
				ua := testUserAgents["uc-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.ShortVersion(), ShouldEqual, "10")
			})
		})

		Convey("When the browser is BlackBerry", func() {
			Convey("It should return the short version", func() {
				ua := testUserAgents["blackberry"]
				b, _ := NewBrowser(ua.Windows)
				So(b.ShortVersion(), ShouldEqual, "7")
			})
		})

		Convey("When the browser is Opera", func() {
			Convey("It should return the short version", func() {
				ua := testUserAgents["opera"]
				b, _ := NewBrowser(ua.Windows)
				So(b.ShortVersion(), ShouldEqual, "31")
			})
		})
	})
}
