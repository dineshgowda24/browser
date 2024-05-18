package browser

import (
	"strings"

	"github.com/dineshgowda24/browser/matchers"
)

const (
	userAgentSizeLimit = 2048 // 2KiB
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
	userAgent  string         // the user agent string
	platform   *Platform      // detected platform
	device     *Device        // detected device
	bot        *Bot           // detected bot
	matcher    BrowserMatcher // detected browser matcher
	registered bool           // whether the matcher has been registered
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

	p, _ := NewPlatform(userAgent)
	d, _ := NewDevice(userAgent)

	bot, err := NewBot(userAgent)
	if err != nil {
		return nil, err
	}

	b := &Browser{
		userAgent: userAgent,
		platform:  p,
		device:    d,
		bot:       bot,
	}
	b.register()

	return b, nil
}

// register registers the browser matcher detected from the user agent string.
func (b *Browser) register() {
	// add all your browser matchers here, order matters
	parser := matchers.NewUAParser(b.userAgent)
	matchers := []BrowserMatcher{
		matchers.NewAlipay(parser),
		matchers.NewNokia(parser),
		matchers.NewUCBrowser(parser),
		matchers.NewBlackBerry(parser),
		matchers.NewBrave(parser),
		matchers.NewOpera(parser),
		matchers.NewOtter(parser),
		matchers.NewInstagram(parser),
		matchers.NewSnapchat(parser),
		matchers.NewWeibo(parser),
		matchers.NewMicroMessenger(parser),
		matchers.NewQQ(parser),
		matchers.NewElectron(parser),
		matchers.NewDuckDuckGo(parser),
		matchers.NewGoogleSearchApp(parser),
		matchers.NewHuaweiBrowser(parser),
		matchers.NewKonqueror(parser),
		matchers.NewMaxthon(parser),
		matchers.NewMiuiBrowser(parser),
		matchers.NewPaleMoon(parser),
		matchers.NewPuffin(parser),
		matchers.NewEdge(parser),
		matchers.NewInternetExplorer(parser),
		matchers.NewSamsungBrowser(parser),
		matchers.NewSogouBrowser(parser),
		matchers.NewVivaldi(parser),
		matchers.NewVivoBrowser(parser),
		matchers.NewSputnik(parser),
		matchers.NewYaaniBrowser(parser),
		matchers.NewYandex(parser),
		matchers.NewFirefox(parser),
		matchers.NewChrome(parser), // chrome should be before safari
		matchers.NewSafari(parser), // chrome and safari must be at the end
		matchers.NewUnknown(parser),
	}

	for _, matcher := range matchers {
		if matcher.Match() {
			b.matcher = matcher
			break
		}
	}

	b.registered = true
}

