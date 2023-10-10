package browser

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"
)

// testPlatforms is a map of user agent strings for each platform from the platforms.yml file
var testPlatforms map[string]string

// initTestPlatforms reads the platforms.yml file and
// unmarshals it into the testPlatforms map.
func initTestPlatforms() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	wd = strings.Split(wd, "/browser")[0]
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/browser/assets/test/platforms.yml", wd))
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var data map[string]string
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	testPlatforms = data
}

func TestNewPlatform(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When NewPlatform is called", func() {
			platform, _ := NewPlatform(testPlatforms["adobe-air"])
			Convey("It returns a platform", func() {
				So(platform, ShouldHaveSameTypeAs, &Platform{})
			})
		})

		Convey("When NewPlatform is called with a user agent string larger than the limit", func() {
			Convey("It returns ErrUserAgentSizeExceeded error", func() {
				// generate a string that is larger than the limit of 2k
				largeString := strings.Repeat("a", 2049)
				_, err := NewPlatform(largeString)
				So(err, ShouldEqual, ErrUserAgentSizeExceeded)
			})
		})
	})
}

func TestPlatformName(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When NewPlatform is called", func() {
			p, _ := NewPlatform(testPlatforms["adobe-air"])
			Convey("It returns the correct platform name", func() {
				So(p.Name(), ShouldEqual, "Adobe AIR")
			})
		})
	})
}

