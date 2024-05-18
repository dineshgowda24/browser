package browser

import (
	"log"
	"os"
	"strings"
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
	yamlFile, err := os.ReadFile("assets/test/matchers.yml")
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
// It initializes all the test data.
func TestMain(m *testing.M) {
	initTestUserAgents()
	initTestDevices()
	initTestPlatforms()
	initTestBots()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestNewBrowser(t *testing.T) {
	Convey("Subject: #NewBrowser", t, func() {
		Convey("When the user agent is not empty", func() {
			Convey("It should return a browser", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b, ShouldNotBeNil)
			})
		})

		Convey("When the user agent is over 2kb", func() {
			Convey("It should return an error", func() {
				ua := strings.Repeat("a", 2049)
				_, err := NewBrowser(ua)
				So(err, ShouldEqual, ErrUserAgentSizeExceeded)
			})
		})
	})
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

func TestBrowserPlatform(t *testing.T) {
	Convey("Subject: #Platform", t, func() {
		Convey("It should return the platform", func() {
			ua := testUserAgents["chrome"]
			b, _ := NewBrowser(ua.Windows)
			So(b.Platform(), ShouldHaveSameTypeAs, &Platform{})
		})
	})
}

func TestBrowserDevice(t *testing.T) {
	Convey("Subject: #Device", t, func() {
		Convey("It should return the device", func() {
			ua := testUserAgents["chrome"]
			b, _ := NewBrowser(ua.Windows)
			So(b.Device(), ShouldHaveSameTypeAs, &Device{})
		})
	})
}

func TestBrowserBot(t *testing.T) {
	Convey("Subject: #Bot", t, func() {
		Convey("It should return the bot", func() {
			ua := testUserAgents["googlebot"]
			b, _ := NewBrowser(ua.Windows)
			So(b.Bot(), ShouldHaveSameTypeAs, &Bot{})
		})
	})
}

func TestBrowserIsAlipay(t *testing.T) {
	Convey("Subject: #IsAlipay", t, func() {
		Convey("When the browser is Alipay", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["alipay"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsAlipay(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Alipay", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsAlipay(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsNokia(t *testing.T) {
	Convey("Subject: #IsNokia", t, func() {
		Convey("When the browser is Nokia", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["nokia"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsNokia(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Nokia", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsNokia(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsUCBrowser(t *testing.T) {
	Convey("Subject: #IsUCBrowser", t, func() {
		Convey("When the browser is UC Browser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["uc-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsUCBrowser(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not UC Browser", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsUCBrowser(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsBlackBerry(t *testing.T) {
	Convey("Subject: #IsBlackBerry", t, func() {
		Convey("When the browser is BlackBerry", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["blackberry"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsBlackBerry(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not BlackBerry", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsBlackBerry(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsDuckDuckGo(t *testing.T) {
	Convey("Subject: #IsDuckDuckGo", t, func() {
		Convey("When the browser is DuckDuckGo", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["duck_duck_go"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsDuckDuckGo(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not DuckDuckGo", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsDuckDuckGo(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsGoogleSearchApp(t *testing.T) {
	Convey("Subject: #IsGoogleSearchApp", t, func() {
		Convey("When the browser is Google Search App", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["gsa"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsGoogleSearchApp(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Google Search App", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsGoogleSearchApp(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsHuaweiBrowser(t *testing.T) {
	Convey("Subject: #IsHuaweiBrowser", t, func() {
		Convey("When the browser is Huawei Browser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["huawei-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsHuaweiBrowser(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Huawei Browser", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsHuaweiBrowser(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsKonqueror(t *testing.T) {
	Convey("Subject: #IsKonqueror", t, func() {
		Convey("When the browser is Konqueror", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["konqueror"]
				b, _ := NewBrowser(ua.Linux)
				So(b.IsKonqueror(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Konqueror", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Linux)
				So(b.IsKonqueror(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsMaxthon(t *testing.T) {
	Convey("Subject: #IsMaxthon", t, func() {
		Convey("When the browser is Maxthon", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["maxthon"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsMaxthon(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Maxthon", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsMaxthon(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsMiuiBrowser(t *testing.T) {
	Convey("Subject: #IsMiuiBrowser", t, func() {
		Convey("When the browser is MiuiBrowser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["miui-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsMiuiBrowser(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not MiuiBrowser", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsMiuiBrowser(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsPaleMoon(t *testing.T) {
	Convey("Subject: #IsPaleMoon", t, func() {
		Convey("When the browser is Pale Moon", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["pale-moon"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsPaleMoon(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Pale Moon", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsPaleMoon(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsPuffin(t *testing.T) {
	Convey("Subject: #IsPuffin", t, func() {
		Convey("When the browser is Puffin", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["puffin"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsPuffin(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Puffin", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsPuffin(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsOpera(t *testing.T) {
	Convey("Subject: #IsOpera", t, func() {
		Convey("When the browser is Opera", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["opera"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsOpera(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Opera", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsOpera(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsOtter(t *testing.T) {
	Convey("Subject: #IsOtter", t, func() {
		Convey("When the browser is Otter", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["otter"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsOtter(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Otter", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsOtter(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsInstagram(t *testing.T) {
	Convey("Subject: #IsInstagram", t, func() {
		Convey("When the browser is Instagram", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["instagram"]
				b, _ := NewBrowser(ua.IOS)
				So(b.IsInstagram(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Instagram", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.IOS)
				So(b.IsInstagram(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsSnapchat(t *testing.T) {
	Convey("Subject: #IsSnapchat", t, func() {
		Convey("When the browser is Snapchat", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["snapchat"]
				b, _ := NewBrowser(ua.IOS)
				So(b.IsSnapchat(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Snapchat", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.IOS)
				So(b.IsSnapchat(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsWeibo(t *testing.T) {
	Convey("Subject: #IsWeibo", t, func() {
		Convey("When the browser is Weibo", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["weibo"]
				b, _ := NewBrowser(ua.IOS)
				So(b.IsWeibo(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Weibo", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.IOS)
				So(b.IsWeibo(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsMicroMessenger(t *testing.T) {
	Convey("Subject: #IsMicroMessenger", t, func() {
		Convey("When the browser is MicroMessenger", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["micro-messenger"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsMicroMessenger(), ShouldBeTrue)
				So(b.IsWechat(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not MicroMessenger", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsMicroMessenger(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsQQ(t *testing.T) {
	Convey("Subject: #IsQQ", t, func() {
		Convey("When the browser is QQ", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["qq"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsQQ(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not QQ", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsQQ(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsElectron(t *testing.T) {
	Convey("Subject: #IsElectron", t, func() {
		Convey("When the browser is Electron", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["electron"]
				b, _ := NewBrowser(ua.Mac)
				So(b.IsElectron(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Electron", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Mac)
				So(b.IsElectron(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsUnknown(t *testing.T) {
	Convey("Subject: #IsUnknown", t, func() {
		Convey("When the browser is Unknown", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["unknown"]
				b, _ := NewBrowser(ua.Mac)
				So(b.IsUnknown(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Unknown", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsUnknown(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsBrowserKnown(t *testing.T) {
	Convey("Subject: #IsBrowserKnown", t, func() {
		Convey("When the browser is known", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsBrowserKnown(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not known", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["unknown"]
				b, _ := NewBrowser(ua.Mac)
				So(b.IsBrowserKnown(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsInternetExplorer(t *testing.T) {
	Convey("Subject: #IsInternetExplorer", t, func() {
		Convey("When the browser is Internet Explorer", func() {
			Convey("It should return true", func() {
				ie := testUserAgents["ie"]
				oldIE := testUserAgents["old-ie"]

				uas := []string{ie.Windows, oldIE.Mac, oldIE.Windows}

				for _, ua := range uas {
					b, _ := NewBrowser(ua)
					So(b.IsInternetExplorer(), ShouldBeTrue)
				}
			})
		})

		Convey("When the browser is not Internet Explorer", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsInternetExplorer(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsSamsungBrowser(t *testing.T) {
	Convey("Subject: #IsSamsungBrowser", t, func() {
		Convey("When the browser is Samsung Browser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["samsung-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsSamsungBrowser(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Samsung Browser", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsSamsungBrowser(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsSougouBrowser(t *testing.T) {
	Convey("Subject: #IsSougouBrowser", t, func() {
		Convey("When the browser is Sougou Browser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["sogou-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsSougouBrowser(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Sougou Browser", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsSougouBrowser(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsVivoBrowser(t *testing.T) {
	Convey("Subject: #IsVivoBrowser", t, func() {
		Convey("When the browser is Vivo Browser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["vivo-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsVivoBrowser(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Vivo Browser", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsVivoBrowser(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsSputnik(t *testing.T) {
	Convey("Subject: #IsSputnik", t, func() {
		Convey("When the browser is Sputnik", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["sputnik"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsSputnik(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Sputnik", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsSputnik(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsYaaniBrowser(t *testing.T) {
	Convey("Subject: #IsYaaniBrowser", t, func() {
		Convey("When the browser is YaaniBrowser", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["yaani-browser"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsYaaniBrowser(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not YaaniBrowser", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsYaaniBrowser(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsYandex(t *testing.T) {
	Convey("Subject: #IsYandex", t, func() {
		Convey("When the browser is Yandex", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["yandex"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsYandex(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Yandex", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsYandex(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsChrome(t *testing.T) {
	Convey("Subject: #IsChrome", t, func() {
		Convey("When the browser is Chrome", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsChrome(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Chrome", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["alipay"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsChrome(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsSafari(t *testing.T) {
	Convey("Subject: #IsSafari", t, func() {
		Convey("When the browser is Safari", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["safari"]
				b, _ := NewBrowser(ua.Mac)
				So(b.IsSafari(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Safari", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["alipay"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsSafari(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsSafariWebappMode(t *testing.T) {
	Convey("Subject: #IsSafariWebappMode", t, func() {
		Convey("When the browser is Safari Webapp Mode", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.IOS)
				So(b.IsSafariWebappMode(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Safari Webapp Mode", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["alipay"]
				b, _ := NewBrowser(ua.Android)
				So(b.IsSafariWebappMode(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsEdge(t *testing.T) {
	Convey("Subject: #IsEdge", t, func() {
		Convey("When the browser is Edge", func() {
			uas := []string{testUserAgents["edge"].Windows, testUserAgents["edge"].Android}

			Convey("It should return true", func() {
				for _, ua := range uas {
					b, _ := NewBrowser(ua)
					So(b.IsEdge(), ShouldBeTrue)
				}
			})
		})

		Convey("When the browser is not Edge", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsEdge(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsFirefox(t *testing.T) {
	Convey("Subject: #IsFirefox", t, func() {
		Convey("When the browser is Firefox", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["firefox"]
				b, _ := NewBrowser(ua.Mac)
				So(b.IsFirefox(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Firefox", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsFirefox(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsBrave(t *testing.T) {
	Convey("Subject: #IsBrave", t, func() {
		Convey("When the browser is Brave", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["brave"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsBrave(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Brave", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsBrave(), ShouldBeFalse)
			})
		})
	})
}

func TestBrowserIsVivaldi(t *testing.T) {
	Convey("Subject: #IsVivaldi", t, func() {
		Convey("When the browser is Vivaldi", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["vivaldi"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsVivaldi(), ShouldBeTrue)
			})
		})

		Convey("When the browser is not Vivaldi", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]
				b, _ := NewBrowser(ua.Windows)
				So(b.IsVivaldi(), ShouldBeFalse)
			})
		})
	})
}