// getMatcher returns the browser matcher detected from the user agent string.
// If the browser matcher has not been registered, it will be registered.
func (b *Browser) getMatcher() BrowserMatcher {
	if !b.registered {
		b.register()
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

// Bot returns the bot of the browser.
func (b *Browser) Bot() *Bot {
	return b.bot
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

// IsNokia returns true if the browser is Nokia.
//
// https://www.nokia.com/
func (b *Browser) IsNokia() bool {
	if _, ok := b.getMatcher().(*matchers.Nokia); ok {
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

// IsBlackBerry returns true if the browser is BlackBerry.
//
// https://www.blackberry.com/
func (b *Browser) IsBlackBerry() bool {
	if _, ok := b.getMatcher().(*matchers.BlackBerry); ok {
		return true
	}

	return false
}

// IsBrave returns true if the browser is Brave.
//
// https://brave.com/
func (b *Browser) IsBrave() bool {
	if _, ok := b.getMatcher().(*matchers.Brave); ok {
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

// IsInstagram returns true if the browser is Instagram.
//
// https://www.instagram.com/
func (b *Browser) IsInstagram() bool {
	if _, ok := b.getMatcher().(*matchers.Instagram); ok {
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

// IsWeibo returns true if the browser is Weibo.
//
// https://weibo.com/
func (b *Browser) IsWeibo() bool {
	if _, ok := b.getMatcher().(*matchers.Weibo); ok {
		return true
	}

	return false
}

// IsWechat returns true if the browser is Wechat.
// Wechat is an alias for MicroMessenger.
func (b *Browser) IsWechat() bool {
	return b.IsMicroMessenger()
}

// IsMicroMessenger returns true if the browser is MicroMessenger.
func (b *Browser) IsMicroMessenger() bool {
	if _, ok := b.getMatcher().(*matchers.MicroMessenger); ok {
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

// IsElectron returns true if the browser is Electron.
//
// https://www.electronjs.org/
func (b *Browser) IsElectron() bool {
	if _, ok := b.getMatcher().(*matchers.Electron); ok {
		return true
	}

	return false
}

// IsFirefox returns true if the browser is Firefox.
//
// https://www.mozilla.org/firefox/
func (b *Browser) IsFirefox() bool {
	if _, ok := b.getMatcher().(*matchers.Firefox); ok {
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

// IsGoogleSearchApp returns true if the browser is GoogleSearchApp.
//
// https://www.google.com/search/about/
func (b *Browser) IsGoogleSearchApp() bool {
	if _, ok := b.getMatcher().(*matchers.GoogleSearchApp); ok {
		return true
	}

	return false
}

// IsHuaweiBrowser returns true if the browser is HuaweiBrowser.
//
// https://consumer.huawei.com/en/mobileservices/browser/
func (b *Browser) IsHuaweiBrowser() bool {
	if _, ok := b.getMatcher().(*matchers.HuaweiBrowser); ok {
		return true
	}

	return false
}

// IsKonqueror returns true if the browser is Konqueror.
//
// https://konqueror.org/
func (b *Browser) IsKonqueror() bool {
	if _, ok := b.getMatcher().(*matchers.Konqueror); ok {
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

// IsMiuiBrowser returns true if the browser is MiuiBrowser.
func (b *Browser) IsMiuiBrowser() bool {
	if _, ok := b.getMatcher().(*matchers.MiuiBrowser); ok {
		return true
	}

	return false
}

// IsPaleMoon returns true if the browser is PaleMoon.
//
// https://www.palemoon.org/
func (b *Browser) IsPaleMoon() bool {
	if _, ok := b.getMatcher().(*matchers.PaleMoon); ok {
		return true
	}

	return false
}

// IsPuffin returns true if the browser is Puffin.
//
// https://www.puffin.com/
func (b *Browser) IsPuffin() bool {
	if _, ok := b.getMatcher().(*matchers.Puffin); ok {
		return true
	}

	return false
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

// IsSamsungBrowser returns true if the browser is SamsungBrowser.
func (b *Browser) IsSamsungBrowser() bool {
	if _, ok := b.getMatcher().(*matchers.SamsungBrowser); ok {
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

// IsVivaldi returns true if the browser is Vivaldi.
//
// https://vivaldi.com/
func (b *Browser) IsVivaldi() bool {
	if _, ok := b.getMatcher().(*matchers.Vivaldi); ok {
		return true
	}

	return false
}

// IsVivoBrowser returns true if the browser is VivoBrowser.
//
// https://www.vivo.com/
func (b *Browser) IsVivoBrowser() bool {
	if _, ok := b.getMatcher().(*matchers.VivoBrowser); ok {
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

// IsYaaniBrowser returns true if the browser is YaaniBrowser.
func (b *Browser) IsYaaniBrowser() bool {
	if _, ok := b.getMatcher().(*matchers.YaaniBrowser); ok {
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

// IsChrome returns true if the browser is Chrome.
//
// https://www.google.com/chrome/
func (b *Browser) IsChrome() bool {
	if _, ok := b.getMatcher().(*matchers.Chrome); ok {
		return true
	}

	return false
}

// IsSafari returns true if the browser is Safari.
//
// https://www.apple.com/safari/
func (b *Browser) IsSafari() bool {
	if _, ok := b.getMatcher().(*matchers.Safari); ok {
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
