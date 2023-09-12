package browser

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

