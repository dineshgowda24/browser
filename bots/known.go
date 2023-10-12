package bots

import (
	"regexp"
)

// Known is a bot that matches when the user agent contains a keyword.
// It is used for bots that are known to be bots
type Known struct {
	userAgent  string
	bots       map[string]string
	matchedBot string
}

// NewKnown returns a new Known bot.
func NewKnown(userAgent string, bots map[string]string) *Known {
	return &Known{
		userAgent: userAgent,
		bots:      bots,
	}
}

// Name returns the keyword that matched.
func (k *Known) Name() string {
	if !k.matched() {
		k.Match()
	}

	return k.matchedBot
}

// Match returns true if the user agent contains a keyword.
func (k *Known) Match() bool {
	if k.matched() {
		return true
	}

	for b, n := range k.bots {
		rg := regexp.MustCompile(`(?i)` + b)
		if rg.MatchString(k.userAgent) {
			k.matchedBot = n
			return true
		}
	}
	return false
}

func (k *Known) matched() bool {
	return k.matchedBot != ""
}
