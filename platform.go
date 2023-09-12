package browser

// PlatformMatcher is an interface for user agent platform matchers.
// A platform matcher is used to detect the platform from the user agent string.
type PlatformMatcher interface {
	Matcher
	Version() string // Version returns the version of the platform.
}

