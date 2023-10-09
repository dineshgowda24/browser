package browser

import (
	"regexp"
	"strings"

	"github.com/dineshgowda24/browser/platforms"
	"github.com/dineshgowda24/browser/utils"
)

// PlatformMatcher is an interface for user agent platform matchers.
// A platform matcher is used to detect the platform from the user agent string.
type PlatformMatcher interface {
	Matcher
	Version() string // Version returns the version of the platforms.
}

// Platform is a struct that contains information about the user agent's platforms.
type Platform struct {
	userAgent string          // the user agent string
	matcher   PlatformMatcher // detected platform matcher
}

// NewPlatform creates a new platform based on the user agent string.
//
// UserAgentSizeError is returned if the user agent string is larger than the limit.
func NewPlatform(userAgent string) (*Platform, error) {
	// if user agent is larger than the limit, return error
	if len(userAgent) > userAgentSizeLimit {
		return nil, ErrUserAgentSizeExceeded
	}

	return &Platform{
		userAgent: userAgent,
	}, nil
}

// getMatcher returns the platform matcher detected from the user agent string.
// It sets the matcher if it is not already set.
// If the matcher is already set, it returns the matcher.
func (p *Platform) getMatcher() PlatformMatcher {
	if p.matcher != nil {
		return p.matcher
	}

	// define all your platform matchers here
	matchers := []PlatformMatcher{
		platforms.NewAdobeAir(p.userAgent),
		platforms.NewBlackBerry(p.userAgent),
		platforms.NewKaiOS(p.userAgent),
		platforms.NewIOS(p.userAgent),
		platforms.NewWatchOS(p.userAgent),
		platforms.NewWindowsMobile(p.userAgent),
		platforms.NewWindowsPhone(p.userAgent),
		platforms.NewWindows(p.userAgent),
		platforms.NewAndroid(p.userAgent),
		platforms.NewLinux(p.userAgent),
		platforms.NewFirefoxOS(p.userAgent),
		platforms.NewChromeOS(p.userAgent),
		platforms.NewUnknown(p.userAgent),
	}

	for _, matcher := range matchers {
		if matcher.Match() {
			p.matcher = matcher
			break
		}
	}

	return p.matcher
}

// Name returns the name of the platform.
func (p *Platform) Name() string {
	return p.getMatcher().Name()
}

// IsAdobeAir returns true if the user agent string matches Adobe AIR.
func (p *Platform) IsAdobeAir() bool {
	if _, ok := p.getMatcher().(*platforms.AdobeAir); ok {
		return true
	}

	return false
}

func (p *Platform) IsAdobeAirVersionCompatible(version string) bool {
	actualVersion := p.getMatcher().Version()
	return p.IsAdobeAir() && utils.VersionGTE(actualVersion, version)
}

// IsAndroid returns true if the user agent string matches Android.
func (p *Platform) IsAndroid() bool {
	if _, ok := p.getMatcher().(*platforms.Android); ok {
		return true
	}

	return false
}

// IsChromeOS returns true if the user agent string matches Chrome OS.
func (p *Platform) IsChromeOS() bool {
	if _, ok := p.getMatcher().(*platforms.ChromeOS); ok {
		return true
	}

	return false
}

func (p *Platform) IsChromeOSVersionCompatible(version string) bool {
	return p.IsChromeOS() &&
		utils.VersionGTE(p.getMatcher().Version(), version)
}

// IsIOS returns true if the user agent string matches iOS.
func (p *Platform) IsIOS() bool {
	if _, ok := p.getMatcher().(*platforms.IOS); ok {
		return true
	}

	return false
}

// IsWatchOS returns true if the user agent string matches Watch OS.
func (p *Platform) IsWatchOS() bool {
	if _, ok := p.getMatcher().(*platforms.WatchOS); ok {
		return true
	}

	return false
}

func (p *Platform) IsKaiOS() bool {
	if _, ok := p.getMatcher().(*platforms.KaiOS); ok {
		return true
	}

	return false
}

// IsLinux returns true if the user agent string matches Linux.
func (p *Platform) IsLinux() bool {
	if _, ok := p.getMatcher().(*platforms.Linux); ok {
		return true
	}

	return false
}

func (p *Platform) IsBlackBerry() bool {
	if _, ok := p.getMatcher().(*platforms.BlackBerry); ok {
		return true
	}

	return false
}

func (p *Platform) IsWindowsMobile() bool {
	if _, ok := p.getMatcher().(*platforms.WindowsMobile); ok {
		return true
	}

	return false
}

func (p *Platform) IsWindowsPhone() bool {
	if _, ok := p.getMatcher().(*platforms.WindowsPhone); ok {
		return true
	}

	return false
}

func (p *Platform) IsIOSApp() bool {
	if p.IsIOS() && !strings.Contains(p.userAgent, "Safari") {
		return true
	}
	return false
}

func (p *Platform) IsIOSWebview() bool {
	return p.IsIOSApp()
}

func (p *Platform) IsAndroidApp() bool {
	rg := regexp.MustCompile(`\bwv\b`)
	if p.IsAndroid() && rg.MatchString(p.userAgent) {
		return true
	}
	return false
}

func (p *Platform) IsAndroidWebview() bool {
	return p.IsAndroidApp()
}

func (p *Platform) IsWindows() bool {
	if _, ok := p.getMatcher().(*platforms.Windows); ok {
		return true
	}

	return false
}

func (p *Platform) IsWindowsXP() bool {
	v := p.getMatcher().Version()
	rg := regexp.MustCompile(`5\.[12]`)

	if p.IsWindows() && rg.MatchString(v) {
		return true
	}

	return false
}

func (p *Platform) IsWindowsVista() bool {
	v := p.getMatcher().Version()
	if p.IsWindows() && v == "6.0" {
		return true
	}
	return false
}

func (p *Platform) IsWindows7() bool {
	ver := p.getMatcher().Version()
	if p.IsWindows() && ver == "6.1" {
		return true
	}
	return false
}

func (p *Platform) IsWindows8() bool {
	ver := p.getMatcher().Version()
	rg := regexp.MustCompile(`6\.[2-3]`)
	if p.IsWindows() && rg.MatchString(ver) {
		return true
	}
	return false
}

func (p *Platform) IsWindows8_1() bool {
	ver := p.getMatcher().Version()
	if p.IsWindows() && ver == "6.3" {
		return true
	}
	return false
}

func (p *Platform) IsWindows10() bool {
	ver := p.getMatcher().Version()
	if p.IsWindows() && ver == "10.0" {
		return true
	}
	return false
}

func (p *Platform) IsWindowsRT() bool {
	if p.IsWindows8() && strings.Contains(p.userAgent, "ARM") {
		return true
	}
	return false
}

func (p *Platform) IsWindowsX64() bool {
	rg := regexp.MustCompile(`(Win64|x64|Windows NT 5\.2)`)
	if p.IsWindows() && rg.MatchString(p.userAgent) {
		return true
	}
	return false
}

func (p *Platform) IsWindowsWOW64() bool {
	rg := regexp.MustCompile(`(?i)WOW64`)
	if p.IsWindows() && rg.MatchString(p.userAgent) {
		return true
	}
	return false
}

func (p *Platform) IsWindowsX64Inclusive() bool {
	return p.IsWindowsX64() || p.IsWindowsWOW64()
}

// IsWindowsTouchScreenDesktop returns true if the platform is Windows 8 and the user agent string contains Touch.
func (p *Platform) IsWindowsTouchScreenDesktop() bool {
	if p.IsWindows() && strings.Contains(p.userAgent, "Touch") {
		return true
	}
	return false
}
