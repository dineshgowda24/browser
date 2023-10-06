package browser

import (
	"regexp"

	"github.com/dineshgowda24/browser/devices"
)


// DeviceMatcher is an interface for device matchers
type DeviceMatcher interface {
	Matcher
}

// Device is a struct that contains information about the user agent's device
type Device struct {
	userAgent string
	matcher   DeviceMatcher
}

// NewDevice creates a new device based on the user agent string
func NewDevice(userAgent string) (*Device, error) {
	// if user agent is larger than the limit, return error
	if len(userAgent) > userAgentSizeLimit {
		return nil, ErrUserAgentSizeExceeded
	}

	return &Device{
		userAgent: userAgent,
	}, nil
}

// getMatcher returns the device matcher detected from the user agent string
func (d *Device) getMatcher() DeviceMatcher {
	if d.matcher != nil {
		return d.matcher
	}

	// define all your device matchers here
	matchers := []DeviceMatcher{
		devices.NewBlackberryPlaybook(d.userAgent),
		devices.NewIphone(d.userAgent),
		devices.NewIpad(d.userAgent),
		devices.NewIpodTouch(d.userAgent),
		devices.NewKindleFire(d.userAgent),
		devices.NewKindle(d.userAgent),
		devices.NewPlayStation3(d.userAgent),
		devices.NewPlayStation4(d.userAgent),
		devices.NewPlayStation5(d.userAgent),
		devices.NewPSVita(d.userAgent),
		devices.NewPSP(d.userAgent),
		devices.NewWiiU(d.userAgent),
		devices.NewWii(d.userAgent),
		devices.NewXboxOne(d.userAgent), // this must be before xbox 360
		devices.NewXbox360(d.userAgent),
		devices.NewSwitch(d.userAgent),
		devices.NewIpad(d.userAgent),
		devices.NewKindle(d.userAgent),
		devices.NewTV(d.userAgent), // this must be before android
		devices.NewAndroid(d.userAgent),
		devices.NewUnknown(d.userAgent),
	}

	for _, matcher := range matchers {
		if matcher.Match() {
			d.matcher = matcher
			break
		}
	}
	return d.matcher
}

// Name returns the name of the device
func (d *Device) Name() string {
	return d.getMatcher().Name()
}

// IsBlackberryPlaybook returns true if the device is a Blackberry Playbook
func (d *Device) IsBlackberryPlaybook() bool {
	if _, ok := d.getMatcher().(*devices.BlackberryPlaybook); ok {
		return true
	}

	return false
}

// IsIPad returns true if the device is an iPad
func (d *Device) IsIPad() bool {
	if _, ok := d.getMatcher().(*devices.Ipad); ok {
		return true
	}

	return false
}

// IsIPhone returns true if the device is an iPhone
func (d *Device) IsIPhone() bool {
	if _, ok := d.getMatcher().(*devices.Iphone); ok {
		return true
	}

	return false
}

// IsIPodTouch returns true if the device is an iPod touch
func (d *Device) IsIPodTouch() bool {
	if _, ok := d.getMatcher().(*devices.IpodTouch); ok {
		return true
	}

	return false
}

// IsConsole returns true if the device is a console
// This includes PlayStation, Xbox, and Nintendo devices
func (d *Device) IsConsole() bool {
	return d.IsPlayStation() || d.IsXbox() || d.IsNintendo()
}

// IsPlayStation returns true if the device is a PlayStation
// This includes PlayStation 3, PlayStation 4, and PlayStation 5
func (d *Device) IsPlayStation() bool {
	return d.IsPs3() || d.IsPs4() || d.IsPs5()
}

// IsPs3 returns true if the device is a PlayStation 3
func (d *Device) IsPs3() bool {
	if _, ok := d.getMatcher().(*devices.PlayStation3); ok {
		return true
	}

	return false
}

// IsPs4 returns true if the device is a PlayStation 4
func (d *Device) IsPs4() bool {
	if _, ok := d.getMatcher().(*devices.PlayStation4); ok {
		return true
	}

	return false
}

// IsPs5 returns true if the device is a PlayStation 5
func (d *Device) IsPs5() bool {
	if _, ok := d.getMatcher().(*devices.PlayStation5); ok {
		return true
	}

	return false
}

// IsXbox returns true if the device is an Xbox
// This includes Xbox 360 and Xbox One
func (d *Device) IsXbox() bool {
	return d.IsXbox360() || d.IsXboxOne()
}

// IsXbox360 returns true if the device is an Xbox 360
func (d *Device) IsXbox360() bool {
	if _, ok := d.getMatcher().(*devices.Xbox360); ok {
		return true
	}

	return false
}

// IsXboxOne returns true if the device is an Xbox One
func (d *Device) IsXboxOne() bool {
	if _, ok := d.getMatcher().(*devices.XboxOne); ok {
		return true
	}

	return false
}

// IsNintendo returns true if the device is a Nintendo
// This includes Wii, Wii U, and Switch
func (d *Device) IsNintendo() bool {
	return d.IsWii() || d.IsWiiU() || d.IsSwitch()
}

// IsWii returns true if the device is a Wii
func (d *Device) IsWii() bool {
	if _, ok := d.getMatcher().(*devices.Wii); ok {
		return true
	}

	return false
}

// IsNintendoWii returns true if the device is a Nintendo Wii
// IsNintendoWii is an alias for IsWii
func (d *Device) IsNintendoWii() bool {
	return d.IsWii()
}

// IsWiiU returns true if the device is a Wii U
func (d *Device) IsWiiU() bool {
	if _, ok := d.getMatcher().(*devices.WiiU); ok {
		return true
	}

	return false
}

// IsNintendoWiiU returns true if the device is a Nintendo Wii U
// IsNintendoWiiU is an alias for IsWiiU
func (d *Device) IsNintendoWiiU() bool {
	return d.IsWiiU()
}

// IsSwitch returns true if the device is a Switch
func (d *Device) IsSwitch() bool {
	if _, ok := d.getMatcher().(*devices.Switch); ok {
		return true
	}

	return false
}

// IsNintendoSwitch returns true if the device is a Nintendo Switch
// IsNintendoSwitch is an alias for IsSwitch
func (d *Device) IsNintendoSwitch() bool {
	return d.IsSwitch()
}

// IsKindle returns true if the device is a Kindle
// This includes Kindle Fire
func (d *Device) IsKindle() bool {
	isKindle := false
	if _, ok := d.getMatcher().(*devices.Kindle); ok {
		isKindle = true
	}

	return isKindle || d.IsKindleFire()
}

// IsKindleFire returns true if the device is a Kindle Fire
func (d *Device) IsKindleFire() bool {
	if _, ok := d.getMatcher().(*devices.KindleFire); ok {
		return true
	}

	return false
}

// IsPSP returns true if the device is a PlayStation Portable
func (d *Device) IsPSP() bool {
	if _, ok := d.getMatcher().(*devices.PSP); ok {
		return true
	}

	return false
}

// IsTV returns true if the device is a TV
func (d *Device) IsTV() bool {
	if _, ok := d.getMatcher().(*devices.TV); ok {
		return true
	}

	return false
}
