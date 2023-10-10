package browser

import (
	"strings"

	"github.com/dineshgowda24/browser/matchers"
)

const (
	userAgentSizeLimit = 2048 // 2KB
)

// Matcher is an interface for user agent matchers.
// A matcher is used to detect a match from the user agent string.
type Matcher interface {
	Match() bool
	Name() string
}

// BrowserMatcher is an interface for browser matchers.
type BrowserMatcher interface {
	Matcher          // BrowserMatcher implements Matcher interface.
	Version() string // Version returns the full version of the browser.
}

// Browser is a struct that contains information about the user agent's browser.
// It contains information about the browser's name, version, platform and device.
type Browser struct {
	userAgent string         // the user agent string
	platform  *Platform      // detected platform
	device    *Device        // detected device
	matcher   BrowserMatcher // detected browser matcher
}

// NewBrowser creates a new browser based on the user agent string.
// It builds the browser by matching the user agent string with the known browsers.
//
// UserAgentSizeError is returned if the user agent string is larger than the limit.
func NewBrowser(userAgent string) (*Browser, error) {
	// if user agent is larger than the limit, return error
	if len(userAgent) > userAgentSizeLimit {
		return nil, ErrUserAgentSizeExceeded
	}

	b := &Browser{
		userAgent: userAgent,
		platform: &Platform{
			userAgent: userAgent,
		},
		device: &Device{
			userAgent: userAgent,
		},
	}

	b.getMatcher()
	return b, nil
}

func (b *Browser) getMatcher() BrowserMatcher {
	if b.matcher != nil {
		return b.matcher
	}

	// add all your browser matchers here, order matters
	matchers := []BrowserMatcher{
		matchers.NewAlipay(b.userAgent),
		matchers.NewNokia(b.userAgent),
		matchers.NewUCBrowser(b.userAgent),
		matchers.NewBlackBerry(b.userAgent),
		matchers.NewOpera(b.userAgent),
		matchers.NewOtter(b.userAgent),
		matchers.NewInstagram(b.userAgent),
		matchers.NewSnapchat(b.userAgent),
		matchers.NewWeibo(b.userAgent),
		matchers.NewMicroMessenger(b.userAgent),
		matchers.NewQQ(b.userAgent),
		matchers.NewElectron(b.userAgent),
		matchers.NewDuckDuckGo(b.userAgent),
		matchers.NewGoogleSearchApp(b.userAgent),
		matchers.NewHuaweiBrowser(b.userAgent),
		matchers.NewKonqueror(b.userAgent),
		matchers.NewMaxthon(b.userAgent),
		matchers.NewMiuiBrowser(b.userAgent),
		matchers.NewPaleMoon(b.userAgent),
		matchers.NewPuffin(b.userAgent),
		matchers.NewEdge(b.userAgent),
		matchers.NewInternetExplorer(b.userAgent),
		matchers.NewSamsungBrowser(b.userAgent),
		matchers.NewSogouBrowser(b.userAgent),
		matchers.NewVivoBrowser(b.userAgent),
		matchers.NewSputnik(b.userAgent),
		matchers.NewYaaniBrowser(b.userAgent),
		matchers.NewYandex(b.userAgent),
		matchers.NewChrome(b.userAgent), // chrome should be before safari
		matchers.NewSafari(b.userAgent), // chrome and safari must be at the end
		matchers.NewUnknown(b.userAgent),
	}

	for _, matcher := range matchers {
		if matcher.Match() {
			b.matcher = matcher
			break
		}
	}

	return b.matcher
}

// Name returns the name of the matched browser.
func (b *Browser) Name() string {
	return b.getMatcher().Name()
}

// Version returns the full version of the matched browser.
// Examples:
// - 12.1.1
// - 04.34
func (b *Browser) Version() string {
	return b.getMatcher().Version()
}

// ShortVersion returns the short version of the version of the matched browser.
// Below are some examples of how the short version is returned:
//
// 12.1.1 => 12
//
// 04.34  => 04
func (b *Browser) ShortVersion() string {
	return strings.Split(b.Version(), ".")[0]
}

// Platform returns the platform of the browser.
// It can used to further query information about the platform.
func (b *Browser) Platform() *Platform {
	return b.platform
}

// Device returns the device of the browser.
// It can used to further query information about the device.
func (b *Browser) Device() *Device {
	return b.device
}

// IsAliPay returns true if the browser is AliPay.
//
// https://www.alipay.com/
func (b *Browser) IsAlipay() bool {
	if _, ok := b.getMatcher().(*matchers.Alipay); ok {
		return true
	}

	return false
}

// IsBlackBerry returns true if the browser is BlackBerry.
//
// https://www.blackberry.com/
func (b *Browser) IsBlackBerry() bool {
	if _, ok := b.getMatcher().(*matchers.BlackBerry); ok {
		return true
	}

	return false
}

// IsDuckDuckGo returns true if the browser is DuckDuckGo.
//
// https://duckduckgo.com/
func (b *Browser) IsDuckDuckGo() bool {
	if _, ok := b.getMatcher().(*matchers.DuckDuckGo); ok {
		return true
	}

	return false
}

