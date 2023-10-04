package browser

import (
	"strings"

	"github.com/dineshgowda24/browser/matchers"
)

const (
	userAgentSizeLimit      = 2048 // 2KB
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

