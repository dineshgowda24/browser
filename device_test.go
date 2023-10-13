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

// testDevices is a map of user agent strings for each device from the devices.yml file
var testDevices map[string]string

// initTestDevices reads the devices.yml file and
// unmarshals it into the testDevices map.
func initTestDevices() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	wd = strings.Split(wd, "/browser")[0]
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/browser/assets/test/devices.yml", wd))
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var data map[string]string
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	testDevices = data
}

func TestNewDevice(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When NewDevice is called", func() {
			Convey("It returns a device", func() {
				d, _ := NewDevice(testDevices["iphone-7"])
				So(d, ShouldHaveSameTypeAs, &Device{})
			})
		})

		Convey("When NewDevice is called with a user agent string larger than the limit", func() {
			Convey("It returns ErrUserAgentSizeExceeded error", func() {
				// generate a string that is larger than the limit of 2k
				largeString := strings.Repeat("a", 2049)
				_, err := NewDevice(largeString)
				So(err, ShouldEqual, ErrUserAgentSizeExceeded)
			})
		})
	})
}

func TestDeviceName(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When NewDevice is called", func() {
			Convey("Then it should return the correct device name", func() {
				d, _ := NewDevice(testDevices["iphone-7"])
				So(d.Name(), ShouldEqual, "iPhone")
			})
		})
	})
}

func TestDeviceIsBlackberryPlaybook(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a Blackberry Playbook", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["bb-playbook-1"])
				So(d.IsBlackberryPlaybook(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsIPad(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is an iPad", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["ipad-4"])
				So(d.IsIPad(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsIPhone(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is an iPhone", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["iphone-7"])
				So(d.IsIPhone(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsIPodTouch(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is an iPod touch", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["ipod-touch-1"])
				So(d.IsIPodTouch(), ShouldBeTrue)
			})
		})

		Convey("When the device is not an iPod touch", func() {
			Convey("It should return false", func() {
				d, _ := NewDevice(testDevices["iphone-7"])
				So(d.IsIPodTouch(), ShouldBeFalse)
			})
		})
	})
}

func TestDeviceIsPlayStation3(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a PlayStation 3", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["ps3"])
				So(d.IsPs3(), ShouldBeTrue)
				So(d.IsPlayStation(), ShouldBeTrue)
				So(d.IsConsole(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsPlayStation4(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a PlayStation 4", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["ps4"])
				So(d.IsPs4(), ShouldBeTrue)
				So(d.IsPlayStation(), ShouldBeTrue)
				So(d.IsConsole(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsPlayStation5(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a PlayStation 5", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["ps5"])
				So(d.IsPs5(), ShouldBeTrue)
				So(d.IsPlayStation(), ShouldBeTrue)
				So(d.IsConsole(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsXbox360(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is an Xbox 360", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["xbox360-1"])
				So(d.IsXbox360(), ShouldBeTrue)
				So(d.IsXbox(), ShouldBeTrue)
				So(d.IsConsole(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsXboxOne(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is an Xbox One", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["xbox-one-1"])
				So(d.IsXboxOne(), ShouldBeTrue)
				So(d.IsXbox(), ShouldBeTrue)
				So(d.IsConsole(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsWii(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a Wii", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["wii-1"])
				So(d.IsWii(), ShouldBeTrue)
				So(d.IsNintendoWii(), ShouldBeTrue)
				So(d.IsNintendo(), ShouldBeTrue)
				So(d.IsConsole(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsWiiU(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a Wii U", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["wiiu-1"])
				So(d.IsWiiU(), ShouldBeTrue)
				So(d.IsNintendoWiiU(), ShouldBeTrue)
				So(d.IsNintendo(), ShouldBeTrue)
				So(d.IsConsole(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsSwitch(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a Switch", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["switch-1"])
				So(d.IsSwitch(), ShouldBeTrue)
				So(d.IsNintendoSwitch(), ShouldBeTrue)
				So(d.IsNintendo(), ShouldBeTrue)
				So(d.IsConsole(), ShouldBeTrue)
			})
		})

		Convey("When the device is not a Switch", func() {
			Convey("It should return false", func() {
				d, _ := NewDevice(testDevices["wiiu-1"])
				So(d.IsSwitch(), ShouldBeFalse)
			})
		})
	})
}

func TestDeviceIsSurface(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a Surface", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["surface-1"])
				So(d.IsSurface(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsKindle(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a Kindle or Kindle fire", func() {
			Convey("It should return true", func() {
				kindleUserAgents := []string{testDevices["kindle-1"], testDevices["kindle-fire-1"]}
				for _, ua := range kindleUserAgents {
					d, _ := NewDevice(ua)
					So(d.IsKindle(), ShouldBeTrue)
				}
			})
		})
	})
}

func TestDeviceIsKindleFire(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a Kindle Fire", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["kindle-fire-1"])
				So(d.IsKindleFire(), ShouldBeTrue)
			})
		})

		Convey("When the device is not a Kindle Fire", func() {
			Convey("It should return false", func() {
				d, _ := NewDevice(testDevices["kindle-1"])
				So(d.IsKindleFire(), ShouldBeFalse)
			})
		})
	})
}

func TestDeviceIsPSP(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a PSP", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["psp-1"])
				So(d.IsPSP(), ShouldBeTrue)
			})
		})
	})
}

func TestDeviceIsSamsung(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a Samsung", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["samsung-1"])
				So(d.IsSamsung(), ShouldBeTrue)
			})
		})

		Convey("When the device is not a Samsung", func() {
			Convey("It should return false", func() {
				d, _ := NewDevice(testDevices["iphone-7"])
				So(d.IsSamsung(), ShouldBeFalse)
			})
		})
	})
}

func TestDeviceIsTV(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a TV", func() {
			Convey("It should return true", func() {
				d, _ := NewDevice(testDevices["tv-1"])
				So(d.IsTV(), ShouldBeTrue)
			})
		})

		Convey("When the device is not a TV", func() {
			Convey("It should return false", func() {
				d, _ := NewDevice(testDevices["iphone-7"])
				So(d.IsTV(), ShouldBeFalse)
			})
		})
	})
}

func TestDeviceIsMobile(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a mobile", func() {
			Convey("It should return true", func() {
				mobiles := []string{testDevices["iphone-7"], testDevices["android-froyo-2-2"], testDevices["mobile-1"], testDevices["mobile-2"]}
				for _, ua := range mobiles {
					d, _ := NewDevice(ua)
					So(d.IsMobile(), ShouldBeTrue)
				}
			})
		})

		Convey("When the device is not a mobile", func() {
			Convey("It should return false", func() {
				d, _ := NewDevice(testDevices["ipad-1"])
				So(d.IsMobile(), ShouldBeFalse)
			})
		})
	})
}

func TestDeviceIsTablet(t *testing.T) {
	Convey("Given a user agent string", t, func() {
		Convey("When the device is a tablet", func() {
			Convey("It should return true", func() {
				tablets := []string{testDevices["ipad-4"], testDevices["surface-1"], testDevices["bb-playbook-1"]}
				for _, ua := range tablets {
					d, _ := NewDevice(ua)
					So(d.IsTablet(), ShouldBeTrue)
				}
			})
		})

		Convey("When the device is not a tablet", func() {
			Convey("It should return false", func() {
				d, _ := NewDevice(testDevices["iphone-7"])
				So(d.IsTablet(), ShouldBeFalse)
			})
		})
	})
}