// IsElectron returns true if the browser is Electron.
//
// https://www.electronjs.org/
func (b *Browser) IsElectron() bool {
	if _, ok := b.getMatcher().(*matchers.Electron); ok {
		return true
	}

	return false
}

// IsInstagram returns true if the browser is Instagram.
//
// https://www.instagram.com/
func (b *Browser) IsInstagram() bool {
	if _, ok := b.getMatcher().(*matchers.Instagram); ok {
		return true
	}

	return false
}

// IsMaxthon returns true if the browser is Maxthon.
//
// https://www.maxthon.com/
func (b *Browser) IsMaxthon() bool {
	if _, ok := b.getMatcher().(*matchers.Maxthon); ok {
		return true
	}

	return false
}

// IsMicroMessenger returns true if the browser is MicroMessenger.
func (b *Browser) IsMicroMessenger() bool {
	if _, ok := b.getMatcher().(*matchers.MicroMessenger); ok {
		return true
	}

	return false
}

// IsWechat returns true if the browser is Wechat.
// Wechat is an alias for MicroMessenger.
func (b *Browser) IsWechat() bool {
	return b.IsMicroMessenger()
}

// IsMiuiBrowser returns true if the browser is MiuiBrowser.
func (b *Browser) IsMiuiBrowser() bool {
	if _, ok := b.getMatcher().(*matchers.MiuiBrowser); ok {
		return true
	}

	return false
}

func (b *Browser) IsNokia() bool {
	if _, ok := b.getMatcher().(*matchers.Nokia); ok {
		return true
	}

	return false
}

// IsOpera returns true if the browser is Opera.
//
// https://www.opera.com/
func (b *Browser) IsOpera() bool {
	if _, ok := b.getMatcher().(*matchers.Opera); ok {
		return true
	}

	return false
}

// IsOtter returns true if the browser is Otter.
//
// https://otter-browser.org/
func (b *Browser) IsOtter() bool {
	if _, ok := b.getMatcher().(*matchers.Otter); ok {
		return true
	}

	return false
}

// IsQQ returns true if the browser is QQ.
//
// https://browser.qq.com/
func (b *Browser) IsQQ() bool {
	if _, ok := b.getMatcher().(*matchers.QQ); ok {
		return true
	}

	return false
}

// IsSnapchat returns true if the browser is Snapchat.
//
// https://www.snapchat.com/
func (b *Browser) IsSnapchat() bool {
	if _, ok := b.getMatcher().(*matchers.Snapchat); ok {
		return true
	}

	return false
}

// IsSougouBrowser returns true if the browser is SougouBrowser.
//
// https://www.sogou.com/
func (b *Browser) IsSougouBrowser() bool {
	if _, ok := b.getMatcher().(*matchers.SogouBrowser); ok {
		return true
	}

	return false
}

// IsSputnik returns true if the browser is Sputnik.
func (b *Browser) IsSputnik() bool {
	if _, ok := b.getMatcher().(*matchers.Sputnik); ok {
		return true
	}

	return false
}

// IsUCBrowser returns true if the browser is UCBrowser.
//
// https://www.ucweb.com/
func (b *Browser) IsUCBrowser() bool {
	if _, ok := b.getMatcher().(*matchers.UCBrowser); ok {
		return true
	}

	return false
}

// IsWeibo returns true if the browser is Weibo.
//
// https://weibo.com/
func (b *Browser) IsWeibo() bool {
	if _, ok := b.getMatcher().(*matchers.Weibo); ok {
		return true
	}

	return false
}

// IsYandex returns true if the browser is Yandex.
//
// https://yandex.com/
func (b *Browser) IsYandex() bool {
	if _, ok := b.getMatcher().(*matchers.Yandex); ok {
		return true
	}

	return false
}

// IsUnknown returns true if the browser is Unknown.
// Browsers are considered unknown when they do not match any of the known browsers.
func (b *Browser) IsUnknown() bool {
	if _, ok := b.getMatcher().(*matchers.Unknown); ok {
		return true
	}

	return false
}

// IsBrowserKnown returns true if a match was found for the browser.
func (b *Browser) IsBrowserKnown() bool {
	return !b.IsUnknown()
}

// IsSafariWebappMode returns true if the browser is Safari Webapp Mode.
// It is true for iOS devices which have AppleWebKit in the user agent string.
func (b *Browser) IsSafariWebappMode() bool {
	d := b.Device()
	return (d.IsIPad() || d.IsIPhone()) &&
		strings.Contains(b.userAgent, "AppleWebKit")
}

// IsEdge returns true if the browser is Edge.
func (b *Browser) IsEdge() bool {
	if _, ok := b.getMatcher().(*matchers.Edge); ok {
		return true
	}

	return false
}

// IsInternetExplorer returns true if the browser is Internet Explorer.
func (b *Browser) IsInternetExplorer() bool {
	if _, ok := b.getMatcher().(*matchers.InternetExplorer); ok {
		return true
	}

	return false
}