func TestPlatformIsAdobeAir(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Adobe AIR", func() {
			p, _ := NewPlatform(testPlatforms["adobe-air"])
			Convey("It returns true", func() {
				So(p.IsAdobeAir(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Adobe AIR", func() {
			p, _ := NewPlatform(testPlatforms["blackberry"])
			Convey("It returns false", func() {
				So(p.IsAdobeAir(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsAdobeAirVersionCompatible(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Adobe AIR", func() {
			p, _ := NewPlatform(testPlatforms["adobe-air"])
			Convey("And the version is compatible", func() {
				Convey("It returns true", func() {
					So(p.IsAdobeAirVersionCompatible("17"), ShouldBeTrue)
				})
			})

			Convey("And the version is not compatible", func() {
				Convey("It returns false", func() {
					So(p.IsAdobeAirVersionCompatible("19"), ShouldBeFalse)
				})
			})
		})
	})
}

func TestPlatformIsAndroid(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Android", func() {
			p, _ := NewPlatform(testPlatforms["android-jelly-bean-4.1"])
			Convey("It returns true", func() {
				So(p.IsAndroid(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Android", func() {
			p, _ := NewPlatform(testPlatforms["blackberry"])
			Convey("It returns false", func() {
				So(p.IsAndroid(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsAndroidApp(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Android App", func() {
			Convey("It returns true", func() {
				apps := []string{"android-app-1", "android-app-2"}
				for _, app := range apps {
					p, _ := NewPlatform(testPlatforms[app])

					So(p.IsAndroidApp(), ShouldBeTrue)
					So(p.IsAndroidWebview(), ShouldBeTrue)
				}
			})
		})

		Convey("When the platform is not Android App", func() {
			p, _ := NewPlatform(testPlatforms["android-jelly-bean-4.1"])
			Convey("It returns false", func() {
				So(p.IsAndroidApp(), ShouldBeFalse)
				So(p.IsAndroidWebview(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsChromeOS(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Chrome OS", func() {
			p, _ := NewPlatform(testPlatforms["chrome-os-armv7i"])
			Convey("It returns true", func() {
				So(p.IsChromeOS(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Chrome OS", func() {
			p, _ := NewPlatform(testPlatforms["blackberry"])
			Convey("It returns false", func() {
				So(p.IsChromeOS(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsChromeOSVersionCompatible(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Chrome OS", func() {
			p, _ := NewPlatform(testPlatforms["chrome-os-armv7i"])
			Convey("And the version is compatible", func() {
				Convey("It returns true", func() {
					So(p.IsChromeOSVersionCompatible("6157"), ShouldBeTrue)
				})
			})

			Convey("And the version is not compatible", func() {
				Convey("It returns false", func() {
					So(p.IsChromeOSVersionCompatible("6158.50.0"), ShouldBeFalse)
				})
			})
		})
	})
}

func TestPlatformIsIOS(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is iOS", func() {
			p, _ := NewPlatform(testPlatforms["ios-12"])
			Convey("It returns true", func() {
				So(p.IsIOS(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not iOS", func() {
			p, _ := NewPlatform(testPlatforms["blackberry"])
			Convey("It returns false", func() {
				So(p.IsIOS(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsWatchOS(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is watchOS", func() {
			p, _ := NewPlatform(testPlatforms["apple-watch-6"])
			Convey("It returns true", func() {
				So(p.IsWatchOS(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not watchOS", func() {
			p, _ := NewPlatform(testPlatforms["blackberry"])
			Convey("It returns false", func() {
				So(p.IsWatchOS(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsKaiOS(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is KaiOS", func() {
			p, _ := NewPlatform(testPlatforms["kai-os-2"])
			Convey("It returns true", func() {
				So(p.IsKaiOS(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not KaiOS", func() {
			p, _ := NewPlatform(testPlatforms["blackberry"])
			Convey("It returns false", func() {
				So(p.IsKaiOS(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsLinux(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Linux", func() {
			p, _ := NewPlatform(testPlatforms["linux"])
			Convey("It returns true", func() {
				So(p.IsLinux(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Linux", func() {
			p, _ := NewPlatform(testPlatforms["blackberry"])
			Convey("It returns false", func() {
				So(p.IsLinux(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsBlackberry(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is BlackBerry", func() {
			p, _ := NewPlatform(testPlatforms["blackberry"])
			Convey("It returns true", func() {
				So(p.IsBlackBerry(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not BlackBerry", func() {
			p, _ := NewPlatform(testPlatforms["linux"])
			Convey("It returns false", func() {
				So(p.IsBlackBerry(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsWindowsMobile(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows Mobile", func() {
			p, _ := NewPlatform(testPlatforms["windows-ce"])
			Convey("It returns true", func() {
				So(p.IsWindowsMobile(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows Mobile", func() {
			p, _ := NewPlatform(testPlatforms["linux"])
			Convey("It returns false", func() {
				So(p.IsWindowsMobile(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsWindowsPhone(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows Phone", func() {
			p, _ := NewPlatform(testPlatforms["windows-phone"])
			Convey("It returns true", func() {
				So(p.IsWindowsPhone(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows Phone", func() {
			p, _ := NewPlatform(testPlatforms["linux"])
			Convey("It returns false", func() {
				So(p.IsWindowsPhone(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsIOSApp(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is iOS App", func() {
			p, _ := NewPlatform(testPlatforms["ios-app"])
			Convey("It returns true", func() {
				So(p.IsIOSApp(), ShouldBeTrue)
				So(p.IsIOSWebview(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not iOS App", func() {
			p, _ := NewPlatform(testPlatforms["ios-3"])
			Convey("It returns false", func() {
				So(p.IsIOSApp(), ShouldBeFalse)
				So(p.IsIOSWebview(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsWindows(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows", func() {
			p, _ := NewPlatform(testPlatforms["windows-10"])
			Convey("It returns true", func() {
				So(p.IsWindows(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows", func() {
			p, _ := NewPlatform(testPlatforms["linux"])
			Convey("It returns false", func() {
				So(p.IsWindows(), ShouldBeFalse)
			})
		})
	})
}

func TestPlatformIsWindowsXP(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows XP", func() {
			p, _ := NewPlatform(testPlatforms["windows-xp"])
			Convey("It returns true", func() {
				So(p.IsWindowsXP(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows XP", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-8-1", "windows-8", "windows-7", "windows-vista"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindowsXP(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindowsVista(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows Vista", func() {
			p, _ := NewPlatform(testPlatforms["windows-vista"])
			Convey("It returns true", func() {
				So(p.IsWindowsVista(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows Vista", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-8-1", "windows-8", "windows-7", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindowsVista(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindows7(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows 7", func() {
			p, _ := NewPlatform(testPlatforms["windows-7"])
			Convey("It returns true", func() {
				So(p.IsWindows7(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows 7", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-8-1", "windows-8", "windows-vista", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindows7(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindows8(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows 8", func() {
			p, _ := NewPlatform(testPlatforms["windows-8"])
			Convey("It returns true", func() {
				So(p.IsWindows8(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows 8", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-7", "windows-vista", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindows8(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindows8_1(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows 8.1", func() {
			p, _ := NewPlatform(testPlatforms["windows-8-1"])
			Convey("It returns true", func() {
				So(p.IsWindows8_1(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows 8.1", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-7", "windows-vista", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindows8_1(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindows10(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows 10", func() {
			p, _ := NewPlatform(testPlatforms["windows-10"])
			Convey("It returns true", func() {
				So(p.IsWindows10(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows 10", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-8-1", "windows-8", "windows-7", "windows-vista", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindows10(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindowsRT(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows RT", func() {
			p, _ := NewPlatform(testPlatforms["windows-rt"])
			Convey("It returns true", func() {
				So(p.IsWindowsRT(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows RT", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-8-1", "windows-8", "windows-7", "windows-vista", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindowsRT(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindowsX64(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows x64", func() {
			p, _ := NewPlatform(testPlatforms["windows-8-1"])
			Convey("It returns true", func() {
				So(p.IsWindowsX64(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows x64", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-8", "windows-7", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindowsX64(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindowsWOW64(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows WOW64", func() {
			p, _ := NewPlatform(testPlatforms["windows-10"])
			Convey("It returns true", func() {
				So(p.IsWindowsWOW64(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows WOW64", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-vista", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindowsWOW64(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindowsX64Inclusive(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows X64 Inclusive", func() {
			p, _ := NewPlatform(testPlatforms["windows-inclusive"])
			Convey("It returns true", func() {
				So(p.IsWindowsX64Inclusive(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows X64 Inclusive", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-8", "windows-7", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindowsX64(), ShouldBeFalse)
				}
			})
		})
	})
}

func TestPlatformIsWindowsTouchScreenDesktop(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the platform is Windows Touch Screen Desktop", func() {
			p, _ := NewPlatform(testPlatforms["windows-touch"])
			Convey("It returns true", func() {
				So(p.IsWindowsTouchScreenDesktop(), ShouldBeTrue)
			})
		})

		Convey("When the platform is not Windows Touch Screen Desktop", func() {
			Convey("It returns false", func() {
				platforms := []string{"windows-10", "windows-8", "windows-7", "windows-xp"}
				for _, platform := range platforms {
					p, _ := NewPlatform(testPlatforms[platform])
					So(p.IsWindowsTouchScreenDesktop(), ShouldBeFalse)
				}
			})
		})
	})
}
