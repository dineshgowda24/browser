package browser

import (
	"regexp"
	"strings"

	"github.com/dineshgowda24/browser/devices"
)

var (
	// regex copied from http://detectmobilebrowsers.com/mobile
	regMobileB = regexp.MustCompile(`(?i)(android|bb\d+|meego).+mobile|avantgo|bada\/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)\/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows ce|xda|xiino`)
	regMobileV = regexp.MustCompile(`(?i)1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s\-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|\-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw\-(n|u)|c55\/|capi|ccwa|cdm\-|cell|chtm|cldc|cmd\-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc\-s|devi|dica|dmob|do(c|p)o|ds(12|\-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(\-|_)|g1 u|g560|gene|gf\-5|g\-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd\-(m|p|t)|hei\-|hi(pt|ta)|hp( i|ip)|hs\-c|ht(c(\-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i\-(20|go|ma)|i230|iac( |\-|\/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |\/)|klon|kpt |kwc\-|kyo(c|k)|le(no|xi)|lg( g|\/(k|l|u)|50|54|\-[a-w])|libw|lynx|m1\-w|m3ga|m50\/|ma(te|ui|xo)|mc(01|21|ca)|m\-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(\-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)\-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|\-([1-8]|c))|phil|pire|pl(ay|uc)|pn\-2|po(ck|rt|se)|prox|psio|pt\-g|qa\-a|qc(07|12|21|32|60|\-[2-7]|i\-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55\/|sa(ge|ma|mm|ms|ny|va)|sc(01|h\-|oo|p\-)|sdk\/|se(c(\-|0|1)|47|mc|nd|ri)|sgh\-|shar|sie(\-|m)|sk\-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h\-|v\-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl\-|tdg\-|tel(i|m)|tim\-|t\-mo|to(pl|sh)|ts(70|m\-|m3|m5)|tx\-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|\-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(\-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas\-|your|zeto|zte\-`)
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
		devices.NewSurface(d.userAgent),
		devices.NewKindle(d.userAgent),
		devices.NewSamsung(d.userAgent),
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

// IsSurface returns true if the device is a Surface
func (d *Device) IsSurface() bool {
	if _, ok := d.getMatcher().(*devices.Surface); ok {
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

// IsSamsung returns true if the device is a Samsung device
func (d *Device) IsSamsung() bool {
	if _, ok := d.getMatcher().(*devices.Samsung); ok {
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

// IsMobile returns true if the device is a mobile device
func (d *Device) IsMobile() bool {
	return d.IsPSP() ||
		regexp.MustCompile(`zunewp7`).MatchString(d.userAgent) ||
		regMobileB.MatchString(d.userAgent) ||
		regMobileV.MatchString(d.userAgent[:4])
}

// IsTablet returns true if the device is a tablet
// This includes iPad, PlayBook, Surface, and Android based tablets
func (d *Device) IsTablet() bool {
	return d.IsIPad() ||
		d.IsBlackberryPlaybook() ||
		d.IsSurface() ||
		d.isAndroidTablet()
}

// isAndroidTablet returns true if the device is an Android based tablet
func (d *Device) isAndroidTablet() bool {
	a := regexp.MustCompile(`(?i)android`)

	return !d.IsMobile() &&
		a.MatchString(d.userAgent) &&
		!strings.Contains(d.userAgent, "KAIOS")
}
